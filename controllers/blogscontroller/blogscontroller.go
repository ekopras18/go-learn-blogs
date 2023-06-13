package blogscontroller

import (
	"context"
	"fmt"
	"go-learn-blogs/controllers/base"
	"go-learn-blogs/entities"
	"go-learn-blogs/models"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/go-chi/chi"
)

func BlogCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		result, err := models.Show(id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "blog", result)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	result, err := models.Get()
	base.CatchWithMessage(err, "Error getting all articles")
	data := map[string]any{
		"blogs":       result,
		"title":       "Blogs",
		"page_tittle": "List of Blogs.",
		"page_active": "blog",
	}

	t, _ := template.ParseFiles("views/layout/base.html", "views/blogs/index.html")
	err = t.Execute(w, data)
	base.CatchWithMessage(err, "Error executing template")

}

func CreateBlogs(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		data := map[string]any{
			"title":       "Post Blog",
			"page_tittle": "Create Blog",
			"page_active": "blog",
		}

		t, _ := template.ParseFiles("views/layout/base.html", "views/blogs/create.html")
		err := t.Execute(w, data)
		base.CatchWithMessage(err, "Error executing template")
	}

	if r.Method == "POST" {
		blog := &entities.Blogs{
			Title:     r.FormValue("title"),
			Author:    r.FormValue("author"),
			Tags:      r.FormValue("tags"),
			Content:   []byte(r.FormValue("content")),
			CreatedAt: time.Now(),
		}
		err := models.Store(blog)
		base.CatchWithMessage(err, "Error creating Blog")
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	blog := r.Context().Value("blog").(*entities.Blogs)
	blogs, _ := models.Get()
	data := map[string]any{
		"blog":        blog,
		"blogs":       blogs,
		"title":       "Blogs",
		"page_tittle": "List of Blogs.",
		"page_active": "blog",
	}

	t, _ := template.ParseFiles("views/layout/base.html", "views/blogs/show.html")
	err := t.Execute(w, data)
	base.CatchWithMessage(err, "Error Get Blog")
}

func EditBlog(w http.ResponseWriter, r *http.Request) {
	blog := r.Context().Value("blog").(*entities.Blogs)
	if r.Method == "GET" {

		data := map[string]any{
			"blog":        blog,
			"title":       "Blogs",
			"page_tittle": "Edit Blogs.",
			"page_active": "blog",
		}
		t, _ := template.ParseFiles("views/layout/base.html", "views/blogs/edit.html")
		err := t.Execute(w, data)
		base.CatchWithMessage(err, "Error Get Blog")
	}

	if r.Method == "POST" {
		newBlog := &entities.Blogs{
			Title:     r.FormValue("title"),
			Author:    r.FormValue("author"),
			Tags:      r.FormValue("tags"),
			Content:   []byte(r.FormValue("content")),
			UpdatedAt: time.Now(),
		}

		err := models.Update(strconv.Itoa(blog.Id), newBlog)
		base.CatchWithMessage(err, "Error Get Blog")
		http.Redirect(w, r, "/blogs", http.StatusFound)
	}
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	blog := r.Context().Value("blog").(*entities.Blogs)
	err := models.Delete(strconv.Itoa(blog.Id))
	base.CatchWithMessage(err, "Error deleting Blog")
	http.Redirect(w, r, "/blogs", http.StatusFound)
}
