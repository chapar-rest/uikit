package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func MonokaiProClassic() *theme.Theme {
	return &theme.Theme{
		Id:     "monokai-pro-classic",
		Name:   "Monokai Pro Classic",
		IsDark: true,
		Base: theme.Base{
			Surface: colors.FromHex("282923"), // hsl(70,8%,15%)
			Text: colors.FromHex("fdfff0"), // hsl(69,100%,97%)
			TextSubtle: colors.FromHex("bfc0b4"), // hsl(65,9%,73%)
			TextSubtlest: colors.FromHex("909188"), // hsl(66,4%,55%)
			Primary: colors.FromHex("ac80ff"), // hsl(261,100%,75%)
			Secondary: colors.FromHex("b2b9bd"), // hsl(202,8%,72%)
			Info: colors.FromHex("67d8ef"), // hsl(190,81%,67%)
			Success: colors.FromHex("a6e22c"), // hsl(80,76%,53%)
			Notice: colors.FromHex("e7db74"), // hsl(54,70%,68%)
			Warning: colors.FromHex("fd9621"), // hsl(32,98%,56%)
			Danger: colors.FromHex("f92472"), // hsl(338,95%,56%)
		},
		Components: map[string]theme.Base{
			"appHeader": {
				Surface: colors.FromHex("1e1f1a"), // hsl(72,9%,11%)
				Text: colors.FromHex("b2b9bd"), // hsl(202,8%,72%)
				TextSubtle: colors.FromHex("767a7f"), // hsl(213,4%,48%)
				TextSubtlest: colors.FromHex("696d77"), // hsl(223,6%,44%)
			},
			"button": {
				Primary: colors.FromHex("955cff"), // hsl(261,100%,68%)
				Secondary: colors.FromHex("9fa8ad"), // hsl(202,8%,65%)
				Info: colors.FromHex("46d0ec"), // hsl(190,81%,60%)
				Success: colors.FromHex("99d71d"), // hsl(80,76%,48%)
				Notice: colors.FromHex("e2d455"), // hsl(54,71%,61%)
				Warning: colors.FromHex("fc8803"), // hsl(32,98%,50%)
				Danger: colors.FromHex("f9065f"), // hsl(338,95%,50%)
			},
		},
	}
}