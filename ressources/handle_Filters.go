package groupie

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func HandleFilters(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}

	creationDateEnd := r.FormValue("creationDateEnd")
	creationDateStart := r.FormValue("creationDateStart")
	var endDateConverted int
	var startDateConverted int
	var err error
	var err1 error

	if creationDateEnd != "" && creationDateStart != "" {
		endDateConverted, err = strconv.Atoi(creationDateEnd)
		startDateConverted, err1 = strconv.Atoi(creationDateStart)
		if err != nil && err1 != nil {
			HandleError(w, http.StatusInternalServerError)
			return
		}
	} else {
		HandleError(w, http.StatusBadRequest)
		return
	}

	firstAlbumDateStart := r.FormValue("firstAlbumDateStart")
	firstAlbumDateEnd := r.FormValue("firstAlbumDateEnd")

	if firstAlbumDateEnd == "" || firstAlbumDateStart == "" {
		HandleError(w, http.StatusBadRequest)
		return
	}

	numOfMembers := r.URL.Query()["Member"]
	concertsLocations := strings.ToLower(r.FormValue("filter"))
	concertsLocations = strings.ReplaceAll(concertsLocations, ", ", "-")

	var artists []Artists

	if err := fetch("https://groupietrackers.herokuapp.com/api/", "artists", &artists); err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	var locations Locations2

	if err := fetch("https://groupietrackers.herokuapp.com/api/", "locations", &locations); err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	var res []Artists

	for _, artist := range artists {
		isFound := true

		if !(endDateConverted >= artist.CreationDate && startDateConverted <= artist.CreationDate) {
			isFound = false
			continue
		}

		if firstAlbumDateStart != "" && firstAlbumDateEnd != "" {
			if !(firstAlbumDateEnd >= artist.FirstAlbum[6:] && firstAlbumDateStart <= artist.FirstAlbum[6:]) {
				isFound = false
				continue
			}
		}

		if len(numOfMembers) > 0 {
			memberMatch := false
			for _, v := range numOfMembers {
				n, _ := strconv.Atoi(v)
				if len(artist.Members) == n {
					memberMatch = true
					break
				}
			}
			if !memberMatch {
				isFound = false
				continue
			}
		}

		if concertsLocations != "" {
			locationMatch := false
			for _, location := range locations.Index {
				if location.Id == artist.Id {
					for _, locat := range location.Locations {
						if strings.Contains(strings.ToLower(locat), concertsLocations) {
							locationMatch = true
							break
						}
					}
				}
			}
			if !locationMatch {
				isFound = false
				continue
			}
		}

		if isFound {
			res = append(res, artist)
		}
	}

	tmpl, err := template.ParseFiles("templates/search_Filters.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, res)
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
