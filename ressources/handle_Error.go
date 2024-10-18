package groupie

import (
	"html/template"
	"net/http"
)

// HandleError serves an error page based on the given status code or returns a generic 500 error.
func HandleError(w http.ResponseWriter, StatusCodes int) {
	templates := map[int]string{
		400: "templates/400.html",
		404: "templates/404.html",
		500: "templates/500.html",
	}
	tmplName := templates[StatusCodes]
	tmpl, err := template.ParseFiles(tmplName)
	if err != nil {
		tmpl, err = template.ParseFiles("templates/500.html")
		if err != nil {
			http.Error(w, "internal server error 500", http.StatusInternalServerError)
			return
		}

	}
	w.WriteHeader(StatusCodes)
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
		return
	}
}
