package router

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/hlscalon/go-react-boilerplate/models"
	"github.com/hlscalon/go-react-boilerplate/controllers"
)

type PostResponse struct {
	*models.Post
}

func (pr *PostResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func newPostResponse(post *models.Post) *PostResponse {
	resp := &PostResponse{post}
	return resp
}

func newPostListResponse(posts []*models.Post) []render.Renderer {
	list := []render.Renderer{}
	for _, post := range posts {
		list = append(list, newPostResponse(post))
	}
	return list
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	posts := controllers.GetAllPosts()
	if err := render.RenderList(w, r, newPostListResponse(posts)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
