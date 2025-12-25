package functions

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Artstr struct to hold art data
type Artstr struct {
	text  string
	style string
	Art   string
}

func GetMethodChecker(w http.ResponseWriter, r *http.Request, Art Artstr) bool {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		tmpl400.Execute(w, nil)
		return false
	}
	if r.URL.Path != "/" && !strings.HasPrefix(r.URL.Path, "/dependencies") && !strings.HasPrefix(r.URL.Path, "/static") && !strings.HasPrefix(r.URL.Path, "/styles") {
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
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if GetMethodChecker(w, r, Art) {
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
