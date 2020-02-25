package main

import (
	"net/http"
	"strings"
	"log"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Route("/", func(root chi.Router) {
		fileServer(root, "", "/public/", http.Dir("../public/"))
		fileServer(root, "", "/", http.Dir("../public/views"))
	})

	log.Printf("Up and running on port 8080...")
	http.ListenAndServe(":8080", r)
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

	log.Printf("path: %s\n", path)

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
