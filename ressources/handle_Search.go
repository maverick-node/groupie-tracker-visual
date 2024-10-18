package groupie

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// HandleSearch processes search requests and filters artists based on user input (search term).
func HandleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed)
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

	var filteredArtists []Artists

	search := r.FormValue("search")
	search = strings.TrimSpace(search)
	if search == "" {
		HandleError(w, http.StatusBadRequest)
		return
	}
	for _, artist := range artists {
		isFound := false
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(search)) {
			isFound = true
		}

		if strconv.Itoa(artist.CreationDate) == search {
			isFound = true
		}

		if strings.Contains(artist.FirstAlbum, search) {
			isFound = true
		}

		for i := 0; i < len(artist.Members); i++ {
			if strings.Contains(strings.ToLower(artist.Members[i]), strings.ToLower(search)) {
				isFound = true
			}
		}

		for _, location := range locations.Index {
			if location.Id == artist.Id {
				for _, locat := range location.Locations {
					if strings.Contains(strings.ToLower(locat), strings.ToLower(search)) {
						isFound = true
					}
				}
			}
		}

		for _, date := range dates.Index {
			if date.Id == artist.Id {
				for _, dat := range date.Dates {
					if strings.Contains(dat, search) {
						isFound = true
					}
				}
			}
		}
		if isFound {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	if len(filteredArtists) == 0 {
		HandleError(w, http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("templates/search_Filters.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, filteredArtists)
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
