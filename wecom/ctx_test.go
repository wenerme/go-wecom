package wecom

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCtx(t *testing.T) {
	assert.Equal(t, "wecom.contextKey(Client)", ClientContextKey.String())
	assert.Nil(t, FromContext(context.Background()))
}
