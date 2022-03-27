package codewars

import (
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

func Construct(settings url.Values, user Userdata) (template string, err error) {
	content, err := ioutil.ReadFile("./codewars/templates/codewarscard.svg")
	if err != nil {
		return
	}
	template = string(content)
	theme := "default"
	if settings.Get("theme") != "" {
		theme = settings.Get("theme")
	}

	if strings.Contains(user.Ranks.Overall.Name, "dan") {
		user.Ranks.Overall.Name = "dan"
	}
	if settings.Get("name") == "true" {
		template = strings.Replace(template, "{name}", user.Name, 1)
	} else {
		template = strings.Replace(template, "{name}", user.Username, 1)
	}
	if settings.Get("top_languages") == "true" {
		template = SetIcons(template, user.Ranks.Languages)
	}
	stroke := ""
	if settings.Get("stroke") != "" {
		stroke = "stroke: " + settings.Get("stroke")
	}
	template = strings.Replace(template, "{rankName}", user.Ranks.Overall.Name, 1)
	template = strings.Replace(template, "{clan}", user.Clan, 1)
	template = strings.Replace(template, "{leaderboardPosition}", strconv.Itoa(user.LeaderboardPosition), 1)
	template = strings.Replace(template, "{honor}", strconv.Itoa(user.Honor), 1)
	template = strings.Replace(template, "{score}", strconv.Itoa(user.Ranks.Overall.Score), 1)
	template = strings.Replace(template, "{rankColor}", LevelColors[user.Ranks.Overall.Name], -1)
	template = strings.Replace(template, "{totalCompleted}", strconv.Itoa(user.CodeChallenges.TotalCompleted), 1)
	template = strings.Replace(template, "{strokeColor}", stroke, 1)
	template = strings.Replace(template, "{cardColor}", Themes[theme].Card, -1)
	template = strings.Replace(template, "{headlineFontColor}", Themes[theme].Headline_font, -1)
	template = strings.Replace(template, "{bodyFontColor}", Themes[theme].Body_font, -1)
	template = strings.Replace(template, "{badgeColor}", Themes[theme].Rank_badge, -1)
	template = strings.Replace(template, "{iconColor}", Themes[theme].Icon, -1)
	return
}
