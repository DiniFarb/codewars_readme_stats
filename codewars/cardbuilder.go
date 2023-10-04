package codewars

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"

	svg "github.com/ajstarks/svgo"
)

type svgWriter struct {
	content string
}

func (c *svgWriter) Write(data []byte) (n int, err error) {
	c.content += string(data)
	return len(data), nil
}

type CardData struct {
	Svg          *svg.SVG
	Theme        Theme
	User         User
	StrokeColor  string
	LevelColor   string
	ShowStroke   bool
	ShowTopLangs bool
	Nickname     bool
	HideClan     bool
	HasGradient  bool
}

func CreateSvg(settings url.Values, user *User) (string, error) {
	svgWriter := svgWriter{
		content: "",
	}
	s := *svg.New(&svgWriter)
	theme, exists := Themes[settings.Get("theme")]
	if !exists {
		theme = Themes["default"]
	}
	card := CardData{
		Svg:          &s,
		Theme:        theme,
		User:         *user,
		LevelColor:   LevelColors[user.Ranks.Overall.Name],
		StrokeColor:  settings.Get("stroke"),
		ShowStroke:   settings.Get("stroke") != "",
		Nickname:     settings.Get("name") == "true",
		ShowTopLangs: settings.Get("top_languages") == "true",
		HideClan:     settings.Get("hide_clan") == "true",
		HasGradient:  strings.HasPrefix(settings.Get("theme"), "gradient"),
	}
	card.CreateSvg()
	if card.HasGradient {
		card.SetGradient()
	}
	card.SetTitle()
	card.SetStatsTexts()
	card.SetLevel()
	if card.ShowTopLangs {
		card.SetIcons()
	}
	s.End()
	return svgWriter.content, nil
}

func (c *CardData) CreateSvg() {
	height := 195
	switch {
	case c.ShowTopLangs && c.HideClan:
		height = 280
	case c.ShowTopLangs:
		height = 255
	case c.HideClan:
		height = 170
	}
	box := fmt.Sprintf(`viewBox="0 0 500 %d"`, height)
	c.Svg.Start(500, height, box)
	attr := []string{`rx="4.5"`}
	if c.ShowStroke {
		attr = append(attr, fmt.Sprintf(`stroke="%s"`, c.StrokeColor))
	}
	if c.HasGradient {
		attr = append(attr, `fill="url(#grad)"`)
	} else {
		attr = append(attr, fmt.Sprintf(`fill="%s"`, c.Theme.CardColor))
	}
	c.Svg.Rect(0, 0, 500, height, strings.Join(attr, " "))
}

func (c *CardData) SetTitle() {
	attr := []string{
		fmt.Sprintf(`fill="%s"`, c.Theme.HeadlineFontColor),
		`font-weight="600"`,
		`font-size="20px"`,
		fmt.Sprintf(`font-family="%s"`, c.Theme.Font),
		`opacity="0"`,
		`id="title"`,
	}
	name := "%s's Codewars Stats"
	if c.Nickname {
		name = fmt.Sprintf(name, c.User.Name)
	} else {
		name = fmt.Sprintf(name, c.User.Username)
	}
	c.Svg.Text(10, 25, name, attr...)
	c.Svg.Animate("#title", "opacity", 0, 1, 2, 1, `begin="0.3"`, `fill="freeze"`)
}

func (c *CardData) SetStatsTexts() {
	stats := []string{
		fmt.Sprintf("Position:-%d", c.User.LeaderboardPosition),
		fmt.Sprintf("Honor:-%d", c.User.Honor),
		fmt.Sprintf("Score:-%d", c.User.Ranks.Overall.Score),
		fmt.Sprintf("SolvedKatas:-%d", c.User.CodeChallenges.TotalCompleted),
	}
	if !c.HideClan {
		clan := []string{fmt.Sprintf("Clan:-%s", c.User.Clan)}
		stats = append(clan, stats...)
	}
	attr := []string{
		fmt.Sprintf(`fill="%s"`, c.Theme.BodyFontColor),
		`font-weight="500"`,
		`font-size="15px"`,
		fmt.Sprintf(`font-family="%s"`, c.Theme.Font),
		`id="title"`,
	}
	height := 60
	delay := 0.2
	c.Svg.Group(attr...)
	for i, stat := range stats {
		key := strings.Split(stat, "-")[0]
		value := strings.Split(stat, "-")[1]
		idKey := fmt.Sprintf("k-%d", i)
		delayString := fmt.Sprintf(`begin="%.1f"`, delay)
		c.Svg.Text(15, height, key, `opacity="0"`, fmt.Sprintf(`id="%s"`, idKey))
		c.Svg.Animate("#"+idKey, "opacity", 0, 1, 2, 1, delayString, `fill="freeze"`)
		idValue := fmt.Sprintf("v-%d", i)
		c.Svg.Text(150, height, value, `opacity="0"`, fmt.Sprintf(`id="%s"`, idValue))
		c.Svg.Animate("#"+idValue, "opacity", 0, 1, 2, 1, delayString, `fill="freeze"`)
		height += 25
		delay += 0.2
	}
	c.Svg.Gend()
}

func (c *CardData) SetLevel() {
	c.Svg.Group()
	polyAttr := []string{
		fmt.Sprintf(`fill="%s"`, c.Theme.RankBadgeColor),
		fmt.Sprintf(`stroke="%s"`, c.LevelColor),
		`stroke-width="3"`,
		`opacity="0"`,
		`id="level"`,
	}
	c.Svg.Polygon([]int{340, 355, 435, 450, 435, 355}, []int{107, 80, 80, 107, 135, 135}, polyAttr...)
	c.Svg.Animate("#level", "opacity", 0, 1, 2, 1, `begin="1"`, `fill="freeze"`)
	textAttr := []string{
		fmt.Sprintf(`fill="%s"`, c.LevelColor),
		`font-weight="600"`,
		`font-size="30px"`,
		fmt.Sprintf(`font-family="%s"`, c.Theme.Font),
		`opacity="0"`,
		`id="level-text"`,
	}
	c.Svg.Text(360, 118, c.User.Ranks.Overall.Name, textAttr...)
	c.Svg.Animate("#level-text", "opacity", 0, 1, 2, 1, `begin="1"`, `fill="freeze"`)
	c.Svg.Gend()
}

func (c *CardData) SetGradient() error {
	vals := strings.Split(c.Theme.CardColor, ",")
	if len(vals) != 8 {
		err := fmt.Errorf("gradient values must have 8 values")
		fmt.Printf("error parsing gradient values: %v\n", err)
		return err
	}
	for i, v := range vals[0:2] {
		if v == "{LEVEL}" {
			vals[i] = c.LevelColor
		}
	}
	var parsedVals []uint8
	for _, v := range vals[2:8] {
		x, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("error parsing gradient values: %v\n", err)
			return err
		}
		parsedVals = append(parsedVals, uint8(x))
	}
	c.Svg.Group()
	c.Svg.LinearGradient("grad",
		parsedVals[0],
		parsedVals[1],
		parsedVals[2],
		parsedVals[3],
		[]svg.Offcolor{
			{Color: vals[0], Offset: parsedVals[4], Opacity: 1},
			{Color: vals[1], Offset: parsedVals[5], Opacity: 1},
		},
	)
	c.Svg.Gend()
	return nil
}

const noIcon = `
<svg icon="this_is_a_hook">
	<text x="10" y="10" alignment-baseline="central" dominant-baseline="central" text-anchor="middle">
	%s
	</text>
</svg>`

func (c *CardData) SetIcons() {
	c.Svg.Group()
	textAttr := []string{
		fmt.Sprintf(`fill="%s"`, c.Theme.BodyFontColor),
		`font-weight="500"`,
		`font-size="15px"`,
		fmt.Sprintf(`font-family="%s"`, c.Theme.Font),
		`opacity="0"`,
		`id="top-languages"`,
	}
	c.Svg.Text(199, 190, "Top Languages", textAttr...)
	c.Svg.Animate("#top-languages", "opacity", 0, 1, 2, 1, `begin="0.8"`, `fill="freeze"`)
	keys := make([]string, 0, len(c.User.Ranks.Languages))
	for key := range c.User.Ranks.Languages {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return c.User.Ranks.Languages[keys[i]].Score > c.User.Ranks.Languages[keys[j]].Score
	})
	var x int
	switch len(keys) {
	case 1:
		x = 110
	case 2:
		x = 80
	default:
		x = 50
	}
	for i, k := range keys {
		if i > 0 {
			x += 60
		}
		if i > 2 {
			continue
		}
		icon, err := os.ReadFile("./codewars/icons/" + k + ".svg")
		if err != nil {
			log.Printf("Could not get icon svg for: %v => %v", k, err)
			icon = []byte(fmt.Sprintf(noIcon, k[0:2]))
		}
		attr := []string{
			fmt.Sprintf(`fill="%s"`, c.Theme.IconColor),
			fmt.Sprintf(`id="icon-%d"`, i),
			`opacity="0"`,
			fmt.Sprintf(`x="%d"`, x),
			`y="200"`,
			`viewBox="0 0 150 150"`,
		}
		ic := strings.Replace(string(icon), `icon="this_is_a_hook"`, strings.Join(attr, " "), 1)
		c.Svg.Writer.Write([]byte(ic))
		c.Svg.Animate("#icon-"+strconv.Itoa(i), "opacity", 0, 1, 2, 1, `begin="1.2s"`, `fill="freeze"`)
		i++
	}
	c.Svg.Gend()
}
