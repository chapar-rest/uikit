package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func MonokaiProSpectrum() *theme.Theme {
	return &theme.Theme{
		Id:     "monokai-pro-spectrum",
		Name:   "Monokai Pro Spectrum",
		IsDark: true,
		Base: theme.Base{
			Surface: colors.FromHex("212121"), // hsl(0,0%,13%)
			Text: colors.FromHex("f6f0ff"), // hsl(266,100%,97%)
			TextSubtle: colors.FromHex("b9b5bf"), // hsl(264,7%,73%)
			TextSubtlest: colors.FromHex("8c8990"), // hsl(266,3%,55%)
			Primary: colors.FromHex("968ce3"), // hsl(247,61%,72%)
			Secondary: colors.FromHex("b9b5bf"), // hsl(264,7%,73%)
			Info: colors.FromHex("5bd4e6"), // hsl(188,74%,63%)
			Success: colors.FromHex("79d78e"), // hsl(133,54%,66%)
			Notice: colors.FromHex("fce564"), // hsl(51,96%,69%)
			Warning: colors.FromHex("fd9453"), // hsl(23,98%,66%)
			Danger: colors.FromHex("fc5f8b"), // hsl(343,96%,68%)
		},
		Components: map[string]theme.Base{
			"appHeader": {
				Surface: colors.FromHex("1a1a1a"), // hsl(0,0%,10%)
				Text: colors.FromHex("b9b5bf"), // hsl(264,7%,73%)
				TextSubtle: colors.FromHex("8c8990"), // hsl(266,3%,55%)
				TextSubtlest: colors.FromHex("68666b"), // hsl(264,2%,41%)
			},
			"button": {
				Primary: colors.FromHex("7c6fdc"), // hsl(247,61%,65%)
				Secondary: colors.FromHex("a7a2ae"), // hsl(264,7%,66%)
				Info: colors.FromHex("40cde2"), // hsl(188,74%,57%)
				Success: colors.FromHex("5ecf76"), // hsl(133,54%,59%)
				Notice: colors.FromHex("fbdf41"), // hsl(51,96%,62%)
				Warning: colors.FromHex("fd7f30"), // hsl(23,98%,59%)
				Danger: colors.FromHex("fb3c72"), // hsl(343,96%,61%)
			},
		},
	}
}