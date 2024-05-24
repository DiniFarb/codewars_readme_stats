package routes

import (
	"dinifarb/codewars_readme_stats/codewars"
	"log"
	"net/http"
	"os"
)

func GetCodewarsCard(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing Query param => [user={yourname}]"))
		return
	}

	var user codewars.User
	err := user.GetUserData(username)
	if err != nil {
		log.Println("Get Userdata failed with: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not get Userdata from codewars"))
		return
	}

	card, err := codewars.CreateSvg(r.URL.Query(), &user)
	if err != nil {
		log.Println("Cunstruct codewars card failed with: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error while constructing codewars card"))
		return
	}

	cache := r.URL.Query().Get("cache_control")
	if cache == "" {
		w.Header().Set("Cache-Control", "public, max-age=no-cache")
	} else {
		w.Header().Set("Cache-Control", "public, max-age="+cache)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(card))
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "public, max-age=no-cache")
	content, err := os.ReadFile("./routes/assets/on.svg")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not read health check file"))
		return
	}
	w.Write([]byte(content))
}
