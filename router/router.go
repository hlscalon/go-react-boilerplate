package router

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/hlscalon/go-react-boilerplate/models"
)

type Env struct {
	db models.Datastore
}

type Router struct {
	router *chi.Mux
	env *Env
}

func New() (*Router, error) {
	db, err := models.NewDB()
	if err != nil {
		return nil, err
	}

	r := chi.NewRouter()
	env := &Env{db}

	return &Router{r, env}, nil
}

func (r *Router) Init(port string) {
	router := r.router
	env := r.env

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/api/public/v1", func(r chi.Router) {
		// config to all public routes

		// specific routes
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", env.listPosts)
			// r.Post("/", createPost)

			// r.Route("/{postID}", func(r chi.Router) {
			// 	r.Use(postCtx)
			// 	r.Get("/", getPost)
			// 	r.Put("/", updatePost)
			// 	r.Delete("/", deletePost)
			// })
		})
	})

	// r.Route("/api/admin/v1", func(r chi.Router) {})

	router.Route("/", func(root chi.Router) {
		fileServer(root, "", "/dist/", http.Dir("assets/public/dist/"))
		fileServer(root, "", "/", http.Dir("assets/public/static/"))
	})

	log.Printf("Up and running on port %s...", port)
	http.ListenAndServe(":"+port, r.router)
}

func fileServer(r chi.Router, basePath string, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("fileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(basePath+path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
