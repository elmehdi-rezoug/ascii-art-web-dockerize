package handlers

import (
	"ascii-art-web/asciigenerator"
	"bytes"
	"html/template"
	"net/http"
	"strings"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	input := r.FormValue("input")
	banner := r.FormValue("banner")

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		ErrorHandler(w, "400 wrong banner", http.StatusBadRequest)
		return
	}

	_, ok := r.Form["input"]
	if !ok {
		ErrorHandler(w, "400 must provide an input", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	output, err := asciigenerator.GenerateAsciiArt(input, banner)
	if err != nil {
		if strings.HasPrefix(err.Error(), "400") {
			ErrorHandler(w, err.Error(), http.StatusBadRequest)
		} else {
			ErrorHandler(w, "500 internal server error", http.StatusInternalServerError)
		}
		return
	}

	if strings.HasPrefix(output, "\n") {
		output = "\n" + output
	}

	data := Data{
		Output: output,
	}

	var buff bytes.Buffer
	err = tmpl.Execute(&buff, data)
	if err != nil {
		ErrorHandler(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(buff.Bytes())
}
