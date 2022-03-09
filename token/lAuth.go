package token

type AuthorizationInfo struct {
	UserID         string   `json:"id"`
	Username       string   `json:"username"`
	IsOrgAdmin     *int     `json:"isorgadmin"`
	IsActive       int      `json:"isactive"`
	OrganizationId string   `json:"organizationid"`
	AppId          string   `json:"app"`
	Exp            int      `json:"exp"`
	UserAccess     []string `json:"user_access"`
}
