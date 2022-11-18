package WeWorkFinanceSDK

import "C"
import (
	"crypto/rsa"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

type GetChatDataOptions struct {
	Sequence        uint64
	Limit           int64
	Proxy           string
	ProxyCredential string
	Timeout         int
	SkipDecrypt     bool
}
type GetMediaDataOptions struct {
	Index           string
	FileID          string
	Proxy           string
	ProxyCredential string
	Timeout         int
}
type ClientOptions struct {
	Proxy           string
	TempDir         string
	ProxyCredential string
	Timeout         int
	PrivateKey      *rsa.PrivateKey
	PrivateKeyFn    func(ver int) (*rsa.PrivateKey, error)
}
type ClientOptionBuilder struct {
	Options ClientOptions
	Errors  []error
	Client  *client
}

func (b *ClientOptionBuilder) PrivateKeys(m map[int]string) *ClientOptionBuilder {
	b.PrivateKeyFn(func(ver int) (string, error) {
		return m[ver], nil
	})
	return b
}

func (b *ClientOptionBuilder) PrivateKeyFn(fn func(ver int) (string, error)) *ClientOptionBuilder {
	keys := sync.Map{}
	b.Options.PrivateKeyFn = func(ver int) (key *rsa.PrivateKey, err error) {
		val, _ := keys.Load(ver)
		if key, _ = val.(*rsa.PrivateKey); key != nil {
			return
		}
		s, err := fn(ver)
		if err != nil || s == "" {
			return nil, err
		}
		key, err = ParsePrivateKey(s)
		if err == nil {
			keys.Store(ver, key)
		}
		return
	}
	return b
}

func (b *ClientOptionBuilder) Proxy(value string) *ClientOptionBuilder {
	b.Options.Proxy = value
	return b
}

func (b *ClientOptionBuilder) ProxyCredential(value string) *ClientOptionBuilder {
	b.Options.ProxyCredential = value
	return b
}

func (b *ClientOptionBuilder) ParseEnv() *ClientOptionBuilder {
	if b.Options.PrivateKey == nil {
		pkFile := os.Getenv("WWF_PRIVATE_KEY_FILE")
		pk := os.Getenv("WWF_PRIVATE_KEY")

		if pkFile != "" && pk == "" {
			file, err := os.ReadFile(pkFile)
			if err != nil {
				b.Errors = append(b.Errors, err)
				return b
			}
			pk = string(file)
		}
		if pk != "" {
			b.PrivateKey(pk)
		}
	}
	if b.Options.Proxy == "" {
		b.Proxy(os.Getenv("WWF_PROXY"))
	}
	if b.Options.ProxyCredential == "" {
		b.ProxyCredential(os.Getenv("WWF_PROXY_CREDENTIAL"))
	}
	if b.Options.Timeout == 0 {
		b.Options.Timeout, _ = strconv.Atoi(os.Getenv("WWF_TIMEOUT"))
	}
	return b
}

func (b *ClientOptionBuilder) PrivateKey(v string) *ClientOptionBuilder {
	key, err := ParsePrivateKey(v)
	if err != nil {
		b.Errors = append(b.Errors, err)
	} else {
		b.Options.PrivateKey = key
	}
	return b
}

func (b *ClientOptionBuilder) Apply() (Client, error) {
	options, err := b.Build()
	if err != nil {
		return nil, err
	}
	b.Client.options = options
	return b.Client, nil
}

func (b *ClientOptionBuilder) MustApply() Client {
	client, err := b.Apply()
	if err != nil {
		panic(err)
	}
	return client
}

func (b *ClientOptionBuilder) MustBuild() ClientOptions {
	options, err := b.Build()
	if err != nil {
		panic(err)
	}
	return options
}

func (b *ClientOptionBuilder) Build() (ClientOptions, error) {
	if len(b.Errors) > 0 {
		return ClientOptions{}, b.Errors[0]
	}
	o := b.Options
	if o.PrivateKey == nil && o.PrivateKeyFn == nil {
		b.Errors = append(b.Errors, ErrorOfCode(10000, "no private key"))
	}
	return o, nil
}

func NewClientFromEnv() (Client, error) {
	corpID, _ := os.LookupEnv("WWF_CORP_ID")
	corpSecret, _ := os.LookupEnv("WWF_CORP_SECRET")
	corpSecretFile, _ := os.LookupEnv("WWF_CORP_SECRET_FILE")
	if corpSecretFile != "" && corpSecret == "" {
		file, err := os.ReadFile(corpSecretFile)
		if err != nil {
			panic(err)
		}
		corpSecret = string(file)
	}
	if corpID == "" {
		return nil, fmt.Errorf("corpId not founed from env")
	}
	if corpSecret == "" {
		return nil, fmt.Errorf("corpSecret not founed from env")
	}
	cli, err := NewClient(corpID, corpSecret)
	if err == nil {
		cli, err = cli.Options().ParseEnv().Apply()
	}
	return cli, err
}

func (c *client) Options() *ClientOptionBuilder {
	return &ClientOptionBuilder{Options: c.options, Client: c}
}

type SaveMediaOptions struct {
	Media   *Media
	TempDir string
	Dir     string // save to dir
	// Name    string // save as name
	// Path    string // save as final path
	// Meta    string

	Writer   io.Writer
	KeepData bool
	Force    bool
	CheckSum bool
}

func (c *client) GetCorpID() string {
	return c.corpID
}

type Client interface {
	GetCorpID() string
	Options() *ClientOptionBuilder
	CopyMediaData(o GetMediaDataOptions, w io.Writer) (sum int, err error)
	ReadMediaData(o GetMediaDataOptions) (data []byte, err error)
	GetMediaData(o GetMediaDataOptions) (*MediaData, error)

	// SaveMedia(o SaveMediaOptions) error

	GetChatData(o GetChatDataOptions) ([]*ChatData, error)
	Close()
}
