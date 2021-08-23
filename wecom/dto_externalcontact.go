package wecom

type ListExternalContactRequest struct {
	UserID string `required:"true" json:"userid"`
}

type ListExternalContactResponse struct {
	ExternalUserID string `json:"external_userid"`
}
