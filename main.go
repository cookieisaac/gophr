package main

import (
	"net/http"
	"log"
)

func main() {
		mux := http.NewServeMux()
		mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
		
		log.Fatal(http.ListenAndServe(":3000", mux))
}