package db

import (
	"AuthApp/models"
	"database/sql"
)

type PermissionRepository interface {
	GetPermissionById(id int64) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermission() ([]*models.Permission, error)
	CreatePermission(name string, description string, resource string, action string) (*models.Permission, error)
	DeletePermissionById(id int64) error
	UpdatePermission(id int64, name string, description string, resource string, action string) (*models.Permission, error)
}

type PermissionRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionRepository(_db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: _db,
	}
}

func (pr *PermissionRepositoryImpl) GetPermissionById(id int64) (*models.Permission, error) {
	return nil, nil
}

func (pr *PermissionRepositoryImpl) GetPermissionByName(name string) (*models.Permission, error) {
	return nil, nil
}

func (pr *PermissionRepositoryImpl) GetAllPermission() ([]*models.Permission, error) {
	return nil, nil
}

func (pr *PermissionRepositoryImpl) CreatePermission(name string, description string, resource string, action string) (*models.Permission, error) {
	return nil, nil
}

func (pr *PermissionRepositoryImpl) DeletePermissionById(id int64) error {
	return nil
}

func (pr *PermissionRepositoryImpl) UpdatePermission(id int64, name string, description string, resource string, action string) (*models.Permission, error) {
	return nil, nil
}
