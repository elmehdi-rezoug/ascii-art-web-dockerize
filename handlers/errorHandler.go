package handlers

import (
	"bytes"
	"html/template"
	"net/http"
)

type ErrorData struct {
	Error string
}

func ErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, message, statusCode)
		return
	}
	errorData := ErrorData{
		Error: message,
	}
	var buff bytes.Buffer
	err = tmpl.Execute(&buff, errorData)
	if err != nil {
		http.Error(w, message, statusCode)
		return
	}
	w.Write(buff.Bytes())
}
