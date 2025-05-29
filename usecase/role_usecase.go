package usecase

import (
	"time"

	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/domain"
)

/*To Do List in here
CreateRole
ListRole
GetRole
UpdateRole
DeleteRole
*/

type RoleUsecase struct {
	Repository domain.RoleRepository
}

func NewRoleUsecase(repo domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{Repository: repo}
}

func (r *RoleUsecase) CreateRole(role *domain.Role) (*domain.Role, error) {
	return r.Repository.Create(role)
	//doesnt have a nil because Create() already has 2 returns
}

//*RoleUsecase lets you modify RoleUsecase itself

func (r *RoleUsecase) GetRole(roleid int) (*domain.Role, error) {
	findrole, err := r.Repository.FindById(roleid)
	if findrole == nil || err != nil {
		return nil, err
	} else {
		return findrole, nil
	}
}

func (r *RoleUsecase) ListRoles() ([]*domain.Role, error) {
	return r.Repository.FindAll()
}

func (r *RoleUsecase) UpdateRole(role *domain.Role) (*domain.Role, error) {
	//this is working on the belief that role is already the updated role
	oldrole, err := r.Repository.FindById(role.ID)
	if oldrole == nil || err != nil {
		return nil, err
	}
	now := time.Now()
	oldrole.Name = role.Name //role contains the updated fields
	oldrole.UpdatedAt = &now
	newrole, err := r.Repository.Update(oldrole)
	if err != nil {
		return nil, err
	}
	return newrole, nil
}

func (r *RoleUsecase) DeleteRole(role *domain.Role) (string, error) {
	//concepts i want to do
	//find role, delete role, find role again to make sure
	roletobedeleted, err := r.Repository.FindById(role.ID)
	if roletobedeleted == nil || err != nil {
		return "role not found", err
	}
	return r.Repository.Delete(roletobedeleted)
}
