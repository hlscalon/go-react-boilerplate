package controllers

import (
	"github.com/hlscalon/go-react-boilerplate/models"
)

func AllPosts(db models.Datastore) ([]models.Post, error) {
	return db.AllPosts()
}
