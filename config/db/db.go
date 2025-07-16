package config

import (
	env "AuthApp/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()

	cfg.User = env.GetString("DB_USER", "root")
	cfg.Addr = env.GetString("DB_ADDR", "localhost:3306")
	cfg.Passwd = env.GetString("DB_PASSWORD", "root")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.DBName = env.GetString("DB_NAME", "auth_app")

	fmt.Println("Connecting to db", cfg.FormatDSN())
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error connecting to db", err)
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error pinging to db", pingErr)
		return nil, pingErr
	}
	fmt.Println("connected to db successfully", cfg.DBName)

	return db, nil
}
