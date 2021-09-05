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
	assert.Equal(t, "cancel_auth", e.EventType())
	assert.Equal(t, "test", e.(*CancelAuthPushEvent).AuthCorpID)
}

func TestPushEventSerialization(t *testing.T) {
	tdFs := os.DirFS("./testdata")
	files, err := fs.Glob(tdFs, "push/*.xml")
	assert.NoError(t, err)
	for _, file := range files {
		data, err := fs.ReadFile(tdFs, file)
		assert.NoError(t, err)
		e, ce, err := UnmarshalEvent(data)
		if !assert.NoError(t, err, file) {
			if ce != nil {
				log.Println("event", ce.GetEventType(), ce.GetTimestamp())
			}
		}
		_ = e
	}
}
