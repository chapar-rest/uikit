package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func CatppuccinMocha() *theme.Theme {
	return &theme.Theme{
		Id:     "catppuccin-mocha",
		Name:   "Catppuccin Mocha",
		IsDark: true,
		Base: theme.Base{
			Surface:      colors.FromHex("181825"), // hsl(240,21%,12%)
			Text:         colors.FromHex("cdd6f4"), // hsl(226,64%,88%)
			TextSubtle:   colors.FromHex("a6adc9"), // hsl(228,24%,72%)
			TextSubtlest: colors.FromHex("7d829b"), // hsl(230,13%,55%)
			Primary:      colors.FromHex("c8a2f6"), // hsl(267,83%,80%)
			Secondary:    colors.FromHex("bac2de"), // hsl(227,35%,80%)
			Info:         colors.FromHex("89b5fa"), // hsl(217,92%,76%)
			Success:      colors.FromHex("a6e3a1"), // hsl(115,54%,76%)
			Notice:       colors.FromHex("f9e1ae"), // hsl(41,86%,83%)
			Warning:      colors.FromHex("fab285"), // hsl(23,92%,75%)
			Danger:       colors.FromHex("f38ca9"), // hsl(343,81%,75%)
		},
		Components: map[string]theme.Base{
			"dialog": {
				Surface: colors.FromHex("181825"), // hsl(240,21%,12%)
			},
			"sidebar": {
				Surface: colors.FromHex("1e1e2e"), // hsl(240,21%,15%)
				Border:  colors.FromHex("26263b"), // hsl(240,21%,19%)
			},
			"appHeader": {
				Surface: colors.FromHex("12121c"), // hsl(240,23%,9%)
				Border:  colors.FromHex("242438"), // hsl(240,22%,18%)
			},
			"button": {
				Primary:   colors.FromHex("a06ae2"), // hsl(267,67%,65%)
				Secondary: colors.FromHex("8995bd"), // hsl(227,28%,64%)
				Info:      colors.FromHex("528ae5"), // hsl(217,74%,61%)
				Success:   colors.FromHex("78c671"), // hsl(115,43%,61%)
				Notice:    colors.FromHex("e4be6c"), // hsl(41,69%,66%)
				Warning:   colors.FromHex("e4874e"), // hsl(23,74%,60%)
				Danger:    colors.FromHex("db577c"), // hsl(343,65%,60%)
			},
		},
	}
}
