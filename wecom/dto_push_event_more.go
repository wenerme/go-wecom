package wecom

type UnsubscribePushEvent SubscribePushEvent

// EventType impl EventModel
func (UnsubscribePushEvent) EventType() string {
	return "unsubscribe" //nolint:goconst
}

func init() {
	RegisterEventModel(UnsubscribePushEvent{})
}
