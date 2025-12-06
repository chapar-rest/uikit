package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func MonokaiPro() *theme.Theme {
	return &theme.Theme{
		Id:     "monokai-pro",
		Name:   "Monokai Pro",
		IsDark: true,
		Base: theme.Base{
			Surface: colors.FromHex("2c292e"), // hsl(285,5%,17%)
			Text: colors.FromHex("fbfbf9"), // hsl(60,25%,98%)
			TextSubtle: colors.FromHex("c0bfbf"), // hsl(0,1%,75%)
			TextSubtlest: colors.FromHex("919191"), // hsl(300,0%,57%)
			Primary: colors.FromHex("aa9cf2"), // hsl(250,77%,78%)
			Secondary: colors.FromHex("c0bfbf"), // hsl(0,1%,75%)
			Info: colors.FromHex("78dde8"), // hsl(186,71%,69%)
			Success: colors.FromHex("a8db75"), // hsl(90,59%,66%)
			Notice: colors.FromHex("ffd966"), // hsl(45,100%,70%)
			Warning: colors.FromHex("fc9a69"), // hsl(20,96%,70%)
			Danger: colors.FromHex("ff6188"), // hsl(345,100%,69%)
		},
		Components: map[string]theme.Base{
			"appHeader": {
				Surface: colors.FromHex("231f23"), // hsl(300,5%,13%)
				Text: colors.FromHex("c0bfbf"), // hsl(0,1%,75%)
				TextSubtle: colors.FromHex("919191"), // hsl(300,0%,57%)
				TextSubtlest: colors.FromHex("716f71"), // hsl(300,1%,44%)
			},
			"button": {
				Primary: colors.FromHex("8b78ed"), // hsl(250,77%,70%)
				Secondary: colors.FromHex("aeadad"), // hsl(0,1%,68%)
				Info: colors.FromHex("59d5e3"), // hsl(186,71%,62%)
				Success: colors.FromHex("96d459"), // hsl(90,59%,59%)
				Notice: colors.FromHex("ffd042"), // hsl(45,100%,63%)
				Warning: colors.FromHex("fb8246"), // hsl(20,96%,63%)
				Danger: colors.FromHex("ff3d6e"), // hsl(345,100%,62%)
			},
		},
	}
}