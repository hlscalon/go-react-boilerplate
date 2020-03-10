package models

import (
	"errors"
)

type MockDB struct{}

var postsMockDB = []*Post{
	&Post{1, "hlscalon", "hello", "Hello, World!"},
	&Post{2, "user", "summer", "Summer is ending!"},
	&Post{3, "tester", "day", "Today is a sunny day"},
}

func (mdb *MockDB) AllPosts() ([]*Post, error) {
	return postsMockDB, nil
}

func (mdb *MockDB) Post(ID int) (*Post, error) {
	for _, p := range postsMockDB {
		if p.ID == ID {
			return p, nil
		}
	}

	return nil, errors.New("Resource not found")
}

func (mdb *MockDB) UpdatePost(post *Post) (*Post, error) {
	for _, p := range postsMockDB {
		if p.ID == post.ID {
			p = post

			return post, nil
		}
	}

	return nil, errors.New("Resource not found")
}

func (mdb *MockDB) CreatePost(post *Post) (*Post, error) {
	lastID := postsMockDB[len(postsMockDB)-1 : len(postsMockDB)][0].ID

	post.ID = lastID + 1
	postsMockDB = append(postsMockDB, post)

	return post, nil
}

func (mdb *MockDB) DeletePost(ID int) (*Post, error) {
	var posts []*Post
	var post *Post

	for _, p := range postsMockDB {
		if p.ID == ID {
			post = p
		} else {
			posts = append(posts, p)
		}
	}

	postsMockDB = posts
	if post != nil {
		return post, nil
	}

	return nil, errors.New("Resource not found")
}
