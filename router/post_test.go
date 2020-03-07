package router

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "context"

    "github.com/go-chi/chi"

    "github.com/hlscalon/go-react-boilerplate/models"
)

func TestListPosts(t *testing.T) {
    rec := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/", nil) // api/public/v1/posts

    env := Env{db: &models.MockDB{}}
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

func TestGetPost(t *testing.T) {
    req, _ := http.NewRequest("GET", "/", nil) // api/public/v1/posts/1

    rctx := chi.NewRouteContext()
    rctx.URLParams.Add("postID", "1")

    req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

    rec := httptest.NewRecorder()
    env := Env{db: &models.MockDB{}}
    env.postCtx(http.HandlerFunc(env.getPost)).ServeHTTP(rec, req)

    expected :=
        "{" +
            "\"id\":1,\"author\":\"hlscalon\",\"title\":\"hello\",\"description\":\"Hello, World!\"" +
        "}" +
        "\n"

    if expected != rec.Body.String() {
        t.Errorf("\nExpected = %#v\nObtained = %#v", expected, rec.Body.String())
    }
}
