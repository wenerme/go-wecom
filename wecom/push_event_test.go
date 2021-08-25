package wecom

import (
	"encoding/xml"
	"io/fs"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterEventModel(t *testing.T) {
	e := NewEventModel("cancel_auth", "")
	assert.NoError(t, xml.Unmarshal([]byte(`<xml><AuthCorpId>test</AuthCorpId></xml>`), e))
	assert.Equal(t, "cancel_auth", e.EventInfoType())
	assert.Equal(t, "test", e.(*CancelAuthPushEvent).AuthCorpID)
}

func TestPushEventSerialization(t *testing.T) {
	tdFs := os.DirFS("./testdata")
	files, err := fs.Glob(tdFs, "push/*.xml")
	assert.NoError(t, err)
	for _, file := range files {
		data, err := fs.ReadFile(tdFs, file)
		assert.NoError(t, err)
		e, err := UnmarshalEvent(data)
		if !assert.NoError(t, err) {
			log.Println("failed", file)
		}
		_ = e
	}
}
