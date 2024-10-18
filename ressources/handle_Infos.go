package groupie

import (
	"bytes"
	"html/template"
	"net/http"
)

// HandleInfos fetches and displays detailed information about an artist, including their locations, tour dates, and related data, handling errors as needed.
func HandleInfos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")

	var artist Artists
	var locations Locations
	var dates Dates
	var relations Relations

	if err := fetch("https://groupietrackers.herokuapp.com/api/", "artists/"+id, &artist); err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	if err := fetch("https://groupietrackers.herokuapp.com/api/", "locations/"+id, &locations); err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	if err := fetch("https://groupietrackers.herokuapp.com/api/", "dates/"+id, &dates); err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	if err := fetch("https://groupietrackers.herokuapp.com/api/", "relation/"+id, &relations); err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/infos.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	if artist.Id == 0 {
		HandleError(w, http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"artist":         artist,
		"locations":      locations.Locations,
		"dates":          dates.Dates,
		"DatesLocations": relations.DatesLocations,
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
