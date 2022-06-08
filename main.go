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
		c.Redirect(http.StatusTemporaryRedirect, "https://github.com/dinifarb/codewars_readme_stats")
	})
	r.GET("/codewars", routes.GET_CodewarsCard)
	r.GET("/health", routes.Health)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("Start service on port::: ", port)
	r.Run(":" + port)
}
