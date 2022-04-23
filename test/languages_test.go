package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestLanguagesForIcons(t *testing.T) {
	resp, err := http.Get("https://www.codewars.com/kata/search/")
	if err != nil {
		log.Println("Could not get data from codewars: ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Could parse data from codewars: ", err)
	}
	bodyString := string(body)
	log.Println("Skip Language check because could parse data from codewars: ", err)
	s1 := strings.Split(bodyString, `<option value="my-languages">My Languages</option>`)[1]
	s2 := strings.Split(s1, `</select>`)[0]
	cw_languages := strings.Split(s2, `<option value="`)
	file, err := os.Open("../codewars/templates/icons/")
	if err != nil {
		log.Println("Cloud not open icons dir: ", err)
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
				log.Println("No icon for: ", lang)
				count++
			}
		}
	}
	assert.Equal(t, count, 0)
}
