package codewars

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Userdata struct {
	Username            string         `json:"username"`
	Name                string         `json:"name"`
	Honor               int            `json:"honor"`
	Clan                string         `json:"clan"`
	LeaderboardPosition int            `json:"leaderboardPosition"`
	Ranks               Ranks          `json:"ranks"`
	CodeChallenges      CodeChallenges `json:"codeChallenges"`
}
type Ranks struct {
	Overall   Overall   `json:"overall"`
	Languages Languages `json:"languages"`
}
type Overall struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type Languages map[string]Language

type Language struct {
	Score int `json:"score"`
}

type CodeChallenges struct {
	TotalCompleted int `json:"totalCompleted"`
}

var LevelColors = map[string]string{
	"1 kyu:": "#866CC7",
	"2 kyu":  "#866CC7",
	"3 kyu":  "#3C7EBB",
	"4 kyu":  "#3C7EBB",
	"5 kyu":  "#ECB613",
	"6 kyu":  "#ECB613",
	"7 kyu":  "#E6E6E6",
	"8 kyu":  "#E6E6E6",
	"dan":    "#999999",
}

func GetUserData(username string) (userdata Userdata, err error) {
	resp, err := http.Get("https://www.codewars.com/api/v1/users/" + username)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(body), &userdata)
	if err != nil {
		return
	}
	return
}

func Construct(params url.Values, user Userdata) (template string, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	content, err := ioutil.ReadFile(dir + "/codewars/templates/codewarscard.svg")
	if err != nil {
		return
	}
	template = string(content)
	theme := "default"
	if params.Get("theme") != "" {
		theme = params.Get("theme")
	}

	if strings.Contains(user.Ranks.Overall.Name, "dan") {
		user.Ranks.Overall.Name = "dan"
	}
	if params.Get("name") == "true" {
		template = strings.Replace(template, "{name}", user.Name, 1)
	} else {
		template = strings.Replace(template, "{name}", user.Username, 1)
	}
	if params.Get("top_languages") == "true" {
		template = SetIcons(template, user.Ranks.Languages)
	}
	template = strings.Replace(template, "{rankName}", user.Ranks.Overall.Name, 1)
	template = strings.Replace(template, "{clan}", user.Clan, 1)
	template = strings.Replace(template, "{leaderboardPosition}", strconv.Itoa(user.LeaderboardPosition), 1)
	template = strings.Replace(template, "{honor}", strconv.Itoa(user.Honor), 1)
	template = strings.Replace(template, "{score}", strconv.Itoa(user.Ranks.Overall.Score), 1)
	template = strings.Replace(template, "{rankColor}", LevelColors[user.Ranks.Overall.Name], -1)
	template = strings.Replace(template, "{totalCompleted}", strconv.Itoa(user.CodeChallenges.TotalCompleted), 1)
	template = strings.Replace(template, "{strokeColor}", params.Get("stroke"), 1)
	template = strings.Replace(template, "{cardColor}", Themes[theme].Card, -1)
	template = strings.Replace(template, "{headlineFontColor}", Themes[theme].Headline_font, -1)
	template = strings.Replace(template, "{bodyFontColor}", Themes[theme].Body_font, -1)
	template = strings.Replace(template, "{badgeColor}", Themes[theme].Rank_badge, -1)
	return
}
