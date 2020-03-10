package router

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/hlscalon/go-react-boilerplate/controllers"
	"github.com/hlscalon/go-react-boilerplate/models"
	"github.com/hlscalon/go-react-boilerplate/utils"
)

type PostResponse struct {
	*models.Post
}

type PostRequest struct {
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

func (pr *PostRequest) Bind(r *http.Request) error {
	if pr.Post == nil {
		return errors.New("Missing required Post fields.")
	}

	return nil
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

// postCtx middleware is used to load a Post object from
// the URL parameters passed through as the request. In case
// the Post could not be found, we stop here and return a 404.
func (env *Env) postCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var post *models.Post
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

		ctx := context.WithValue(r.Context(), "post", post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (env *Env) getPost(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value("post").(*models.Post)

	if err := render.Render(w, r, newPostResponse(post)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (env *Env) updatePost(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value("post").(*models.Post)

	data := &PostRequest{post}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	post, err := controllers.UpdatePost(env.db, data.Post)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	if err := render.Render(w, r, newPostResponse(post)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (env *Env) deletePost(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value("post").(*models.Post)

	post, err := controllers.DeletePost(env.db, post.ID)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	if err := render.Render(w, r, newPostResponse(post)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (env *Env) createPost(w http.ResponseWriter, r *http.Request) {
	data := &PostRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	post, err := controllers.CreatePost(env.db, data.Post)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := render.Render(w, r, newPostResponse(post)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
