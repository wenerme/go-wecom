//go:build !linux || !cgo

package WeWorkFinanceSDK

import (
	"fmt"
	"io"
)

func NewClient(corpID string, corpSecret string) (Client, error) {
	return nil, fmt.Errorf("unsupported platform")
}

func (c *client) GetChatData(o GetChatDataOptions) ([]*ChatData, error) {
	return nil, nil
}

func (c *client) Close() {
}

type client struct {
	options ClientOptions
	corpID  string
}

func (c *client) CopyMediaData(o GetMediaDataOptions, w io.Writer) (sum int, err error) {
	panic("implement me")
}

func (c *client) ReadMediaData(o GetMediaDataOptions) (data []byte, err error) {
	panic("implement me")
}

func (c *client) GetMediaData(o GetMediaDataOptions) (*MediaData, error) {
	panic("implement me")
}

var _ Client = &client{}
