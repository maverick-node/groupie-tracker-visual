package groupie

import (
	"net/http"
	"os"
	"path/filepath"
)
// fileExists checks if a file exists and returns an error if it does not.
func fileExists(filename string) error {
	_, err := os.Stat(filename)
	return err
}
// HandleStatic serves static files if they exist and have allowed extensions, or returns a 404 error.
func HandleStatic(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = path[len("/static/"):]
	allowedExtensions := map[string]bool{
		".css":   true,
		".png":   true,
		".jpg":   true,
		".jpeg":  true,
		".svg":   true,
		".woff2": true,
		".mp4": true,
		".ttf": true,
	}
	ext := filepath.Ext(path)
	if !allowedExtensions[ext] {
		HandleError(w, http.StatusNotFound)
		return
	}
	fullPath := filepath.Join("static", path)
	err := fileExists(fullPath)
	if err == nil {
		http.ServeFile(w, r, fullPath)
	} else {
		HandleError(w, http.StatusNotFound)
		return
	}
}