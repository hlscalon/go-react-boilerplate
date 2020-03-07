package controllers

import (
	"github.com/hlscalon/go-react-boilerplate/models"
)

func AllPosts(db models.Datastore) ([]models.Post, error) {
	return db.AllPosts()
}

func Post(db models.Datastore, ID int) (models.Post, error) {
	return db.Post(ID)
}

func UpdatePost(db models.Datastore, post models.Post) (models.Post, error) {
	return db.UpdatePost(post)
}

func CreatePost(db models.Datastore, post models.Post) (models.Post, error) {
	return db.CreatePost(post)
}

func DeletePost(db models.Datastore, ID int) (models.Post, error) {
	return db.DeletePost(ID)
}
