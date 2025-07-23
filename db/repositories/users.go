package db

import (
	"AuthApp/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create() error
	GetById() (*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) Create() error {
	fmt.Println("Creating user in user repository")
	query := "INSERT INTO users(username, email, password) VALUES(?,?,?)"
	res, err := u.db.Exec(query, "testuser", "test@example.com", "123456")

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
