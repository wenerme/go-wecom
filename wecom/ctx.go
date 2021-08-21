package wecom

import (
	"context"
	"fmt"
)

type contextKey string

func (k contextKey) String() string {
	return fmt.Sprintf("wecom.contextKey(%s)", string(k))
}

const ClientContextKey = contextKey("Client")

func FromContext(ctx context.Context) *Client {
	value := ctx.Value(ClientContextKey)
	if value == nil {
		return nil
	}
	return value.(*Client)
}

func NewContext(ctx context.Context, v *Client) context.Context {
	return context.WithValue(ctx, ClientContextKey, v)
}
