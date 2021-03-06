package models

import (
	"errors"
)

type Post struct {
	ID          int    `json:"id" db:"id,omitempty"`
	Author      string `json:"author" db:"author"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

func (db *DB) AllPosts() ([]*Post, error) {
	var posts []*Post
	err := db.Collection("posts").Find().All(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (db *DB) Post(ID int) (*Post, error) {
	post := &Post{}
	err := db.Collection("posts").Find("id", ID).One(post)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (db *DB) UpdatePost(post *Post) (*Post, error) {
	err := db.Collection("posts").Find("id", post.ID).Update(post)
	if err != nil {
		return nil, err
	}

	return db.Post(post.ID)
}

func (db *DB) CreatePost(post *Post) (*Post, error) {
	ID, err := db.Collection("posts").Insert(post)
	if err != nil {
		return nil, err
	}

	if ID, ok := ID.(int64); ok {
		return db.Post(int(ID))
	}

	return nil, errors.New("Error getting ID of newly created element")
}

func (db *DB) DeletePost(ID int) (*Post, error) {
	post := &Post{}

	res := db.Collection("posts").Find("id", ID)
	err := res.One(post)
	if err != nil {
		return nil, err
	}

	err = res.Delete()
	if err != nil {
		return nil, err
	}

	return post, nil
}
