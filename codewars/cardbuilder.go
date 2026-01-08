package codewars

import (
	"fmt"
	"net/url"
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
	Svg               *svg.SVG
	Theme             Theme
	User              User
	StrokeColor       string
	LevelColor        string
	ShowStroke        bool
	ShowTopLangs      bool
	Nickname          bool
	HideClan          bool
	HideTitle         bool
	HasGradient       bool
	AnimationDisabeld bool
	contentPointer_y  int
	levelPointer_y    int
	cardHeight        int
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
		Svg:               &s,
		Theme:             theme,
		User:              *user,
		LevelColor:        LevelColors[user.Ranks.Overall.Name],
		StrokeColor:       settings.Get("stroke"),
		ShowStroke:        settings.Get("stroke") != "",
		Nickname:          settings.Get("name") == "true",
		ShowTopLangs:      settings.Get("top_languages") == "true",
		HideClan:          settings.Get("hide_clan") == "true",
		HideTitle:         settings.Get("hide_title") == "true",
		HasGradient:       strings.HasPrefix(settings.Get("theme"), "gradient"),
		AnimationDisabeld: settings.Get("animation") == "false",
		contentPointer_y:  25,
		levelPointer_y:    0,
		cardHeight:        155,
	}
	card.CreateSvg()
	if card.HasGradient {
		card.SetGradient()
	}
	if !card.HideTitle {
		card.SetTitle()
	}
	card.SetStatsTexts()
	card.SetLevel()
	if card.ShowTopLangs {
		card.SetIcons()
	}
	s.End()
	return svgWriter.content, nil
}

func (c *CardData) CreateSvg() {
	if !c.HideTitle {
		c.cardHeight += 15
	}
	if !c.HideClan {
		c.cardHeight += 25
	}
	if c.ShowTopLangs {
		c.cardHeight += 85
	}
	box := fmt.Sprintf(`viewBox="0 0 500 %d"`, c.cardHeight)
	c.Svg.Start(500, c.cardHeight, box)
	attr := []string{`rx="4.5"`}
	if c.ShowStroke {
		attr = append(attr, fmt.Sprintf(`stroke="%s"`, c.StrokeColor))
	}
	if c.HasGradient {
		attr = append(attr, `fill="url(#grad)"`)
	} else {
		attr = append(attr, fmt.Sprintf(`fill="%s"`, c.Theme.CardColor))
	}
	c.Svg.Rect(0, 0, 500, c.cardHeight, strings.Join(attr, " "))
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
	c.Svg.Text(10, c.contentPointer_y, name, attr...)
	c.contentPointer_y += 25
	c.AddAnimation("title", "0.3")
}

func (c *CardData) SetStatsTexts() {
	stats := []string{
		fmt.Sprintf("Position:-%d", c.User.LeaderboardPosition),
		fmt.Sprintf("Honor:-%d", c.User.Honor),
		fmt.Sprintf("Score:-%d", c.User.Ranks.Overall.Score),
		fmt.Sprintf("Solved Katas:-%d", c.User.CodeChallenges.TotalCompleted),
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
	delay := 0.2
	c.Svg.Group(attr...)
	start_y := c.contentPointer_y
	for i, stat := range stats {
		key := strings.Split(stat, "-")[0]
		value := strings.Split(stat, "-")[1]
		idKey := fmt.Sprintf("k-%d", i)
		delayString := fmt.Sprintf(`%.1f`, delay)
		c.Svg.Text(15, c.contentPointer_y, key, `opacity="0"`, fmt.Sprintf(`id="%s"`, idKey))
		c.AddAnimation(idKey, delayString)
		idValue := fmt.Sprintf("v-%d", i)
		c.Svg.Text(150, c.contentPointer_y, value, `opacity="0"`, fmt.Sprintf(`id="%s"`, idValue))
		c.AddAnimation(idValue, delayString)
		c.contentPointer_y += 25
		delay += 0.2
	}
	halfstats := ((c.contentPointer_y - 25) - start_y)
	c.levelPointer_y = start_y + halfstats/2
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
	c.Svg.Polygon([]int{340, 355, 435, 450, 435, 355}, []int{c.levelPointer_y, c.levelPointer_y - 27, c.levelPointer_y - 27, c.levelPointer_y, c.levelPointer_y + 27, c.levelPointer_y + 27}, polyAttr...)
	c.AddAnimation("level", "1")
	textAttr := []string{
		fmt.Sprintf(`fill="%s"`, c.LevelColor),
		`font-weight="600"`,
		`font-size="30px"`,
		fmt.Sprintf(`font-family="%s"`, c.Theme.Font),
		`opacity="0"`,
		`id="level-text"`,
	}
	c.Svg.Text(358, c.levelPointer_y+10, c.User.Ranks.Overall.Name, textAttr...)
	c.AddAnimation("level-text", "1")
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
	c.contentPointer_y += 30
	c.Svg.Text(199, c.contentPointer_y, "Top Languages", textAttr...)
	c.contentPointer_y += 50
	c.AddAnimation("top-languages", "0.8")
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
		x = 220
	case 2:
		x = 190
	default:
		x = 160
	}
	for i, k := range keys {
		if i > 0 {
			x += 60
		}
		if i > 2 {
			continue
		}
		attr := []string{
			fmt.Sprintf(`fill="%s"`, c.Theme.IconColor),
			fmt.Sprintf(`id="icon-%d"`, i),
			`opacity="0"`,
			fmt.Sprintf(`transform="translate(%d, %d)"`, x, c.contentPointer_y),
		}
		c.Svg.Group(attr...)
		if icon, ok := Icons[k]; ok {
			c.Svg.Path(icon, `transform="scale(1,-1) scale(0.05)"`)
		} else {
			c.Svg.Text(-10, -10, k[0:2])
		}
		c.AddAnimation("icon-"+strconv.Itoa(i), "1.2")
		c.Svg.Gend()
		i++
	}
	c.Svg.Gend()
}

func (c *CardData) AddAnimation(id string, begin string) {
	if c.AnimationDisabeld {
		c.Svg.Animate("#"+id, "opacity", 1, 1, 1, 1, `fill="freeze"`)
	} else {
		c.Svg.Animate("#"+id, "opacity", 0, 1, 1, 1, fmt.Sprintf(`begin="%s"`, begin), `fill="freeze"`)
	}
}
