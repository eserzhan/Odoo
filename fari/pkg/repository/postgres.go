package repository

import (
	"fmt"
	

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Configs struct{
	Port string 
	Host string
	Dbname string 
	Sslmode string 
	Username string 
	Password string
}

func NewPostgresDB(c Configs) (*sqlx.DB, error){
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", c.Host,
	c.Port, c.Dbname, c.Username, c.Password, c.Sslmode)

	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}