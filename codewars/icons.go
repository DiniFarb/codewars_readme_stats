package codewars

import (
	"sort"
	"strconv"
	"strings"
)

func GetTopLanguagesTemplate(icons string) (template string) {
	template = `
	<g transform="translate(150, 190)">
		<g class="stats" style="animation-delay: 1050ms">      
			<text class="stat bold"  y="12.5">Top trained languages</text>
			<text class="stat" x="170" y="12.5">
			</text>
			{icons}
		</g>
	</g>`
	return strings.Replace(template, "{icons}", icons, 1)
}

func GetIconTemplate(x string) (s string) {
	s = `
	<g transform="translate(${x},20)">
	  {svg}
	</g>`
	return strings.Replace(s, "{x}", x, 1)
}

func GetIconNotFoundSVG(iconName string) (s string) {
	s = `
	<svg viewBox="0 0 150 150" class="fail-icon-text">
		<text x="10" y="10" alignment-baseline="central" dominant-baseline="central" text-anchor="middle">
		{iconName}
		</text>
  	</svg>`
	return strings.Replace(s, "{iconName}", iconName, 1)
}

func SetIcons(template string, languages Languages) (t string) {
	scores := []int{}
	for _, v := range languages {
		scores = append(scores, v.Score)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	count := 0
	langs := []string{}
	for k, v := range languages {
		if len(langs) < 3 {
			if ArrContainsInt(v.Score, scores) {
				langs = append(langs, k)
			}
		}
		count++
	}
	i := 0
	x := -108
	ic := ""
	for k, _ := range languages {
		if i > 0 {
			x += 60
		}
		ic = GetIconTemplate(strconv.Itoa(x))
		icon := GetIcon(k)
		if icon == "" {
			ic = ic + strings.Replace(ic, "{svg}", GetIconNotFoundSVG(k), 1)
		} else {
			ic = ic + strings.Replace(ic, "{svg}", icon, 1)
			ic = ic + strings.Replace(ic, `viewBox="0 0 24 24"`, `viewBox="0 0 150 150" class="icons" fill="{iconColor}"`, 1)
		}
		i++
	}
	template = strings.Replace(template,
		`<svg width="500" height="195" viewBox="0 0 500 195"`,
		`<svg width="500" height="280" viewBox="0 0 500 280"`, 1)
	return strings.Replace(template, "{icons}", GetTopLanguagesTemplate(ic), 1)
}

func GetIcon(language string) (icon string) {
	return ""
}

func ArrContainsInt(si int, arr []int) (r bool) {
	r = false
	for _, v := range arr {
		if v == si {
			r = true
			return
		}
	}
	return
}
