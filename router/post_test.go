package router

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/hlscalon/go-react-boilerplate/models"
)

type mockDB struct{}

func (mdb *mockDB) AllPosts() ([]models.Post, error) {
    posts := make([]models.Post, 0)
    posts = append(posts, models.Post{1, "hlscalon", "hello", "Hello, World!"})
    posts = append(posts, models.Post{2, "user", "summer", "Summer is ending!"})
    posts = append(posts, models.Post{3, "tester", "day", "Today is a sunny day"})
    return posts, nil
}

func (mdb *mockDB) Post(ID int) (models.Post, error) {
    return models.Post{}, nil
}

func (mdb *mockDB) UpdatePost(post models.Post) (models.Post, error) {
    return models.Post{}, nil
}

func (mdb *mockDB) CreatePost(post models.Post) (models.Post, error) {
    return models.Post{}, nil
}

func (mdb *mockDB) DeletePost(ID int) (models.Post, error) {
    return models.Post{}, nil
}

func TestListPosts(t *testing.T) {
    rec := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/posts", nil)

    env := Env{db: &mockDB{}}
    http.HandlerFunc(env.listPosts).ServeHTTP(rec, req)

    expected :=
        "[" +
            "{" +
                "\"id\":1,\"author\":\"hlscalon\",\"title\":\"hello\",\"description\":\"Hello, World!\"" +
            "}," +
            "{" +
                "\"id\":2,\"author\":\"user\",\"title\":\"summer\",\"description\":\"Summer is ending!\"" +
            "}," +
            "{" +
                "\"id\":3,\"author\":\"tester\",\"title\":\"day\",\"description\":\"Today is a sunny day\"" +
            "}" +
        "]" +
        "\n"

    if expected != rec.Body.String() {
        t.Errorf("\nExpected = %#v\nObtained = %#v", expected, rec.Body.String())
    }
}
