package repository

import (
	"fmt"

	"github.com/eserzhan/rest"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user todo.User) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(username, password string) (int, error){
	var id int
	query := fmt.Sprintf("SELECT id FROM %s where username = $1 and password_hash = $2", usersTable)

	err := r.db.Get(&id, query, username, password)

	if err != nil {
		return -1, err
	}

	return id, nil
}