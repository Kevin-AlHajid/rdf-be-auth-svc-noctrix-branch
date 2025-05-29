package usecase

import (
	"testing"

	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/domain"
	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"fmt"
)

var roleRepository = &repository.RoleRepoMock{Mock: mock.Mock{}}
var roleUsecase = NewRoleUsecase(roleRepository)

func TestCreateRoleSuccess(t *testing.T) {
	//Creating test role that will be inputted
	testrole := &domain.Role{
		ID:          01,
		Name:        "Admin",
		Description: "Test Role",
	}
	//Putting in the method and the input
	roleRepository.Mock.On("Create", testrole).Return(testrole, nil)
	result, err := roleUsecase.CreateRole(testrole)
	assert.NotNil(t, result)
	assert.Nil(t, err)
} //SUCCESS

func TestGetRoleSuccess(t *testing.T) {
	//Creating the role that will be searched
	testrole := domain.Role{
		ID:          01,
		Name:        "Admin",
		Description: "Test Role",
	}
	roleRepository.Mock.On("FindById", 01).Return(testrole)
	result, err := roleUsecase.GetRole(01)
	assert.Equal(t, testrole.ID, result.ID)
	assert.Nil(t, err)
	fmt.Println(result)
}

func TestListRoleSuccess(t *testing.T) {
	testrole := []*domain.Role{
		&domain.Role{ID: 01, Name: "Admin"},
		&domain.Role{ID: 02, Name: "Co-Admin"},
		&domain.Role{ID: 03, Name: "Vice Admin"},
	}
	roleRepository.Mock.On("FindAll").Return(testrole, nil)
	result, err := roleUsecase.ListRoles()
	assert.NotNil(t, result)
	assert.Nil(t, err)
	fmt.Println(result)
}

func TestUpdateRoleSuccess(t *testing.T) {
	//old role
	testrole := domain.Role{
		ID:   01,
		Name: "Admin",
	}
	//new updated role
	updatedrole := domain.Role{
		ID:   01,
		Name: "President",
	}
	//testing so that when update method is called with "anything" as input
	//it returns the updated role and nil errors
	roleRepository.Mock.On("FindById", mock.Anything).Return(testrole, nil)
	roleRepository.Mock.On("Update", mock.Anything).Return(&updatedrole, nil)
	result, err := roleUsecase.UpdateRole(&testrole)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, "President", result.Name)
	fmt.Println(result)
}

func TestDeleteRoleSuccess(t *testing.T) {
	//role gonna be deleted
	roletobedeleted := domain.Role{
		ID:   03,
		Name: "Supervisor",
	}
	roleRepository.Mock.On("FindById", mock.Anything).Return(roletobedeleted)
	roleRepository.Mock.On("Delete", &roletobedeleted).Return("role deleted", nil)
	result, err := roleUsecase.DeleteRole(&roletobedeleted)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}
