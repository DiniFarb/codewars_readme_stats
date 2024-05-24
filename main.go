package main

import (
	"dinifarb/codewars_readme_stats/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/dinifarb/codewars_readme_stats", http.StatusPermanentRedirect)
	})
	http.HandleFunc("/codewars", routes.GetCodewarsCard)
	http.HandleFunc("/health", routes.Health)
	log.Println("Start service on port::: ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
