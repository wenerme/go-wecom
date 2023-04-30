package WeWorkFinanceSDK

import "encoding/json"

func Unmarshal(data []byte) (msg Message, err error) {
	var getType struct {
		Type   string `json:"msgtype,omitempty"`
		Action string `json:"action,omitempty"`
	}
	err = json.Unmarshal(data, &getType)
	if err != nil {
		return msg, err
	}
	msg = MessageOfType(getType.Action, getType.Type)
	msg.SetRaw(data)
	err = json.Unmarshal(data, msg)
	return msg, err
}

// UnmarshalToBase can avoid json error
func UnmarshalToBase(data []byte) (msg *BaseMessage, err error) {
	msg = &BaseMessage{}
	msg.SetRaw(data)
	err = json.Unmarshal(data, &msg)
	return msg, err
}
