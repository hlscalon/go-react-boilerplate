package models

type Post struct {
	ID          int    `json:"id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (db *DB) AllPosts() ([]*Post, error) {
	return nil, nil
}
