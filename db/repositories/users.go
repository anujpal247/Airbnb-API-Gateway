package db

import (
	"AuthApp/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create(username string, email string, password string) error
	GetById() (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) Create(username string, email string, password string) error {
	query := "INSERT INTO users(username, email, password) VALUES(?,?,?)"
	res, err := u.db.Exec(query, username, email, password)

	if err != nil {
		fmt.Println("Error inserting user", err)
		return err
	}

	rowsAffected, rowErr := res.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error row affected ", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil
	}

	fmt.Println("User created successfully, rows affected:", rowsAffected)

	return nil
}

func (u *UserRepositoryImpl) GetById() (*models.User, error) {
	fmt.Println("Getting user by id")

	// step 1. prepare query
	query := "SELECT id, username, email FROM users WHERE id = ?"

	// step 2. execute the query
	row := u.db.QueryRow(query, 1)

	// step 3. process the result
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no user found by given id")
			return nil, err
		} else {
			fmt.Println("Error scaning user")
			return nil, err
		}
	}

	// print the user details

	fmt.Println("user found", user)
	return user, nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {

	query := "SELECT id, username, password, email, created_at FROM users WHERE email=?"

	row := u.db.QueryRow(query, email)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found by given email")
			return nil, err
		} else {
			fmt.Println("Error scaning user")
			return nil, err
		}
	}

	fmt.Println("user found", user)
	return user, nil
}
