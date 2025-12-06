package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func CatppuccinMacchiato() *theme.Theme {
	return &theme.Theme{
		Id:     "catppuccin-macchiato",
		Name:   "Catppuccin Macchiato",
		IsDark: true,
		Base: theme.Base{
			Surface:      colors.FromHex("1d202f"), // hsl(233,23%,15%)
			Text:         colors.FromHex("ccd5f5"), // hsl(227,68%,88%)
			TextSubtle:   colors.FromHex("a4adcb"), // hsl(227,27%,72%)
			TextSubtlest: colors.FromHex("8187a2"), // hsl(228,15%,57%)
			Primary:      colors.FromHex("c8a2f6"), // hsl(267,83%,80%)
			Secondary:    colors.FromHex("b8c0e0"), // hsl(228,39%,80%)
			Info:         colors.FromHex("8aaef4"), // hsl(220,83%,75%)
			Success:      colors.FromHex("a6da95"), // hsl(105,48%,72%)
			Notice:       colors.FromHex("eed4a0"), // hsl(40,70%,78%)
			Warning:      colors.FromHex("f5a87f"), // hsl(21,86%,73%)
			Danger:       colors.FromHex("ed8796"), // hsl(351,74%,73%)
		},
		Components: map[string]theme.Base{
			"dialog": {
				Surface: colors.FromHex("181825"), // hsl(240,21%,12%)
			},
			"sidebar": {
				Surface: colors.FromHex("232638"), // hsl(232,23%,18%)
				Border:  colors.FromHex("2b2f45"), // hsl(231,23%,22%)
			},
			"appHeader": {
				Surface: colors.FromHex("181926"), // hsl(236,23%,12%)
				Border:  colors.FromHex("292b42"), // hsl(236,23%,21%)
			},
			"button": {
				Primary:   colors.FromHex("b27df2"), // hsl(267,82%,72%)
				Secondary: colors.FromHex("9ca7d3"), // hsl(228,39%,72%)
				Info:      colors.FromHex("6a97f1"), // hsl(220,83%,68%)
				Success:   colors.FromHex("90d17b"), // hsl(105,48%,65%)
				Notice:    colors.FromHex("e8c47d"), // hsl(40,70%,70%)
				Warning:   colors.FromHex("f3925e"), // hsl(21,86%,66%)
				Danger:    colors.FromHex("e8687b"), // hsl(351,74%,66%)
			},
		},
	}
}
