package repository

import (
	"errors"

	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/domain"
	"github.com/stretchr/testify/mock"
)

type RoleRepoMock struct {
	Mock mock.Mock
}

func (r *RoleRepoMock) Create(name *domain.Role) (*domain.Role, error) {
	args := r.Mock.Called(name)
	return args.Get(0).(*domain.Role), args.Error(1)
	//this gets the 0 position in arg, and converts them into domain.Role
	//while the second position gets converted into error
}

func (r *RoleRepoMock) FindAll() ([]*domain.Role, error) {
	args := r.Mock.Called()
	return args.Get(0).([]*domain.Role), args.Error(1)
}

func (r *RoleRepoMock) FindById(roleid int) (*domain.Role, error) {
	args := r.Mock.Called(roleid)
	if args.Get(0) == nil {
		return nil, errors.New("role tidak ditemukan")
	} else {
		rolesearched := args.Get(0).(domain.Role)
		return &rolesearched, nil
	}
}

func (r *RoleRepoMock) Update(role *domain.Role) (*domain.Role, error) {
	args := r.Mock.Called(role)
	return args.Get(0).(*domain.Role), args.Error(1)
}

func (r *RoleRepoMock) Delete(role *domain.Role) (string, error) {
	args := r.Mock.Called(role)
	return "Role Deleted", args.Error(1)
}
