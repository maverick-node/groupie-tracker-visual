package groupie


// Artists represents the structure of artist data, including their ID, image, name, members, creation date, and first album.
type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// Locations represents the structure of location data, including an ID and a list of locations.
type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

// Dates represents the structure of tour date data, including an ID and a list of dates.
type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Relations represents the structure of relationships between dates and locations, including an ID and a mapping of dates to their respective locations.
type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Locations2 represents the structure of Index data, including an ID and a list of locations.
type Locations2 struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

// Dates2 represents the structure of Index data, including an ID and a list of dates.
type Dates2 struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}
