package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func Relaxing() *theme.Theme {
	return &theme.Theme{
		Id:     "relaxing",
		Name:   "Relaxing",
		IsDark: true,
		Base: theme.Base{
			Surface:   colors.FromHex("2a1d3a"), // hsl(267,33%,17%)
			Text:      colors.FromHex("ece1f5"), // hsl(275,49%,92%)
			Primary:   colors.FromHex("caa6f7"), // hsl(267,84%,81%)
			Secondary: colors.FromHex("bac2de"), // hsl(227,35%,80%)
			Info:      colors.FromHex("89b5fa"), // hsl(217,92%,76%)
			Success:   colors.FromHex("a6e3a1"), // hsl(115,54%,76%)
			Notice:    colors.FromHex("f9e1ae"), // hsl(41,86%,83%)
			Warning:   colors.FromHex("fab285"), // hsl(23,92%,75%)
			Danger:    colors.FromHex("f38ca9"), // hsl(343,81%,75%)
		},
	}
}
