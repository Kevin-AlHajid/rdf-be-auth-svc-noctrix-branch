package domain

import "time"

//Making the main role struct
type Role struct {
	Id          int64
	Name        string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

type RoleRepository interface {
	Create(Name *Role) (*Role, error)
	FindAll() ([]*Role, error)
	FindByName(rolename string) (*Role, error)
	Update(rolename *Role) (*Role, error)
	Delete(rolename *Role) error
}

//The role usecase & it's methods
type RoleUsecase interface {
	CreateRole(Name *Role) (*Role, error)
	GetRole(rolename *Role) (*Role, error)
	ListRoles() ([]*Role, error)
	UpdateRole(rolename *Role) (*Role, error)
	DeleteRole(rolename *Role) error
}
