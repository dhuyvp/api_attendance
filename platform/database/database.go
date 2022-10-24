package database

import (
	"attendance_user/pkg/utils"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectionDB() (*sqlx.DB, error) {
	dbURL, err := utils.URLBuilder("mysql")
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("mysql", dbURL)
	if err != nil {
		return nil, fmt.Errorf("Error, not connected to database, %w", err)
	}
	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("Error, not connected to database, %w", err)

	}

	return db, nil
}
