package icontests

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestLanguagesForIcons(t *testing.T) {
	resp, err := http.Get("https://docs.codewars.com/languages/")
	if err != nil {
		t.Error("TestLanguagesForIcons() failed with error:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("TestLanguagesForIcons() failed with error:", err)
	}
	bodyString := string(body)
	pattern := `href="/languages/([^"]+)"`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(bodyString, -1)
	languages := make([]string, 0)
	fmt.Println(matches)
	for _, match := range matches {
		if len(match) >= 2 {
			languageName := match[1]
			fmt.Println(languageName)
			languages = append(languages, languageName)
		}
	}
	file, err := os.Open("../codewars/icons/")
	if err != nil {
		t.Error("TestLanguagesForIcons() failed with error:", err)
	}
	defer file.Close()
	names, _ := file.Readdirnames(0)
	count := 0
	for i, l := range languages {
		if i != 0 {
			lang := strings.Split(l, `">`)[0]
			contains := false
			for _, n := range names {
				name := strings.Replace(n, ".svg", "", 1)
				if name == lang {
					contains = true
				}
			}
			if !contains {
				t.Error("No icon for: ", lang)
				count++
			}
		}
	}
}
