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

func (db *DB) Post(ID int) (Post, error) {
	var post Post
	err := db.Collection("posts").Find("id", ID).One(&post)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (db *DB) UpdatePost(post Post) (Post, error) {
	err := db.Collection("posts").Find("id", post.ID).Update(post)
	if err != nil {
		return Post{}, err
	}

	return db.Post(post.ID)
}
