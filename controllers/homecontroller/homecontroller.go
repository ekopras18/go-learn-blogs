package homecontroller

import (
	"go-learn-blogs/controllers/base"
	"go-learn-blogs/models"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	result, err := models.Get()
	base.CatchWithMessage(err, "Error getting all articles")
	data := map[string]any{
		"blogs":       result,
		"title":       "Welcome to my blog",
		"page_tittle": "List of Blogs.",
		"page_active": "blog",
	}

	t, _ := template.ParseFiles("views/layout/base.html", "views/index.html")
	err = t.Execute(w, data)
	base.CatchWithMessage(err, "Error executing template")

}
