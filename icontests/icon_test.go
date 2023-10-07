package icontests

import (
	"dinifarb/codewars_readme_stats/codewars"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type CwLangs struct {
	Data []CwLang `json:"data"`
}

type CwLang struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func TestLanguagesForIcons(t *testing.T) {
	resp, err := http.Get("https://www.codewars.com/api/v1/languages/")
	if err != nil {
		t.Error("TestLanguagesForIcons() failed with error:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("TestLanguagesForIcons() failed with error:", err)
	}
	var langs CwLangs
	err = json.Unmarshal(body, &langs)
	if err != nil {
		t.Error("TestLanguagesForIcons() failed with error:", err)
	}
	for _, l := range langs.Data {
		if _, ok := codewars.Icons[l.Id]; !ok {
			t.Error("No icon for: ", l)
		}
	}
}
