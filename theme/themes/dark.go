package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func Dark() *theme.Theme {
	return &theme.Theme{
		Id:     "dark",
		Name:   "Dark",
		IsDark: true,
		Base: theme.Base{
			Surface:          colors.FromHex("1d1b2c"), // #1d1b2c
			SurfaceHighlight: colors.FromHex("29273f"), // #29273f
			Text:             colors.FromHex("d1d0e2"), // #d1d0e2
			TextSubtle:       colors.FromHex("8481a7"), // #8481a7
			TextSubtlest:     colors.FromHex("625e87"), // #625e87
			Border:           colors.FromHex("29273f"), // #29273f
			Primary:          colors.FromHex("8e53ff"), // #8e53ff
			Secondary:        colors.FromHex("5b4f9e"), // #5b4f9e
			Info:             colors.FromHex("0074cc"), // #0074cc
			Success:          colors.FromHex("35a840"), // #35a840
			Notice:           colors.FromHex("ad8200"), // #ad8200
			Warning:          colors.FromHex("b85c00"), // #b85c00
			Danger:           colors.FromHex("d61f6b"), // #d61f6b
		},
		Components: map[string]theme.Base{
			"button": {
				Primary:   colors.FromHex("8e53ff"), // #8e53ff
				Secondary: colors.FromHex("5b4f9e"), // #5b4f9e
				Info:      colors.FromHex("0074cc"), // #0074cc
				Success:   colors.FromHex("35a840"), // #35a840
				Notice:    colors.FromHex("ad8200"), // #ad8200
				Warning:   colors.FromHex("b85c00"), // #b85c00
				Danger:    colors.FromHex("d61f6b"), // #d61f6b
			},
			"dialog": {
				Border: colors.FromHex("29273f"), // #29273f
			},
			"treeview": {
				Surface: colors.FromHex("201f32"), // #243238
			},
		},
	}
}
