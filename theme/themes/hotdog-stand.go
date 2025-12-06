package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func HotdogStand() *theme.Theme {
	return &theme.Theme{
		Id:     "hotdog-stand",
		Name:   "Hotdog Stand",
		IsDark: true,
		Base: theme.Base{
			Surface:          colors.FromHex("ff0000"), // hsl(0,100%,50%)
			SurfaceHighlight: colors.FromHex("000000"), // hsl(0,0%,0%)
			Text:             colors.FromHex("ffffff"), // hsl(0,0%,100%)
			TextSubtle:       colors.FromHex("ffffff"), // hsl(0,0%,100%)
			TextSubtlest:     colors.FromHex("ffff00"), // hsl(60,100%,50%)
			Border:           colors.FromHex("000000"), // hsl(0,0%,0%)
			Primary:          colors.FromHex("ffff00"), // hsl(60,100%,50%)
			Secondary:        colors.FromHex("ffff00"), // hsl(60,100%,50%)
			Info:             colors.FromHex("ffff00"), // hsl(60,100%,50%)
			Success:          colors.FromHex("ffff00"), // hsl(60,100%,50%)
			Notice:           colors.FromHex("ffff00"), // hsl(60,100%,50%)
			Warning:          colors.FromHex("ffff00"), // hsl(60,100%,50%)
			Danger:           colors.FromHex("ffff00"), // hsl(60,100%,50%)
		},
		Components: map[string]theme.Base{
			"appHeader": {
				Surface:      colors.FromHex("000000"), // hsl(0,0%,0%)
				Text:         colors.FromHex("ffffff"), // hsl(0,0%,100%)
				TextSubtle:   colors.FromHex("ffff00"), // hsl(60,100%,50%)
				TextSubtlest: colors.FromHex("ff0000"), // hsl(0,100%,50%)
			},
			"menu": {
				Surface:          colors.FromHex("000000"), // hsl(0,0%,0%)
				Border:           colors.FromHex("ff0000"), // hsl(0,100%,50%)
				SurfaceHighlight: colors.FromHex("ff0000"), // hsl(0,100%,50%)
				Text:             colors.FromHex("ffffff"), // hsl(0,0%,100%)
				TextSubtle:       colors.FromHex("ffff00"), // hsl(60,100%,50%)
				TextSubtlest:     colors.FromHex("ffff00"), // hsl(60,100%,50%)
			},
			"button": {
				Surface:   colors.FromHex("000000"), // hsl(0,0%,0%)
				Text:      colors.FromHex("ffffff"), // hsl(0,0%,100%)
				Primary:   colors.FromHex("000000"), // hsl(0,0%,0%)
				Secondary: colors.FromHex("ffffff"), // hsl(0,0%,100%)
				Info:      colors.FromHex("000000"), // hsl(0,0%,0%)
				Success:   colors.FromHex("ffff00"), // hsl(60,100%,50%)
				Notice:    colors.FromHex("ffff00"), // hsl(60,100%,50%)
				Warning:   colors.FromHex("000000"), // hsl(0,0%,0%)
				Danger:    colors.FromHex("ff0000"), // hsl(0,100%,50%)
			},
		},
	}
}
