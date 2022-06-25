package codewars

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func GetTopLanguagesTemplate(icons string) (template string) {
	template = `
	<g transform="translate(150, 190)">
		<g class="stats" style="animation-delay: 1050ms">      
			<text class="stat bold"  y="12.5">Top trained languages</text>
			<text class="stat" x="170" y="12.5"></text>
			{icons}
		</g>
	</g>`
	return strings.Replace(template, "{icons}", icons, 1)
}

func GetIconTemplate(x string) (s string) {
	s = `
	<g transform="translate({x},20)">
	  {svg}
	</g>`
	return strings.Replace(s, "{x}", x, 1)
}

func GetIconNotFoundSVG(iconName string) (s string) {
	s = `
	<svg viewBox="0 0 150 150" class="fail-icon-text" fill="{iconColor}">
		<text x="10" y="10" alignment-baseline="central" dominant-baseline="central" text-anchor="middle">
		{iconName}
		</text>
  	</svg>`
	return strings.Replace(s, "{iconName}", iconName, 1)
}

func SetIcons(template string, languages Languages) (t string) {
	keys := make([]string, 0, len(languages))
	for key := range languages {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return languages[keys[i]].Score > languages[keys[j]].Score
	})
	var x int
	switch len(keys) {
	case 1:
		x = -48
	case 2:
		x = -78
	default:
		x = -108
	}
	ic := ""
	for i, k := range keys {
		if i > 0 {
			x += 60
		}
		if i > 2 {
			continue
		}
		icon, err := GetIcon(k)
		if err != nil {
			log.Printf("Could not get icon svg for: %v => %v", k, err)
			ic = ic + strings.Replace(GetIconTemplate(strconv.Itoa(x)), "{svg}", GetIconNotFoundSVG(k), 1)
		} else {
			iconstring := strings.Replace(GetIconTemplate(strconv.Itoa(x)), "{svg}", icon, 1)
			ic = ic + strings.Replace(iconstring, `viewBox="0 0 24 24"`, `viewBox="0 0 150 150" class="icons" fill="{iconColor}"`, 1)
		}
		i++
	}
	template = strings.Replace(template,
		`<svg width="500" height="195" viewBox="0 0 500 195"`,
		`<svg width="500" height="280" viewBox="0 0 500 280"`, 1)
	return strings.Replace(template, "{icons}", GetTopLanguagesTemplate(ic), 1)
}

func GetIcon(language string) (icon string, err error) {
	content, err := ioutil.ReadFile("./codewars/templates/icons/" + language + ".svg")
	if err != nil {
		return
	}
	icon = string(content)
	return
}
