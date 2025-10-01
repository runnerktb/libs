package token

const (
	Root  = iota // no realm
	Super        // tti or ktb
	Admin        // org admin
	User         // org user
)

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
	Jti            string   `json:"jti"`
	Sid            string   `json:"sid"`
}

func (a *AuthorizationInfo) IsRoot() bool {
	return a.IsOrgAdmin != nil && *a.IsOrgAdmin == Root
}

func (a *AuthorizationInfo) IsSuper() bool {
	return a.IsOrgAdmin != nil && *a.IsOrgAdmin == Super
}

func (a *AuthorizationInfo) IsAdmin() bool {
	return a.IsOrgAdmin != nil && *a.IsOrgAdmin == Admin
}

func (a *AuthorizationInfo) IsUser() bool {
	return a.IsOrgAdmin != nil && *a.IsOrgAdmin == User
}
