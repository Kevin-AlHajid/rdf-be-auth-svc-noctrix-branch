package domain

import (
	"time"
)

// Making the main role struct
type Role struct {
	Id          int
	Name        string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type RoleRepository interface {
	Create(Name *Role) (*Role, error)
	FindAll() ([]*Role, error)
	FindById(roleid int) (*Role, error)
	Update(*Role) (*Role, error)
	Delete(*Role) (string, error)
}

// The role usecase & it's methods
type RoleUsecase interface {
	CreateRole(*Role) (*Role, error)
	GetRole(roleid int) (*Role, error)
	ListRoles() ([]*Role, error)
	UpdateRole(*Role) (*Role, error)
	DeleteRole(*Role) (string, error)
}

//RoleUsecase is meant to represent business logic, not technical data fetching details.
//It expects ready-to-use data that’s already been validated or collected, like *Role
