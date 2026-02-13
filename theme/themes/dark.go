package themes

import (
	"gioui.org/layout"
	"gioui.org/unit"
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

			Radii: theme.Radii{
				None:            unit.Dp(0),
				Small:           unit.Dp(4),
				Medium:          unit.Dp(8),
				Large:           unit.Dp(16),
				ExtraLarge:      unit.Dp(24),
				ExtraExtraLarge: unit.Dp(32),
			},
			Insets: theme.Insets{
				None:            layout.Inset{},
				Small:           layout.Inset{Top: 4, Bottom: 4, Left: 4, Right: 4},
				Medium:          layout.Inset{Top: 8, Bottom: 8, Left: 8, Right: 8},
				Large:           layout.Inset{Top: 16, Bottom: 16, Left: 16, Right: 16},
				ExtraLarge:      layout.Inset{Top: 24, Bottom: 24, Left: 24, Right: 24},
				ExtraExtraLarge: layout.Inset{Top: 32, Bottom: 32, Left: 32, Right: 32},
			},
			Sizes: theme.Sizes{
				None:            theme.Size{Height: 0, Width: 0},
				Small:           theme.Size{Height: 24, Width: 24},
				Medium:          theme.Size{Height: 32, Width: 32},
				Large:           theme.Size{Height: 40, Width: 40},
				ExtraLarge:      theme.Size{Height: 48, Width: 48},
				ExtraExtraLarge: theme.Size{Height: 56, Width: 56},
			},
		},
		Components: map[string]theme.Base{
			"button": {
				Primary:   colors.FromHex("d1d0e2"), // #d1d0e2
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
