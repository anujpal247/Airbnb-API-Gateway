package db

import (
	"AuthApp/models"
	"database/sql"
)

type RoleRepository interface {
	GetRoleById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRole() ([]*models.Role, error)
	CreateRole(name string, description string) (*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRoleById(id int64, name string, description string) (*models.Role, error)
}

type RoleRepositoryImpl struct {
	db *sql.DB
}

func NewRoleRepository(_db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: _db,
	}
}

func (rr *RoleRepositoryImpl) GetRoleById(id int64) (*models.Role, error) {
	return nil, nil
}

func (rr *RoleRepositoryImpl) GetRoleByName(name string) (*models.Role, error) {
	return nil, nil
}

func (rr *RoleRepositoryImpl) GetAllRole() ([]*models.Role, error) {
	return nil, nil
}

func (rr *RoleRepositoryImpl) CreateRole(name string, description string) (*models.Role, error) {
	return nil, nil
}

func (rr *RoleRepositoryImpl) DeleteRoleById(id int64) error {
	return nil
}

func (rr *RoleRepositoryImpl) UpdateRoleById(id int64, name string, description string) (*models.Role, error) {
	return nil, nil
}
