package main

import (
	"go-learn-blogs/config"
	"go-learn-blogs/controllers/base"
	"go-learn-blogs/controllers/blogscontroller"
	"go-learn-blogs/controllers/homecontroller"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var r *chi.Mux

func main() {

	r = chi.NewRouter()
	r.Use(middleware.Recoverer)

	var err error
	config.ConnectDB()

	r.Use(base.ChangeMethod)
	r.Get("/", homecontroller.Index)
	r.Post("/upload", base.UploadHandler)
	r.Get("/public/images/*", base.ServeImages)
	r.Route("/blogs", func(r chi.Router) {
		r.Get("/", blogscontroller.GetAllBlogs)
		r.Get("/create", blogscontroller.CreateBlogs)
		r.Post("/", blogscontroller.CreateBlogs)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(blogscontroller.BlogCtx)
			r.Get("/", blogscontroller.GetBlog)       // GET /articles/1234
			r.Post("/", blogscontroller.EditBlog)     // PUT /articles/1234
			r.Delete("/", blogscontroller.DeleteBlog) // DELETE /articles/1234
			r.Get("/edit", blogscontroller.EditBlog)  // GET /articles/1234/edit
		})
	})

	log.Println("Server started on: http://localhost:8018")
	err = http.ListenAndServe(":8018", r)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
