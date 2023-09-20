package routes

import (
	"dinifarb/codewars_readme_stats/codewars"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetCodewarsCard(c *gin.Context) {
	username := c.Request.URL.Query().Get("user")
	if username == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Missing Query param => [user={yourname}]"})
		return
	}
	var user codewars.User
	err := user.GetUserData(username)
	if err != nil {
		log.Println("Get Userdata failed with: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not get Userdata from codewars"})
		return
	}
	card, err := codewars.ConstructCard(c.Request.URL.Query(), &user)
	if err != nil {
		log.Println("Cunstruct codewars card failed with: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error while constructing codewars card"})
		return
	}

	cache := c.Request.URL.Query().Get("cache_control")
	if cache == "" {
		c.Writer.Header().Set("Cache-Control", "public, max-age=no-cache")
	} else {
		c.Writer.Header().Set("Cache-Control", "public, max-age="+cache)
	}

	c.Writer.Header().Set("Content-Type", "image/svg+xml")
	c.String(http.StatusOK, card)
}

func Health(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "image/svg+xml")
	c.Writer.Header().Set("Cache-Control", "public, max-age=no-cache")
	content, err := os.ReadFile("./codewars/templates/health/on.svg")
	if err != nil {
		c.AbortWithError(400, err)
	}
	c.String(http.StatusOK, string(content))
}
