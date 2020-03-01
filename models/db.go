package models

import (
	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

type Datastore interface {
	AllPosts() ([]Post, error)
}

type DB struct {
	db.Database
}

func NewDB(host, database, user, password string) (*DB, error) {
	settings := mysql.ConnectionURL{
		Host:     host,     // MySQL server IP or name.
		Database: database, // Database name.
		User:     user,     // Optional user name.
		Password: password, // Optional user password.
	}

	db, err := mysql.Open(settings)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetLogging(true)

	return &DB{db}, nil
}
