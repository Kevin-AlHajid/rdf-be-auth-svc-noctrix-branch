package domain

type RoleScope struct {
	RoleScopeId int64
	RoleId      int64
	ScopeId     int64
}

type RoleScopeRepo interface {
	AssignScopetoRole(*RoleScope) (*RoleScope, error)
	ReduceScopeofRole(*RoleScope) (*RoleScope, error)
}
