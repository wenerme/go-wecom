package wecom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhookModel(t *testing.T) {
	payload := SendPayload{}
	for _, v := range []MessageContent{
		SendTextContent{},
		SendFileContent{},
		SendImageContent{},
		SendTemplateCardContent{},
		SendMarkdownContent{},
		SendNewsContent{},
		&SendTextContent{},
		&SendFileContent{},
		&SendImageContent{},
		&SendTemplateCardContent{},
		&SendMarkdownContent{},
		&SendNewsContent{},
	} {
		assert.NoError(t, payload.SetContent(v))
	}
}
