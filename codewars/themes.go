package codewars

type Theme struct {
	Card           string
	HeadlineFont   string
	BodyFont       string
	Icon           string
	RankBadge      string
	HasGradient    bool
	GradientValues GradientValues
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

var Themes = map[string]Theme{
	"default": {
		Card:         "#262729",
		HeadlineFont: "#F1F5F3",
		BodyFont:     "#BB432C",
		Icon:         "#6795DE",
		RankBadge:    "#181919",
	},
	"gradient": {
		Card:         "#262729,#BB432C,40,10,100,0,0,100",
		HeadlineFont: "#F1F5F3",
		BodyFont:     "#BB432C",
		Icon:         "#6795DE",
		RankBadge:    "#181919",
	},
	"light": {
		Card:         "#fffefe",
		HeadlineFont: "#181919",
		BodyFont:     "#BB432C",
		Icon:         "#6795DE",
		RankBadge:    "#181919",
	},
	"gradient_light": {
		Card:         "#fffefe,#3C7EBB,40,10,100,0,0,100",
		HeadlineFont: "#181919",
		BodyFont:     "#BB432C",
		Icon:         "#6795DE",
		RankBadge:    "#181919",
	},
	"gradient_light_by_level": {
		Card:         "#fffefe,{LEVEL},40,10,100,0,0,100",
		HeadlineFont: "#181919",
		BodyFont:     "#BB432C",
		Icon:         "#6795DE",
		RankBadge:    "#181919",
	},
	"dark": {
		Card:         "#000000",
		HeadlineFont: "#F1F5F3",
		BodyFont:     "#BB432C",
		Icon:         "#6795DE",
		RankBadge:    "#181919",
	},
	"gradient_dark": {
		Card:         "#000000,#8D8D8D,40,10,100,0,0,100",
		HeadlineFont: "#F1F5F3",
		BodyFont:     "#BB432C",
		Icon:         "#6795DE",
		RankBadge:    "#181919",
	},
	"gradient_dark_by_level": {
		Card:         "#000000,{LEVEL},40,10,100,0,0,100",
		HeadlineFont: "#F1F5F3",
		BodyFont:     "#BB432C",
		Icon:         "#6795DE",
		RankBadge:    "#181919",
	},
	"highcontrast": {
		Card:         "#000",
		HeadlineFont: "#e7f216",
		BodyFont:     "#fff",
		Icon:         "#00ffff",
		RankBadge:    "#181919",
	},
	"midnight_purple": {
		Card:         "#000",
		HeadlineFont: "#9745f5",
		BodyFont:     "#fff",
		Icon:         "#9f4bff",
		RankBadge:    "#181919",
	},
	"gradient_midnight_puple": {
		Card:         "#000,#9745f5,40,10,100,0,0,100",
		HeadlineFont: "#9745f5",
		BodyFont:     "#fff",
		Icon:         "#9f4bff",
		RankBadge:    "#181919",
	},
	"gradient_midnight_puple_by_level": {
		Card:         "#000,{LEVEL},40,10,100,0,0,100",
		HeadlineFont: "#9745f5",
		BodyFont:     "#fff",
		Icon:         "#9f4bff",
		RankBadge:    "#181919",
	},
	"midnight_blue": {
		Card:         "#000",
		HeadlineFont: "#0079fa",
		BodyFont:     "#0079fa",
		Icon:         "#0079fa",
		RankBadge:    "#2100fa",
	},
	"purple_dark": {
		Card:         "#2d2b55",
		HeadlineFont: "#c792ea",
		BodyFont:     "#a599e9",
		Icon:         "#b362ff",
		RankBadge:    "#181919",
	},
	"gradient_purple_dark": {
		Card:         "#2d2b55,#c792ea,40,10,100,0,0,100",
		HeadlineFont: "#c792ea",
		BodyFont:     "#a599e9",
		Icon:         "#b362ff",
		RankBadge:    "#181919",
	},
	"gradient_purple_dark_by_level": {
		Card:         "#2d2b55,{LEVEL},40,10,100,0,0,100",
		HeadlineFont: "#c792ea",
		BodyFont:     "#a599e9",
		Icon:         "#b362ff",
		RankBadge:    "#181919",
	},
	"purple_light": {
		Card:         "#ffffff",
		HeadlineFont: "#a84ee5",
		BodyFont:     "#8c64e0",
		Icon:         "#b362ff",
		RankBadge:    "#181919",
	},
	"gradient_purple_light": {
		Card:         "#ffffff,#a84ee5,40,10,100,0,0,100",
		HeadlineFont: "#a84ee5",
		BodyFont:     "#8c64e0",
		Icon:         "#b362ff",
		RankBadge:    "#181919",
	},
	"gradient_purple_light_by_level": {
		Card:         "#ffffff,{LEVEL},40,10,100,0,0,100",
		HeadlineFont: "#a84ee5",
		BodyFont:     "#8c64e0",
		Icon:         "#b362ff",
		RankBadge:    "#181919",
	},
	"nightowl": {
		Card:         "#011627",
		HeadlineFont: "#c792ea",
		BodyFont:     "#7fdbca",
		Icon:         "#ffeb95",
		RankBadge:    "#181919",
	},
	"solarized_dark": {
		Card:         "#002b36",
		HeadlineFont: "#268bd2",
		BodyFont:     "#859900",
		Icon:         "#b58900",
		RankBadge:    "#181919",
	},
	"solarized_light": {
		Card:         "#fdf6e3",
		HeadlineFont: "#268bd2",
		BodyFont:     "#859900",
		Icon:         "#b58900",
		RankBadge:    "#181919",
	},
}
