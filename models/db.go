package models

import (
	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

type Datastore interface {
	AllPosts() ([]*Post, error)
}

type DB struct {
	db.Database
}

func NewDB() (*DB, error) {
	settings := mysql.ConnectionURL{
		Host:     "localhost",  // MySQL server IP or name.
		Database: "go_react_boilerplate",    // Database name.
		User:     "root",     // Optional user name.
		Password: "",     // Optional user password.
	}

	sess, err := mysql.Open(settings)
	if err != nil {
		return nil, err
	}

	if err := sess.Ping(); err != nil {
		return nil, err
	}

	return &DB{sess}, nil
}
