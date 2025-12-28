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

func TemplatesHandler(err int, art Artstr, w http.ResponseWriter) {
	var template *template.Template
	switch err {
	case 400:
		template = tmpl400
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/html")
		template.Execute(w, nil)
	case 404:
		template = tmpl404
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/html")
		template.Execute(w, nil)
	case 500:
		template = tmpl500
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		template.Execute(w, nil)
	case 599:
		template = tmpl
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		template.Execute(w, art)
	default:
		template = tmpl
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		template.Execute(w, art)
	}
}

func ArtHandler(art Artstr) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			TemplatesHandler(400, art, w)
			return
		}
		err := r.ParseForm()
		if err != nil {
			TemplatesHandler(400, art, w)
			return
		}
		art.text = r.FormValue("text")
		if len(art.text) > 10000 {
			art.Art = "Input text is too long."
			TemplatesHandler(599, art, w)
			return
		}
		art.style = r.FormValue("banner")
		if !IsValidInput(art.style) {
			art.Art = "Invalid style selected."
			TemplatesHandler(599, art, w)
			return
		}
		result, errart, i := ArtMaker(art.text, art.style)
		if errart != nil {
			TemplatesHandler(500, art, w)
			return
		} else if i == 1 {
			TemplatesHandler(599, art, w)
			return
		}
		art.Art = string(result)
		TemplatesHandler(200, art, w)
	}
}
