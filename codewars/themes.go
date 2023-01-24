package codewars

type Theme struct {
	Card          string
	Headline_font string
	Body_font     string
	Icon          string
	Rank_badge    string
}

var Themes = map[string]Theme{
	"default": {
		Card:          "#262729",
		Headline_font: "#F1F5F3",
		Body_font:     "#BB432C",
		Icon:          "#6795DE",
		Rank_badge:    "#181919",
	},
	"light": {
		Card:          "#fffefe",
		Headline_font: "#181919",
		Body_font:     "#BB432C",
		Icon:          "#6795DE",
		Rank_badge:    "#181919",
	},
	"dark": {
		Card:          "#000000",
		Headline_font: "#F1F5F3",
		Body_font:     "#BB432C",
		Icon:          "#6795DE",
		Rank_badge:    "#181919",
	},
	"highcontrast": {
		Card:          "#000",
		Headline_font: "#e7f216",
		Body_font:     "#fff",
		Icon:          "#00ffff",
		Rank_badge:    "#181919",
	},
	"midnight_purple": {
		Card:          "#000",
		Headline_font: "#9745f5",
		Body_font:     "#fff",
		Icon:          "#9f4bff",
		Rank_badge:    "#181919",
	},
	"midnight_blue": {
		Card:          "#000",
		Headline_font: "#0079fa",
		Body_font:     "#0079fa",
		Icon:          "#0079fa",
		Rank_badge:    "#2100fa",
	},
	"purple_dark": {
		Card:          "#2d2b55",
		Headline_font: "#c792ea",
		Body_font:     "#a599e9",
		Icon:          "#b362ff",
		Rank_badge:    "#181919",
	},
	"purple_light": {
		Card:          "#ffffff",
		Headline_font: "#a84ee5",
		Body_font:     "#8c64e0",
		Icon:          "#b362ff",
		Rank_badge:    "#181919",
	},
	"nightowl": {
		Card:          "#011627",
		Headline_font: "#c792ea",
		Body_font:     "#7fdbca",
		Icon:          "#ffeb95",
		Rank_badge:    "#181919",
	},
	"solarized_dark": {
		Card:          "#002b36",
		Headline_font: "#268bd2",
		Body_font:     "#859900",
		Icon:          "#b58900",
		Rank_badge:    "#181919",
	},
	"solarized_light": {
		Card:          "#fdf6e3",
		Headline_font: "#268bd2",
		Body_font:     "#859900",
		Icon:          "#b58900",
		Rank_badge:    "#181919",
	},
}
