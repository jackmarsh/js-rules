package main

import (
	"net/http"
)

func main() {	
	fs := http.FileServer(http.Dir("www"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "www/index.html")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
