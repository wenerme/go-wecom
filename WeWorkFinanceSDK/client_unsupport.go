//go:build !linux || !cgo

package WeWorkFinanceSDK

import "fmt"

func NewClient(corpID string, corpSecret string) (*Client, error) {
	return nil, fmt.Errorf("unsupported platform")
}

func (c *Client) GetChatData(o GetChatDataOptions) ([]*ChatData, error) {
	return nil, nil
}

func (c *Client) Close() {
}

type Client struct {
	options ClientOptions
}
