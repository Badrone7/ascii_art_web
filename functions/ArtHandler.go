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

var (
	tmpl404 = template.Must(template.ParseFiles("static/404.html"))
	tmpl400 = template.Must(template.ParseFiles("static/400.html"))
	tmpl500 = template.Must(template.ParseFiles("static/500.html"))
	tmpl    = template.Must(template.ParseFiles("static/index.html"))
)

func ArtHandler(art Artstr) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			tmpl400.Execute(w, nil)
			return
		}
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			tmpl400.Execute(w, nil)
			return
		}
		art.text = r.FormValue("text")
		art.style = r.FormValue("banner")
		if !IsValidInput(art.style) {
			art.Art = "Invalid style selected."
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.Execute(w, art)
			return
		}
		result, errart, i := ArtMaker(art.text, art.style)
		if errart != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tmpl500.Execute(w, nil)
			return
		} else if i == 1 {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text/html")
		art.Art = string(result)
		tmpl.Execute(w, art)
	}
}
