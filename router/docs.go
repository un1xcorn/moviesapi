package router

import (
	"net/http"
)

func DocsHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := GetTemplate("./templates/docs.html")
	tmpl.ExecuteTemplate(w, "docs.html", nil)

}
