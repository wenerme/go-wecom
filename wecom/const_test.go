package wecom

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Gender UserGenderType `json:"gender,omitempty"`
}

func TestMarshal(t *testing.T) {
	tests := []struct {
		in  string
		o   interface{}
		out string
	}{
		{
			in:  `{"gender":"1"}`,
			o:   &testStruct{},
			out: `{"gender":1}`,
		},
		{
			in:  `{"gender":1}`,
			o:   &testStruct{},
			out: `{"gender":1}`,
		},
		{
			in:  `{"gender":null}`,
			o:   &testStruct{},
			out: `{}`,
		},
	}

	for _, test := range tests {
		if assert.NoError(t, json.Unmarshal([]byte(test.in), test.o)) {
			m, err := json.Marshal(test.o)
			if assert.NoError(t, err) {
				assert.Equal(t, test.out, string(m))
			}
		}
	}
}
