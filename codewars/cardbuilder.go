package codewars

import (
	"bytes"
	"fmt"
	"html/template"
	"net/url"
	"os"
	"strings"
)

type CardData struct {
	Theme        Theme
	User         User
	StrokeColor  string
	LevelColor   string
	ShowStroke   bool
	ShowTopLangs bool
	Nickname     bool
	HideClan     bool
}

func ConstructCard(settings url.Values, user User) (string, error) {
	content, err := os.ReadFile("./codewars/templates/codewarscard.svg")
	if err != nil {
		fmt.Printf("error loading template: %v\n", err)
		return "", err
	}
	templateString := string(content)
	data := CardData{
		LevelColor:   LevelColors[user.Ranks.Overall.Name],
		StrokeColor:  settings.Get("stroke"),
		ShowStroke:   settings.Get("stroke") != "",
		Nickname:     settings.Get("name") == "true",
		ShowTopLangs: settings.Get("top_languages") == "true",
		HideClan:     settings.Get("hide_clan") == "true",
		User:         user,
	}
	if data.ShowTopLangs {
		templateString = SetIcons(templateString, user.Ranks.Languages)
	}

	theme, exists := Themes[settings.Get("theme")]
	if exists {
		data.Theme = theme
		if strings.HasPrefix(settings.Get("theme"), "gradient") {
			data.Theme.HasGradient = true
			vals := strings.Split(data.Theme.Card, ",")
			if len(vals) != 8 {
				err := fmt.Errorf("gradient values must have 8 values")
				fmt.Printf("error parsing gradient values: %v\n", err)
				return "", err
			}
			for i, v := range vals {
				if v == "{LEVEL}" {
					vals[i] = LevelColors[user.Ranks.Overall.Name]
				}
			}
			g := GradientValues{
				StartColor:  vals[0],
				StopColor:   vals[1],
				X1:          vals[2],
				X2:          vals[4],
				Y2:          vals[5],
				OffsetStart: vals[6],
				OffsetStop:  vals[7],
			}
			data.Theme.GradientValues = g
		}
	} else {
		data.Theme = Themes["default"]
	}

	templ, err := template.New("svg").Parse(templateString)
	if err != nil {
		fmt.Printf("error creating template: %v\n", err)
		return "", err
	}

	var out = bytes.NewBuffer([]byte{})
	err = templ.ExecuteTemplate(out, "svg", &data)
	if err != nil {
		fmt.Printf("error ExecuteTemplate: %v\n", err)
		return "", err
	}
	templateString = out.String()
	return templateString, nil
}
