package controllers

import (
	"github.com/hlscalon/go-react-boilerplate/models"
)

var posts = []*models.Post{
	{ID: 1, Author: "hlscalon", Title: "this is the first post ever", Description: "this post is awewsome"},
	{ID: 2, Author: "batman", Title: "gotham needs you", Description: "I am out of business"},
	{ID: 3, Author: "robin", Title: "nooooo", Description: "Please Batman, don't go!!!"},
}

func GetAllPosts() []*models.Post {
	return posts
}
