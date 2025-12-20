package functions

import (
	"html/template"
	"net/http"
)

func IsValidInput(style string) bool {
	validStyles := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}
	return validStyles[style]
}

func ArtHandler(art Artstr) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		tmpl := template.Must(template.ParseFiles("page/index.html"))
		art.text = r.FormValue("text")
		art.style = r.FormValue("banner")
		if !IsValidInput(art.style) {
			art.Art = "Invalid style selected."
			w.Header().Set("Content-Type", "text/html")
			tmpl.Execute(w, art)
			return
		}
		result, err := ArtMaker(art.text, art.style)
		if err != nil {
			http.Error(w, "Error generating art", http.StatusInternalServerError)
			return
		}
		if result == nil {
			w.Header().Set("Content-Type", "text/html")
			tmpl.Execute(w, art)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		art.Art = string(result)
		tmpl.Execute(w, art)
	}
}
