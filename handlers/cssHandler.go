package handlers

import "net/http"

func CssHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "assets/style.css")
}