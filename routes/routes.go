package routes

import (
	"andreasvogt/codewars_readme_stats/codewars"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET_CodewarsCard(c *gin.Context) {
	username := c.Request.URL.Query().Get("user")
	if username == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Missing Query param => [user={yourname}]"})
		return
	}
	user, err := codewars.GetUserData(username)
	if err != nil {
		log.Println("Get Userdata failed with: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not get Userdata from codewars"})
		return
	}
	c.Writer.Header().Set("Content-Type", "image/svg+xml")
	c.Writer.Header().Set("Cache-Control", "public, max-age=no-cache")
	data, err := codewars.Construct(c.Request.URL.Query(), user)
	if err != nil {
		log.Println("Cunstruct codewars card failed with: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error while constructing codewars card"})
		return
	}
	c.String(http.StatusOK, data)
}
