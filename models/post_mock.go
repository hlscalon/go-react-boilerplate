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
    return &Post{}, nil
}

func (mdb *MockDB) CreatePost(post *Post) (*Post, error) {
    return &Post{}, nil
}

func (mdb *MockDB) DeletePost(ID int) (*Post, error) {
    return &Post{}, nil
}

