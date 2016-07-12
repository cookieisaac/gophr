package main

import (
	"net/http"
	"log"
)

func main() {
		mux := http.NewServeMux()
		
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			RenderTemplate(w, r, "index/home", nil)
		})
		
		mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
		
		log.Fatal(http.ListenAndServe(":3000", mux))
}