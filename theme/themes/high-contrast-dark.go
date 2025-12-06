package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func HighContrastDark() *theme.Theme {
	return &theme.Theme{
		Id:     "high-contrast-dark",
		Name:   "High Contrast Dark",
		IsDark: true,
		Base: theme.Base{
			Surface:          colors.FromHex("000000"), // hsl(0,0%,0%)
			SurfaceHighlight: colors.FromHex("333333"), // hsl(0,0%,20%)
			Text:             colors.FromHex("ffffff"), // hsl(0,0%,100%)
			TextSubtle:       colors.FromHex("e6e6e6"), // hsl(0,0%,90%)
			TextSubtlest:     colors.FromHex("cccccc"), // hsl(0,0%,80%)
			Border:           colors.FromHex("999999"), // hsl(0,0%,60%)
			Primary:          colors.FromHex("d4b2ff"), // hsl(266,100%,85%)
			Secondary:        colors.FromHex("aaa9c6"), // hsl(242,20%,72%)
			Info:             colors.FromHex("a8d7ff"), // hsl(208,100%,83%)
			Success:          colors.FromHex("42ffa1"), // hsl(150,100%,63%)
			Notice:           colors.FromHex("ffe98a"), // hsl(49,100%,77%)
			Warning:          colors.FromHex("ffb675"), // hsl(28,100%,73%)
			Danger:           colors.FromHex("ff94b2"), // hsl(343,100%,79%)
		},
	}
}
