package groupie

import (
	"html/template"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	// Assuming `data` is the data you want to pass to the template
	err = tmpl.Execute(w, nil) // Passing `nil` if no dynamic data is required
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
	}

}
