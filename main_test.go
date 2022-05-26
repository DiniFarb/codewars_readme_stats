package main

import (
	"andreasvogt/codewars_readme_stats/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
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
		t.Errorf("TestBaseRedirect() failed with error: %s", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 307 {
		t.Errorf("TestBaseRedirect() should end in 307 Not found instead got: %d", resp.Code)
	}
}
func Test404(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/f90123n", nil)
	if err != nil {
		t.Errorf("Test404 failed with error: %s", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 404 {
		t.Errorf("Test404 should end in 404 Not found instead got: %d", resp.Code)
	}
}
func TestWithoutUserParam(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/codewars", nil)
	if err != nil {
		t.Errorf("TestWithoutUserParam() failed with error: %s", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 500 {
		t.Errorf("TestWithoutUserParam() should end in 500 Internal Server Error instead got: %d", resp.Code)
	}
	want := `{"message":"Missing Query param =\u003e [user={yourname}]"}`
	got := resp.Body.String()
	if want != got {
		t.Errorf("TestWithoutUserParam() =  %v, want %v", got, want)
	}
}
func TestBasicUser(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/codewars?user=dinifarb", nil)
	if err != nil {
		t.Errorf("TestBasicUser() failed with error: %s", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	svgString := resp.Body.String()
	containsName := strings.Contains(svgString, "dinifarb's Codewars Stats")
	if !containsName {
		t.Errorf("TestBasicUser() does not contain correct username => dinifarb's Codewars Stats")
	}
}
