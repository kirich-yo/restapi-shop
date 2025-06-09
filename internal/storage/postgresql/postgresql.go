package postgresql

import (
	"fmt"
	"database/sql"
	"restapi-sportshop/internal/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func New(dbConn *config.DatabaseConn) (*Storage, error) {
	const op = "internal.storage.New"
	postgresURI := fmt.Sprintf("postgres://%s:%s@%s/sportshop?sslmode=disable", dbConn.Login, dbConn.Password, dbConn.Address)
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db}, nil
}
