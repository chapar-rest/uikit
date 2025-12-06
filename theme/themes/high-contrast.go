package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func HighContrast() *theme.Theme {
	return &theme.Theme{
		Id:     "high-contrast",
		Name:   "High Contrast Light",
		IsDark: false,
		Base: theme.Base{
			Surface:          colors.FromHex("ffffff"), // white
			SurfaceHighlight: colors.FromHex("e9ecf1"), // hsl(218,24%,93%)
			Text:             colors.FromHex("000000"), // black
			TextSubtle:       colors.FromHex("4e607e"), // hsl(217,24%,40%)
			TextSubtlest:     colors.FromHex("4e607e"), // hsl(217,24%,40%)
			Border:           colors.FromHex("63799c"), // hsl(217,22%,50%)
			Primary:          colors.FromHex("7028c8"), // hsl(267,67%,47%)
			Secondary:        colors.FromHex("72819d"), // hsl(218,18%,53%)
			Info:             colors.FromHex("0068b8"), // hsl(206,100%,36%)
			Success:          colors.FromHex("00854d"), // hsl(155,100%,26%)
			Notice:           colors.FromHex("9e7700"), // hsl(45,100%,31%)
			Warning:          colors.FromHex("ad5701"), // hsl(30,99%,34%)
			Danger:           colors.FromHex("b2004d"), // hsl(334,100%,35%)
		},
	}
}
