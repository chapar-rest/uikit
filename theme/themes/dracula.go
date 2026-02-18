package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func Dracula() *theme.Theme {
	return &theme.Theme{
		Id:     "dracula",
		Name:   "Dracula",
		IsDark: true,
		Base: theme.Base{
			Surface:          colors.FromHex("272935"), // hsl(231,15%,18%)
			SurfaceHighlight: colors.FromHex("343746"), // hsl(230,15%,24%)
			Text:             colors.FromHex("f8f8f2"), // hsl(60,30%,96%)
			TextSubtle:       colors.FromHex("999db2"), // hsl(232,14%,65%)
			TextSubtlest:     colors.FromHex("6e7291"), // hsl(232,14%,50%)
			Border:           colors.FromHex("343746"), // hsl(230,15%,24%)
			Primary:          colors.FromHex("bf95f9"), // hsl(265,89%,78%)
			Secondary:        colors.FromHex("6071a4"), // hsl(225,27%,51%)
			Info:             colors.FromHex("8be8fd"), // hsl(191,97%,77%)
			Success:          colors.FromHex("52fa7c"), // hsl(135,94%,65%)
			Notice:           colors.FromHex("f1fa89"), // hsl(65,92%,76%)
			Warning:          colors.FromHex("ffb86b"), // hsl(31,100%,71%)
			Danger:           colors.FromHex("ff5757"), // hsl(0,100%,67%)
		},
	}
}
