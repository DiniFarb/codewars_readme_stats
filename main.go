package main

import (
	"dinifarb/codewars_readme_stats/routes"
	"dinifarb/codewars_readme_stats/stats"
	"log"
	"net/http"
	"os"
)

func main() {
	stats.Init()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	mu := http.NewServeMux()
	mu.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/dinifarb/codewars_readme_stats", http.StatusPermanentRedirect)
	})
	mu.HandleFunc("/codewars", routes.GetCodewarsCard)
	mu.HandleFunc("/health", routes.Health)
	mu.HandleFunc("/stats", routes.GetStats)
	log.Println("Start service on port::: ", port)
	log.Fatal(http.ListenAndServe(":"+port, mu))
}
