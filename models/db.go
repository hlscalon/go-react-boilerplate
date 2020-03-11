package models

import (
	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

// Interface to our models access database or to be mocked
type Datastore interface {
	AllPosts() ([]*Post, error)
	Post(int) (*Post, error)
	UpdatePost(*Post) (*Post, error)
	CreatePost(*Post) (*Post, error)
	DeletePost(int) (*Post, error)
}

type DB struct {
	db.Database
}

// Open MySQL connection and check if it is working
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
