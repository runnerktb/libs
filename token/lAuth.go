package token

type AuthorizationInfo struct {
	UserID         string   `json:"id"`
	Username       string   `json:"username"`
	IsOrgAdmin     *int     `json:"isorgadmin"`
	IsActive       int      `json:"isactive"`
	OrganizationId string   `json:"organization_id"`
	App            string   `json:"app"`
	Email          string   `json:"email"`
	Exp            int      `json:"exp"`
	Name           string   `json:"name"`
	RealmID        string   `json:"realm_id"`
	UserAccess     []string `json:"user_access"`
}
