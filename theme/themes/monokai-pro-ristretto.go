package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func MonokaiProRistretto() *theme.Theme {
	return &theme.Theme{
		Id:     "monokai-pro-ristretto",
		Name:   "Monokai Pro Ristretto",
		IsDark: true,
		Base: theme.Base{
			Surface: colors.FromHex("2c2525"), // hsl(0,9%,16%)
			Text: colors.FromHex("fff0f2"), // hsl(351,100%,97%)
			TextSubtle: colors.FromHex("c3b7b8"), // hsl(355,9%,74%)
			TextSubtlest: colors.FromHex("938a8b"), // hsl(354,4%,56%)
			Primary: colors.FromHex("a8a9eb"), // hsl(239,63%,79%)
			Secondary: colors.FromHex("c3b7b8"), // hsl(355,9%,74%)
			Info: colors.FromHex("86dacc"), // hsl(170,53%,69%)
			Success: colors.FromHex("acda77"), // hsl(88,57%,66%)
			Notice: colors.FromHex("f9cc6c"), // hsl(41,92%,70%)
			Warning: colors.FromHex("f48e71"), // hsl(13,85%,70%)
			Danger: colors.FromHex("fd6884"), // hsl(349,97%,70%)
		},
		Components: map[string]theme.Base{
			"appHeader": {
				Surface: colors.FromHex("211c1c"), // hsl(0,8%,12%)
				Text: colors.FromHex("c3b7b8"), // hsl(355,9%,74%)
				TextSubtle: colors.FromHex("938a8b"), // hsl(354,4%,56%)
				TextSubtlest: colors.FromHex("72696a"), // hsl(353,4%,43%)
			},
			"button": {
				Primary: colors.FromHex("8688e4"), // hsl(239,63%,71%)
				Secondary: colors.FromHex("b2a3a5"), // hsl(355,9%,67%)
				Info: colors.FromHex("6bd1c0"), // hsl(170,53%,62%)
				Success: colors.FromHex("9ad25b"), // hsl(88,57%,59%)
				Notice: colors.FromHex("f7c04a"), // hsl(41,92%,63%)
				Warning: colors.FromHex("f27350"), // hsl(13,86%,63%)
				Danger: colors.FromHex("fc4567"), // hsl(349,97%,63%)
			},
		},
	}
}