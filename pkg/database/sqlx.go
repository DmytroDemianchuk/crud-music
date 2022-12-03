package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func CreateConn(host, port, username, password, dbname string, secure bool) (*sqlx.DB, error) {
	sslmode := "disable"
	if secure {
		sslmode = "enable"
	}

	fqdn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, dbname, sslmode)

	db, err := sqlx.Connect("postgres", fqdn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
