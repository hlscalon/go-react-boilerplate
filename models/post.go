package models

type Post struct {
	ID          int    `json:"id" db:"id,omitempty"`
	Author      string `json:"author" db:"author"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

func (db *DB) AllPosts() ([]Post, error) {
	var posts []Post
	err := db.Collection("posts").Find().All(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
