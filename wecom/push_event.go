package wecom

import (
	"encoding/xml"
	"fmt"
	"reflect"
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
	if changeType != "" {
		if m := _eventChangeModels[infoType]; m != nil {
			v = m[changeType]
		}
	} else {
		v = _eventModels[infoType]
	}
	if v != nil {
		return reflect.New(reflect.TypeOf(v)).Interface().(EventModel)
	}
	return v
}

func UnmarshalEvent(data []byte) (EventModel, error) {
	b := &CommonPushEvent{}
	err := xml.Unmarshal(data, b)
	if err != nil {
		return nil, err
	}
	var model EventModel
	if b.InfoType != "" {
		model = NewEventModel(b.InfoType, b.ChangeType)
	} else {
		model = NewEventModel(b.Event, b.ChangeType)
	}
	if model == nil {
		return nil, fmt.Errorf("new event model for %q %q", b.InfoType, b.ChangeType)
	}
	return model, xml.Unmarshal(data, model)
}
