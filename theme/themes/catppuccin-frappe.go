package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func CatppuccinFrappe() *theme.Theme {
	return &theme.Theme{
		Id:     "catppuccin-frappe",
		Name:   "Catppuccin Frappé",
		IsDark: true,
		Base: theme.Base{
			Surface:      colors.FromHex("292c3d"), // hsl(231,19%,20%)
			Text:         colors.FromHex("c7d1f5"), // hsl(227,70%,87%)
			TextSubtle:   colors.FromHex("a6aece"), // hsl(228,29%,73%)
			TextSubtlest: colors.FromHex("828aa6"), // hsl(227,17%,58%)
			Primary:      colors.FromHex("ca9ee6"), // hsl(277,59%,76%)
			Secondary:    colors.FromHex("b8c0e0"), // hsl(228,39%,80%)
			Info:         colors.FromHex("8ca9ee"), // hsl(222,74%,74%)
			Success:      colors.FromHex("a6d189"), // hsl(96,44%,68%)
			Notice:       colors.FromHex("e5c88f"), // hsl(40,62%,73%)
			Warning:      colors.FromHex("ef9e76"), // hsl(20,79%,70%)
			Danger:       colors.FromHex("e78384"), // hsl(359,68%,71%)
		},
		Components: map[string]theme.Base{
			"dialog": {
				Surface: colors.FromHex("181825"), // hsl(240,21%,12%)
			},
			"sidebar": {
				Surface: colors.FromHex("303446"), // hsl(229,19%,23%)
				Border:  colors.FromHex("383d52"), // hsl(229,19%,27%)
			},
			"appHeader": {
				Surface: colors.FromHex("232634"), // hsl(229,20%,17%)
				Border:  colors.FromHex("33384c"), // hsl(229,20%,25%)
			},
			"button": {
				Primary:   colors.FromHex("b97dde"), // hsl(277,59%,68%)
				Secondary: colors.FromHex("9ca7d3"), // hsl(228,39%,72%)
				Info:      colors.FromHex("6d92e9"), // hsl(222,74%,67%)
				Success:   colors.FromHex("93c770"), // hsl(96,44%,61%)
				Notice:    colors.FromHex("deba73"), // hsl(40,62%,66%)
				Warning:   colors.FromHex("eb8856"), // hsl(20,79%,63%)
				Danger:    colors.FromHex("e26567"), // hsl(359,68%,64%)
			},
		},
	}
}
