package wecom

func init() {
	registerClientAPIPath("/cgi-bin/webhook/send", "SendWebhookMessage", cRef.SendWebhookMessage)
}
