package groupie

import (
	"bytes"
	"html/template"
	"net/http"
)

// HandleHome serves the home page by fetching a list of artists from an API and rendering it with a template, handling errors appropriately.
func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {

		HandleError(w, http.StatusNotFound)
		return
	}
	var artists []Artists
	err := fetch("https://groupietrackers.herokuapp.com/api/", "artists", &artists)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
	var dates Dates2
	if err := fetch("https://groupietrackers.herokuapp.com/api/", "dates", &dates); err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
	var locations Locations2
	if err := fetch("https://groupietrackers.herokuapp.com/api/", "locations", &locations); err != nil {

		HandleError(w, http.StatusInternalServerError)
		return
	}

	var LocationsOptions []string

	for _, loc := range locations.Index {
		for _, location := range loc.Locations {
			option := location
			LocationsOptions = append(LocationsOptions, option)
		}
	}
	var DatesOptions []string
	for _, dat := range dates.Index {
		for _, dates := range dat.Dates {
			option := dates
			DatesOptions = append(DatesOptions, option)
		}
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	type Data struct {
		OptionsLocations []string
		OptionsDates     []string
		Artists          []Artists
		Locations        []Locations
	}
	data := Data{
		OptionsLocations: LocationsOptions,
		OptionsDates:     DatesOptions,
		Artists:          artists,
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
}
