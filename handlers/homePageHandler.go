package handlers

import (
	"bytes"
	"html/template"
	"net/http"
)

type Data struct {
	Output string
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	var buff bytes.Buffer
	err = tmpl.Execute(&buff, nil)
	if err != nil {
		ErrorHandler(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(buff.Bytes())
}
