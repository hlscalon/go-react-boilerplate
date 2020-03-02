package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/hlscalon/go-react-boilerplate/controllers"
	"github.com/hlscalon/go-react-boilerplate/models"
	"github.com/hlscalon/go-react-boilerplate/utils"
)

type PostResponse struct {
	models.Post
}

func (pr *PostResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func newPostResponse(post models.Post) *PostResponse {
	resp := &PostResponse{post}
	return resp
}

func newPostListResponse(posts []models.Post) []render.Renderer {
	list := []render.Renderer{}
	for _, post := range posts {
		list = append(list, newPostResponse(post))
	}
	return list
}

func (env *Env) listPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := controllers.AllPosts(env.db)
	if err != nil {
		render.Render(w, r, ErrInternal(err))
		return
	}

	if err := render.RenderList(w, r, newPostListResponse(posts)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (env *Env) getPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	var err error

	if postID := chi.URLParam(r, "postID"); postID != "" {
		var id int

		if id, err = utils.StrToInt(postID); err == nil {
			post, err = controllers.Post(env.db, id)
		}
	} else {
		render.Render(w, r, ErrNotFound)
		return
	}

	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	if err := render.Render(w, r, newPostResponse(post)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
