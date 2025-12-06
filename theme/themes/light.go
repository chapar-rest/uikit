package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func Light() *theme.Theme {
	return &theme.Theme{
		Id:     "light",
		Name:   "Light",
		IsDark: false,
		Base: theme.Base{
			Surface:          colors.FromHex("ffffff"), // #ffffff
			SurfaceHighlight: colors.FromHex("d6dce6"), // #d6dce6
			Text:             colors.FromHex("131820"), // #131820
			TextSubtle:       colors.FromHex("4e607e"), // #4e607e
			TextSubtlest:     colors.FromHex("4e607e"), // #4e607e
			Border:           colors.FromHex("e0e4eb"), // #e0e4eb
			Primary:          colors.FromHex("8b33ff"), // #8b33ff
			Secondary:        colors.FromHex("61759e"), // #61759e
			Info:             colors.FromHex("0074cc"), // #0074cc
			Success:          colors.FromHex("1d9042"), // #1d9042
			Notice:           colors.FromHex("ad8200"), // #ad8200
			Warning:          colors.FromHex("b85c00"), // #b85c00
			Danger:           colors.FromHex("d61f6b"), // #d61f6b
		},
	}
}
