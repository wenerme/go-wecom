package wecom

type Conf struct {
	CorpID         string // 企业 ID
	AgentID        string // 应用 ID
	CorpSecret     string // 企业 secret
	ProviderSecret string // 第三方 secret
	SuiteSecret    string // 第三方应用 secret
	PermSecret     string // 第三方企业永久 secret
}
