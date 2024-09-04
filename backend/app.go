package main

import (
	"log"
	"masisserson/mtg-lib/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/api/library", handlers.LibViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
