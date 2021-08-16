package repositories

import (
	"database/sql"
)

type userRepository struct{
	db *sql.DB
}

func UserRepository(db *sql.DB) *userRepository{
	return &userRepository{db}
}