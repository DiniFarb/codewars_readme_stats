package test

import (
	"andreasvogt/codewars_readme_stats/routes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://github.com/andreasvogt89/codewars_readme_stats")
	})
	r.GET("/codewars", routes.GET_CodewarsCard)
	return r
}

func TestBaseRedirect(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, 307, resp.Code)
}
func Test404(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/f90123n", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, 404, resp.Code)
}

func TestWithoutUserParam(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/codewars", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Body.String(), `{"message":"Missing Query param =\u003e [user={yourname}]"}`)
	assert.Equal(t, resp.Code, 500)
}

func TestUnknownUser(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/codewars?user=xxxxxxxxxxxxx", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 500)
}
