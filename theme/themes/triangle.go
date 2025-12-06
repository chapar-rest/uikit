package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func Triangle() *theme.Theme {
	return &theme.Theme{
		Id:     "triangle",
		Name:   "Triangle",
		IsDark: true,
		Base: theme.Base{
			Surface:          colors.FromHex("000000"), // rgb(0,0,0)
			SurfaceHighlight: colors.FromHex("151515"), // rgb(21,21,21)
			Text:             colors.FromHex("ededed"), // rgb(237,237,237)
			TextSubtle:       colors.FromHex("a1a1a1"), // rgb(161,161,161)
			TextSubtlest:     colors.FromHex("737373"), // rgb(115,115,115)
			Border:           colors.FromHex("1f1f1f"), // rgb(31,31,31)
			Primary:          colors.FromHex("c472fb"), // rgb(196,114,251)
			Secondary:        colors.FromHex("a1a1a1"), // rgb(161,161,161)
			Info:             colors.FromHex("47a8ff"), // rgb(71,168,255)
			Success:          colors.FromHex("00ca51"), // rgb(0,202,81)
			Notice:           colors.FromHex("ffaf00"), // rgb(255,175,0)
			Warning:          colors.FromHex("ff4c8d"), // #FF4C8D
			Danger:           colors.FromHex("fd495a"), // #fd495a
		},
		Components: map[string]theme.Base{
			"dialog": {
				Surface: colors.FromHex("0a0a0a"), // rgb(10,10,10)
				Border:  colors.FromHex("1f1f1f"), // rgb(31,31,31)
			},
			"sidebar": {
				Border: colors.FromHex("1f1f1f"), // rgb(31,31,31)
			},
			"appHeader": {
				Surface: colors.FromHex("0a0a0a"), // rgb(10,10,10)
				Border:  colors.FromHex("1f1f1f"), // rgb(31,31,31)
			},
		},
	}
}
