//
// Package: router
//
// All routes on our program are defined in this file. The functions to interact with are defined in other files, like 'post.go'
// It uses chi as a router (a layer above net/http), and some of it's middlewares
//
// Our frontend assets are served here as well, ie. dist/ and images
//

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

// Initialize all routes of our program
// Start server on port and use initialized database connection on our environment
func Init(db models.Datastore, port string) {
	r := chi.NewRouter()
	env := &Env{db}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/public/v1", func(r chi.Router) {
		// config to all public routes

		// specific routes
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", env.listPosts)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(env.postCtx)
				r.Get("/", env.getPost)
			})
		})
	})

	r.Route("/api/admin/v1", func(r chi.Router) {
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", env.listPosts)
			r.Post("/", env.createPost)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(env.postCtx)
				r.Get("/", env.getPost)
				r.Put("/", env.updatePost)
				r.Delete("/", env.deletePost)
			})
		})
	})

	r.Route("/", func(root chi.Router) {
		fileServer(root, "", "/dist/", http.Dir("assets/public/dist/"))
		fileServer(root, "", "/", http.Dir("assets/public/static/"))
	})

	log.Printf("Up and running on port %s...", port)
	http.ListenAndServe(":"+port, r)
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
