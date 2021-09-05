package wecom

import (
	"encoding/xml"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

type EventModel interface {
	EventInfoType() string
}

type EventChangeModel interface {
	EventModel
	EventChangeType() string
}

func RegisterEventModel(e ...EventModel) {
	for _, v := range e {
		if c, ok := v.(EventChangeModel); ok {
			m := _eventChangeModels[v.EventInfoType()]
			if m == nil {
				m = map[string]EventChangeModel{}
				_eventChangeModels[v.EventInfoType()] = m
			}
			m[c.EventChangeType()] = c
		} else {
			_eventModels[v.EventInfoType()] = v
		}
	}
}

var (
	_eventModels       = map[string]EventModel{}
	_eventChangeModels = map[string]map[string]EventChangeModel{}
)

// NewEventModel return a ptr to model struct
// return ptr to prevent xml.Unmarshal(nil, &e) pass a ptr to interface
func NewEventModel(infoType string, changeType string) EventModel {
	var v EventModel
	if changeType == "" {
		v = _eventModels[infoType]
	} else if m := _eventChangeModels[infoType]; m != nil {
		v = m[changeType]
	}
	if v != nil {
		return reflect.New(reflect.TypeOf(v)).Interface().(EventModel)
	}
	return v
}

func UnmarshalCommonEvent(data []byte) (*CommonPushEvent, error) {
	b := &CommonPushEvent{}
	err := xml.Unmarshal(data, b)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal comment event")
	}
	return b, nil
}

func UnmarshalEvent(data []byte) (EventModel, *CommonPushEvent, error) {
	b, err := UnmarshalCommonEvent(data)
	if err != nil {
		return nil, nil, err
	}
	model := NewEventModel(b.GetEventType(), b.ChangeType)
	if model == nil {
		return nil, b, fmt.Errorf("no event model for %q %q", b.GetEventType(), b.ChangeType)
	}
	return model, b, xml.Unmarshal(data, model)
}
