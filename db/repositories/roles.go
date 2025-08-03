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
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE id = ?"

	row := rr.db.QueryRow(query, id)

	role := &models.Role{}

	if err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
		return nil, err
	}

	return role, nil
}

func (rr *RoleRepositoryImpl) GetRoleByName(name string) (*models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE name = ?"

	row := rr.db.QueryRow(query, name)

	role := &models.Role{}

	if err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
		return nil, err
	}
	return role, nil
}

func (rr *RoleRepositoryImpl) GetAllRole() ([]*models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles"

	rows, err := rr.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role

	for rows.Next() {
		role := &models.Role{}

		if err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (rr *RoleRepositoryImpl) CreateRole(name string, description string) (*models.Role, error) {

	query := "INSERT INTO roles (name, description, created_at, updated_at) VALUES(?, ?, NOW(), NOW()) "

	result, err := rr.db.Exec(query)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	role := &models.Role{
		Id:          id,
		Name:        name,
		Description: description,
		CreatedAt:   "",
		UpdatedAt:   "",
	}
	return role, nil
}

func (rr *RoleRepositoryImpl) DeleteRoleById(id int64) error {
	query := "DELETE FROM roles WHERE = ?"

	result, err := rr.db.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (rr *RoleRepositoryImpl) UpdateRoleById(id int64, name string, description string) (*models.Role, error) {
	query := "UPDATE roles SET name = ? description = ?, updated_at = NOW(), WHERE id = ?"

	_, err := rr.db.Exec(query, name, description, id)

	if err != nil {
		return nil, err
	}

	role := &models.Role{
		Id:          id,
		Name:        name,
		Description: description,
		CreatedAt:   "",
		UpdatedAt:   "",
	}

	return role, nil
}
