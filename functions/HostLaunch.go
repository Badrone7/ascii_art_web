package functions

import (
	"fmt"
	"log"
	"net/http"
)

type Artstr struct {
	text  string
	style string
	Art   string
}

func HostLauncher() {
	var Art Artstr
	fs := http.FileServer(http.Dir("page"))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		fs.ServeHTTP(w, r)
	}))
	http.HandleFunc("/ascii-art", ArtHandler(Art))
	fmt.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
