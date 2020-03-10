package router

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

	if rec.Code != 200 || expected != rec.Body.String() {
		t.Errorf("\nExpected = %#v\nObtained = %#v\nExpected = %#v\nObtained = %#v", 200, rec.Code, expected, rec.Body.String())
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

	expected := "{\"id\":1,\"author\":\"hlscalon\",\"title\":\"hello\",\"description\":\"Hello, World!\"}\n"

	if rec.Code != 200 || expected != rec.Body.String() {
		t.Errorf("\nExpected = %#v\nObtained = %#v\nExpected = %#v\nObtained = %#v", 200, rec.Code, expected, rec.Body.String())
	}
}

func TestGetPostNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil) // api/public/v1/posts/1

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("postID", "10000")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()
	env := Env{db: &models.MockDB{}}
	env.postCtx(http.HandlerFunc(env.getPost)).ServeHTTP(rec, req)

	expected := "{\"status\":\"Resource not found.\"}\n"

	if rec.Code != 404 || expected != rec.Body.String() {
		t.Errorf("\nExpected = %#v\nObtained = %#v\nExpected = %#v\nObtained = %#v", 404, rec.Code, expected, rec.Body.String())
	}
}

func TestUpdatePost(t *testing.T) {
	byteData := `{"id":1,"author":"hlscalon","title":"where I come from","description":"Hello, World!"}`
	body := strings.NewReader(byteData)
	req, _ := http.NewRequest("PUT", "/", body) // api/admin/v1/posts/1
	req.Header.Set("Content-Type", "application/json")

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("postID", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()
	env := Env{db: &models.MockDB{}}
	env.postCtx(http.HandlerFunc(env.updatePost)).ServeHTTP(rec, req)

	expected := byteData + "\n"

	if rec.Code != 200 || expected != rec.Body.String() {
		t.Errorf("\nExpected = %#v\nObtained = %#v\nExpected = %#v\nObtained = %#v", 200, rec.Code, expected, rec.Body.String())
	}
}

func TestUpdatePostInvalidContentType(t *testing.T) {
	byteData := `{"id":1,"author":"hlscalon","title":"where I come from","description":"Hello, World!"}`
	body := strings.NewReader(byteData)
	req, _ := http.NewRequest("PUT", "/", body) // api/admin/v1/posts/1

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("postID", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()
	env := Env{db: &models.MockDB{}}
	env.postCtx(http.HandlerFunc(env.updatePost)).ServeHTTP(rec, req)

	expected := "{\"status\":\"Invalid request.\",\"error\":\"render: unable to automatically decode the request content type\"}\n"

	if rec.Code != 400 || expected != rec.Body.String() {
		t.Errorf("\nExpected = %#v\nObtained = %#v\nExpected = %#v\nObtained = %#v", 400, rec.Code, expected, rec.Body.String())
	}
}

func TestCreatePost(t *testing.T) {
	byteData := `{"author":"batman","title":"Gotham I'm coming!","description":"All criminals in town, be aware!"}`
	body := strings.NewReader(byteData)
	req, _ := http.NewRequest("POST", "/", body) // api/admin/v1/posts/
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	env := Env{db: &models.MockDB{}}
	http.HandlerFunc(env.createPost).ServeHTTP(rec, req)

	expected := "{\"id\":4,\"author\":\"batman\",\"title\":\"Gotham I'm coming!\",\"description\":\"All criminals in town, be aware!\"}\n"

	if rec.Code != 201 || expected != rec.Body.String() {
		t.Errorf("\nExpected = %#v\nObtained = %#v\nExpected = %#v\nObtained = %#v", 201, rec.Code, expected, rec.Body.String())
	}
}

func TestDeletePost(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/", nil) // api/admin/v1/posts/4

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("postID", "4")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()
	env := Env{db: &models.MockDB{}}
	env.postCtx(http.HandlerFunc(env.deletePost)).ServeHTTP(rec, req)

	expected := "{\"id\":4,\"author\":\"batman\",\"title\":\"Gotham I'm coming!\",\"description\":\"All criminals in town, be aware!\"}\n"

	if rec.Code != 200 || expected != rec.Body.String() {
		t.Errorf("\nExpected = %#v\nObtained = %#v\nExpected = %#v\nObtained = %#v", 200, rec.Code, expected, rec.Body.String())
	}
}
