package ginp

type UserInfo struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"username"`
	RoleID   int    `json:"role_id"`
	OrgID    int    `json:"org_id"`
	TenantID int    `json:"tenant_id"`
}
