package codewars

type Theme struct {
	CardColor         string
	HeadlineFontColor string
	BodyFontColor     string
	IconColor         string
	RankBadgeColor    string
	Font              string
	GradientValues    GradientValues
}

type GradientValues struct {
	StartColor  string
	StopColor   string
	X1          string
	Y1          string
	X2          string
	Y2          string
	OffsetStart string
	OffsetStop  string
}

// If you want to add a new gradient theme, you need to add some special
// characters to the "Card" value. The format is:
// {StartColor},{StopColor},{X1},{Y1},{X2},{Y2},{OffsetStart},{OffsetStop}
// The {LEVEL} placeholder will be replaced with the color of the user's level.
// You also need to name the theme starting with "gradient" like "gradient_my_theme_name".

var Themes = map[string]Theme{
	"default": {
		CardColor:         "#262729",
		HeadlineFontColor: "#F1F5F3",
		BodyFontColor:     "#BB432C",
		IconColor:         "#6795DE",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient": {
		CardColor:         "#262729,#BB432C,40,10,100,0,0,100",
		HeadlineFontColor: "#F1F5F3",
		BodyFontColor:     "#BB432C",
		IconColor:         "#6795DE",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_by_level": {
		CardColor:         "#262729,{LEVEL},40,10,100,0,0,100",
		HeadlineFontColor: "#F1F5F3",
		BodyFontColor:     "#BB432C",
		IconColor:         "#6795DE",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"light": {
		CardColor:         "#fffefe",
		HeadlineFontColor: "#181919",
		BodyFontColor:     "#BB432C",
		IconColor:         "#6795DE",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_light": {
		CardColor:         "#fffefe,#3C7EBB,40,10,100,0,0,100",
		HeadlineFontColor: "#181919",
		BodyFontColor:     "#BB432C",
		IconColor:         "#6795DE",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_light_by_level": {
		CardColor:         "#fffefe,{LEVEL},40,10,100,0,0,100",
		HeadlineFontColor: "#181919",
		BodyFontColor:     "#BB432C",
		IconColor:         "#6795DE",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"dark": {
		CardColor:         "#000000",
		HeadlineFontColor: "#F1F5F3",
		BodyFontColor:     "#BB432C",
		IconColor:         "#6795DE",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_dark": {
		CardColor:         "#000000,#8D8D8D,40,10,100,0,0,100",
		HeadlineFontColor: "#F1F5F3",
		BodyFontColor:     "#BB432C",
		IconColor:         "#6795DE",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_dark_by_level": {
		CardColor:         "#000000,{LEVEL},40,10,100,0,0,100",
		HeadlineFontColor: "#F1F5F3",
		BodyFontColor:     "#BB432C",
		IconColor:         "#6795DE",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"highcontrast": {
		CardColor:         "#000",
		HeadlineFontColor: "#e7f216",
		BodyFontColor:     "#fff",
		IconColor:         "#00ffff",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"midnight_purple": {
		CardColor:         "#000",
		HeadlineFontColor: "#9745f5",
		BodyFontColor:     "#fff",
		IconColor:         "#9f4bff",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_midnight_puple": {
		CardColor:         "#000,#9745f5,40,10,100,0,0,100",
		HeadlineFontColor: "#9745f5",
		BodyFontColor:     "#fff",
		IconColor:         "#9f4bff",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_midnight_puple_by_level": {
		CardColor:         "#000,{LEVEL},40,10,100,0,0,100",
		HeadlineFontColor: "#9745f5",
		BodyFontColor:     "#fff",
		IconColor:         "#9f4bff",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"midnight_blue": {
		CardColor:         "#000",
		HeadlineFontColor: "#0079fa",
		BodyFontColor:     "#0079fa",
		IconColor:         "#0079fa",
		RankBadgeColor:    "#2100fa",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"purple_dark": {
		CardColor:         "#2d2b55",
		HeadlineFontColor: "#c792ea",
		BodyFontColor:     "#a599e9",
		IconColor:         "#b362ff",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_purple_dark": {
		CardColor:         "#2d2b55,#c792ea,40,10,100,0,0,100",
		HeadlineFontColor: "#c792ea",
		BodyFontColor:     "#a599e9",
		IconColor:         "#b362ff",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_purple_dark_by_level": {
		CardColor:         "#2d2b55,{LEVEL},40,10,100,0,0,100",
		HeadlineFontColor: "#c792ea",
		BodyFontColor:     "#a599e9",
		IconColor:         "#b362ff",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"purple_light": {
		CardColor:         "#ffffff",
		HeadlineFontColor: "#a84ee5",
		BodyFontColor:     "#8c64e0",
		IconColor:         "#b362ff",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_purple_light": {
		CardColor:         "#ffffff,#a84ee5,40,10,100,0,0,100",
		HeadlineFontColor: "#a84ee5",
		BodyFontColor:     "#8c64e0",
		IconColor:         "#b362ff",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"gradient_purple_light_by_level": {
		CardColor:         "#ffffff,{LEVEL},40,10,100,0,0,100",
		HeadlineFontColor: "#a84ee5",
		BodyFontColor:     "#8c64e0",
		IconColor:         "#b362ff",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"nightowl": {
		CardColor:         "#011627",
		HeadlineFontColor: "#c792ea",
		BodyFontColor:     "#7fdbca",
		IconColor:         "#ffeb95",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"solarized_dark": {
		CardColor:         "#002b36",
		HeadlineFontColor: "#268bd2",
		BodyFontColor:     "#859900",
		IconColor:         "#b58900",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
	"solarized_light": {
		CardColor:         "#fdf6e3",
		HeadlineFontColor: "#268bd2",
		BodyFontColor:     "#859900",
		IconColor:         "#b58900",
		RankBadgeColor:    "#181919",
		Font:              "Segoe UI, Ubuntu, Sans-Serif",
	},
}
