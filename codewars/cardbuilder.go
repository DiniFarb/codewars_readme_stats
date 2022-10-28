package codewars

import (
	"bytes"
	"fmt"
	templatePkg "html/template"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

type CardData struct {
	Theme       Theme
	ShowStroke  bool
	StrokeColor string
	LevelColor  string
}

func ConstructCard(settings url.Values, user User) (template string, err error) {
	content, err := ioutil.ReadFile("./codewars/templates/codewarscard.svg")
	if err != nil {
		return
	}
	template = string(content)

	if settings.Get("name") == "true" {
		template = strings.Replace(template, "{name}", user.Name, 1)
	} else {
		template = strings.Replace(template, "{name}", user.Username, 1)
	}
	if settings.Get("top_languages") == "true" {
		template = SetIcons(template, user.Ranks.Languages)
	}
	if settings.Get("hide_clan") != "true" {
		template = SetClan(template, user.Clan)
	}
	template = strings.Replace(template, "{rankName}", user.Ranks.Overall.Name, 1)
	template = strings.Replace(template, "{leaderboardPosition}", strconv.Itoa(user.LeaderboardPosition), 1)
	template = strings.Replace(template, "{honor}", strconv.Itoa(user.Honor), 1)
	template = strings.Replace(template, "{score}", strconv.Itoa(user.Ranks.Overall.Score), 1)
	template = strings.Replace(template, "{totalCompleted}", strconv.Itoa(user.CodeChallenges.TotalCompleted), 1)

	templ, err := templatePkg.New("svg").Parse(template)
	if err != nil {
		fmt.Printf("error creating template: %v\n", err)
		return
	}

	theme := Themes["default"]
	if settings.Get("theme") != "" {
		theme = Themes[settings.Get("theme")]
	}

	data := CardData{
		Theme:       theme,
		ShowStroke:  settings.Get("stroke") != "",
		StrokeColor: settings.Get("stroke"),
		LevelColor:  LevelColors[user.Ranks.Overall.Name],
	}

	var out = bytes.NewBuffer([]byte{})
	err = templ.ExecuteTemplate(out, "svg", &data)
	if err != nil {
		fmt.Printf("error ExecuteTemplate: %v\n", err)
		return
	}

	template = out.String()

	return
}
