package repository

import (
	"fmt"
	"testing"

	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/domain"
	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestCreateRole(t *testing.T) {
	//Create Role with mock
	repoRoleMock := RoleRepoMock{}
	newrole := &domain.Role{
		ID:   01,
		Name: "Admin",
	}
	repoRoleMock.Mock.On("Create", newrole).Return(newrole, nil)

	t.Run("CreateRoleMock", func(t *testing.T) {
		role, err := repoRoleMock.Create(newrole)
		assert.NotNil(t, role)
		assert.Nil(t, err)
	})

	//Create Role with DB
	db := infrastructure.ConnectDB()
	repoRole := NewRoleRepository(db)

	t.Run("CreateRoleDB", func(t *testing.T) {
		role, err := repoRole.Create(newrole)
		assert.NotNil(t, role)
		assert.Nil(t, err)
	})
}

// Note: DB already filled with pgadmin for testing below
func TestFindAllRole(t *testing.T) {
	db := infrastructure.ConnectDB()
	repoRole := NewRoleRepository(db)

	//Finding All roles in DB
	findall, err := repoRole.FindAll()
	assert.NotNil(t, findall)
	assert.Nil(t, err)

	//Checking if the roles returned are right
	assert.Equal(t, "Admin", findall[0].Name)
	assert.Equal(t, "Moderator", findall[2].Name)

	//Checking if the len of slice = all of roles in the db
	assert.Equal(t, 4, len(findall))
	//roles in db currently are: admin, co-admin, moderator, user
}

func TestFindByID(t *testing.T) {
	db := infrastructure.ConnectDB()
	repoRole := NewRoleRepository(db)

	//example input role
	role := &domain.Role{
		ID:   03,
		Name: "Moderator",
	}

	//Finding specific role in DB
	findrole, err := repoRole.FindById(role.ID)
	assert.NotNil(t, findrole)
	assert.Nil(t, err)

	//Checking if the found role is correct
	assert.Equal(t, "Moderator", findrole.Name)
}

func TestUpdateRole(t *testing.T) {
	db := infrastructure.ConnectDB()
	repoRole := NewRoleRepository(db)

	//testing in DB, changing name
	t.Run("UpdateName", func(t *testing.T) {
		existingrole := &domain.Role{}
		err := db.Where("name = ?", "Co-Admin").First(existingrole).Error
		assert.Nil(t, err)

		//Update the name
		existingrole.Name = "Vice-Admin"
		updated, err := repoRole.Update(existingrole)
		assert.Nil(t, err)
		assert.NotNil(t, updated)

		assert.Equal(t, "Vice-Admin", updated.Name)
	})

	//testing in DB, changing
	t.Run("UpdateDesc", func(t *testing.T) {
		existingrole := &domain.Role{}
		err := db.Where("name = ?", "Moderator").First(existingrole).Error
		assert.Nil(t, err)

		//Update the desc
		existingrole.Description = "Middle Brother"
		updated, err := repoRole.Update(existingrole)
		assert.Nil(t, err)
		assert.NotNil(t, updated)

		//Making Sure
		fmt.Println("Role Name:", existingrole.Name, "New Role Desc:", existingrole.Description)
	})
}

func TestDeleteRole(t *testing.T) {
	db := infrastructure.ConnectDB()
	repoRole := NewRoleRepository(db)

	existingrole := &domain.Role{}
	err := db.Where("name = ?", "Vice-Admin").First(existingrole).Error
	assert.Nil(t, err)

	delete, err := repoRole.Delete(existingrole)
	assert.Nil(t, err)
	assert.NotNil(t, delete)

	//Checking the db for the role
	role := &domain.Role{}
	err = db.Where("name = ?", "Vice-Admin").First(role).Error
	assert.Error(t, err)
}
