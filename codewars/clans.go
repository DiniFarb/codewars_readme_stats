package codewars

import (
	"fmt"
	"strings"
)

func SetClan(template, clan string) (t string) {

	body := `<g transform="translate(0, 0)">
				<g class="stats" style="animation-delay: 450ms" transform="translate(25, 0)">
				<text class="stat bold" y="12.5">Clan:</text>
				<text class="stat" x="170" y="12.5">%s
				</text>
				</g>
			</g>`

	body = fmt.Sprintf(body, clan)

	return strings.Replace(template, "{clan}", body, 1)
}
