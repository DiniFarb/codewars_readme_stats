package main

import (
	"dinifarb/codewars_readme_stats/routes"
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
	r.GET("/codewars", routes.GetCodewarsCard)
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
	containsName := strings.Contains(svgString, "dinifarb")
	if !containsName {
		t.Errorf("TestBasicUser() does not contain correct username => dinifarb's Codewars Stats")
	}
}

func TestUserWithoutClan(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/codewars?user=dinifarb&hinde_clan=true", nil)
	if err != nil {
		t.Errorf("TestUserWithoutClan() failed with error: %s", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	svgString := resp.Body.String()
	containsClan := !strings.Contains(svgString, "Clan")
	if containsClan {
		t.Errorf("TestUserWithoutClan() contains clan")
	}
}

func TestUserWithTopLanguages(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/codewars?user=dinifarb&top_languages=true", nil)
	if err != nil {
		t.Errorf("TestUserWithTopLanguages() failed with error: %s", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	svgString := resp.Body.String()
	containsTopLanguages := strings.Contains(svgString, "Top Languages")
	if !containsTopLanguages {
		t.Errorf("TestUserWithTopLanguages() does not contain top languages")
	}
}

func TestUserWithGradient(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/codewars?user=dinifarb&theme=gradient", nil)
	if err != nil {
		t.Errorf("TestUserWithGradient() failed with error: %s", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	svgString := resp.Body.String()
	containsGradient := strings.Contains(svgString, "linearGradient")
	if !containsGradient {
		t.Errorf("TestUserWithGradient() does not contain gradient")
	}
}

func TestUserWithStroke(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/codewars?user=dinifarb&stroke=red", nil)
	if err != nil {
		t.Errorf("TestUserWithStroke() failed with error: %s", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	svgString := resp.Body.String()
	containsStroke := strings.Contains(svgString, "stroke=\"red\"")
	if !containsStroke {
		t.Errorf("TestUserWithStroke() does not contain stroke")
	}
}

func TestUserWithNickname(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/codewars?user=dinifarb&name=true", nil)
	if err != nil {
		t.Errorf("TestUserWithNickname() failed with error: %s", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	svgString := resp.Body.String()
	containsNickname := strings.Contains(svgString, "DiniFarb")
	if !containsNickname {
		t.Errorf("TestUserWithNickname() does not contain nickname")
	}
}

func TestUserWithTheme(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/codewars?user=dinifarb&theme=dark", nil)
	if err != nil {
		t.Errorf("TestUserWithTheme() failed with error: %s", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	svgString := resp.Body.String()
	containsTheme := strings.Contains(svgString, "fill=\"#000000\"")
	if !containsTheme {
		t.Errorf("TestUserWithTheme() does not contain theme")
	}
}
