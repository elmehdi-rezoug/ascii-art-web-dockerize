package main

import (
	"ascii-art-web/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)
	http.HandleFunc("/style.css", handlers.CssHandler)

	port := 8080
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("failed to start server")
	}
}
