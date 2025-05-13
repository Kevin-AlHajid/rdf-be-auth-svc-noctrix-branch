package usecase

import (
	"rdf-be-auth-svc-noctrix-branch/domain"
	"time"
)

type RoleUsecase struct {
	RoleRepository domain.RoleRepository
}

func NewRoleUsecase(repo domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{RoleRepository: repo}
}

func (r *RoleUsecase) CreateRole(Name *domain.Role) (*domain.Role, error) {
	return r.RoleRepository.Create(Name)
}

func (r *RoleUsecase) ListRoles() ([]*domain.Role, error) {
	return r.RoleRepository.FindAll()
}

func (r *RoleUsecase) GetRole(rolename *domain.Role) (*domain.Role, error) {
	findRole, err := r.RoleRepository.FindByName(rolename.Name)
	if err != nil {
		return nil, err
	}
	return findRole, nil
}

func (r *RoleUsecase) UpdateRole(rolename *domain.Role) (*domain.Role, error) {
	oldrole, err := r.RoleRepository.FindByName(rolename.Name)
	if err != nil || oldrole == nil {
		return nil, err
	}
	now := time.Now()
	oldrole.UpdatedAt = &now
	return r.RoleRepository.Update(oldrole)
}

func (r *RoleUsecase) DeleteRole(rolename *domain.Role) error {
	roletobedeleted, err := r.RoleRepository.FindByName(rolename.Name)
	if err != nil || roletobedeleted == nil {
		return err
	}
	return r.RoleRepository.Delete(roletobedeleted)
}
