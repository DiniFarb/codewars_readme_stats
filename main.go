package main

import (
	"andreasvogt/codewars_readme_stats/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://github.com/andreasvogt89/codewars_api")
	})
	r.GET("/codewars", routes.GET_CodewarsCard)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("Start service on port::: ", port)
	r.Run(":" + port)
}
