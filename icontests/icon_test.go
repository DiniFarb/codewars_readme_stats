package icontests

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestLanguagesForIcons(t *testing.T) {
	resp, err := http.Get("https://www.codewars.com/kata/search/")
	if err != nil {
		t.Error("TestLanguagesForIcons() failed with error:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("TestLanguagesForIcons() failed with error:", err)
	}
	bodyString := string(body)
	fmt.Println(bodyString)
	s1 := strings.Split(bodyString, `<option value="my-languages">My Languages</option>`)[1]
	s2 := strings.Split(s1, `</select>`)[0]
	cw_languages := strings.Split(s2, `<option value="`)
	file, err := os.Open("../codewars/templates/icons/")
	if err != nil {
		t.Error("TestLanguagesForIcons() failed with error:", err)
	}
	defer file.Close()
	names, _ := file.Readdirnames(0)
	count := 0
	for i, l := range cw_languages {
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
