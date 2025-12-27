package functions

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Artstr struct to hold art data
type Artstr struct {
	text  string
	style string
	Art   string
}

func PageChecker(w http.ResponseWriter, r *http.Request, Art Artstr) bool {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		tmpl400.Execute(w, nil)
		return false
	}
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		w.WriteHeader(http.StatusNotFound)
		tmpl404.Execute(w, nil)
		return false
	}
	return true
}

// HostLauncher starts the web server
func HostLauncher() {
	var Art Artstr
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := os.Stat("static" + strings.TrimPrefix(r.URL.Path, "/static"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			tmpl404.Execute(w, nil)
			return
		}
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if PageChecker(w, r, Art) {
			Art.Art = ""
			tmpl.Execute(w, Art)
		}
	})
	http.HandleFunc("/ascii-art", ArtHandler(Art))
	fmt.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
