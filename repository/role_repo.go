package repository

import (
	"time"

	"github.com/Kevin-AlHajid/rdf-be-auth-svc-noctrix-branch/domain"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) domain.RoleRepository {
	return &RoleRepository{db}
}

func (r *RoleRepository) Create(role *domain.Role) (*domain.Role, error) {
	now := time.Now()
	newrole := &domain.Role{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   &now,
	}
	err := r.db.Create(newrole).Error
	if err != nil {
		return nil, err
	}
	return newrole, nil
}

func (r *RoleRepository) FindAll() ([]*domain.Role, error) {
	var roles []*domain.Role
	//Initiating roles slice
	err := r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepository) FindById(roleid int) (*domain.Role, error) {
	var role *domain.Role

	err := r.db.Where("id = ?", roleid).First(&role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepository) Update(roledetail *domain.Role) (*domain.Role, error) {
	err := r.db.Save(roledetail).Error
	if err != nil {
		return nil, err
	}
	return roledetail, nil
}

func (r *RoleRepository) Delete(rolename *domain.Role) (string, error) {
	err := r.db.Delete(rolename).Error
	if err != nil {
		return "", err
	}
	return "Role Deleted", nil
}
