package WeWorkFinanceSDK

var MessageOfType = func(action, typ string) Message {
	switch action {
	case "switch":
		return &SwitchMessage{}
	}
	switch typ {
	case "text":
		return &TextMessage{}
	case "image":
		return &ImageMessage{}
	case "revoke":
		return &RevokeMessage{}
	case "agree":
		fallthrough
	case "disagree":
		return &AgreeMessage{}
	case "voice":
		return &VoiceMessage{}
	case "video":
		return &VideoMessage{}
	case "card":
		return &CardMessage{}
	case "location":
		return &LocationMessage{}
	case "emotion":
		return &EmotionMessage{}
	case "link":
		return &LinkMessage{}
	case "chatrecord":
		return &ChatRecordMessage{}
	case "todo":
		return &TodoMessage{}
	case "vote":
		return &VoteMessage{}
	case "collect":
		return &CollectMessage{}
	case "redpacket":
		return &RedPacketMessage{}
	case "meeting":
		return &MeetingMessage{}
	case "docmsg":
		return &DocMessage{}
	case "markdown":
		return &MarkdownMessage{}
	case "news":
		return &NewsMessage{}
	case "calendar":
		return &CalendarMessage{}
	case "mixed":
		return &MixedMessage{}
	case "external_redpacket":
		return &ExternalRedPacketMessage{}
	case "sphfeed":
		return &SphFeedMessage{}
	case "voip_doc_share":
		return &VoipDocShareMessage{}
	case "file":
		return &FileMessage{}
	case "meeting_voice_call":
	case "voiptext":
	case "qydiskfile":
	case "weapp":
	}
	return &BaseMessage{}
}
