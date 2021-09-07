package wecom

type Conf struct {
	CorpID                string // 企业 ID
	AgentID               string // 应用 ID
	CorpSecret            string // 企业 secret
	ProviderSecret        string // 第三方 secret
	SuiteID               string // 第三方应用 ID
	SuiteSecret           string // 第三方应用 secret
	SuiteTicket           string // 第三方应用 ticket - 无法主动获取 - 会尝试从 Provider 获取
	AuthCorpID            string // 企业 ID
	AuthCorpPermanentCode string // 第三方企业永久授权码 - 会尝试从 Provider 获取

	// not used by client
	// ProviderToken  string // used for push event
	// EncodingAESKey string

	TokenProvider TokenProvider `json:"-"`
}
