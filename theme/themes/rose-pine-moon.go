package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func RosePineMoon() *theme.Theme {
	return &theme.Theme{
		Id:     "rose-pine-moon",
		Name:   "Rosé Pine Moon",
		IsDark: true,
		Base: theme.Base{
			Surface:      colors.FromHex("232136"), // hsl(246,24%,17%)
			Text:         colors.FromHex("deddf4"), // hsl(245,50%,91%)
			TextSubtle:   colors.FromHex("918daa"), // hsl(248,15%,61%)
			TextSubtlest: colors.FromHex("6e6986"), // hsl(249,12%,47%)
			Primary:      colors.FromHex("c4a7e7"), // hsl(267,57%,78%)
			Secondary:    colors.FromHex("918daa"), // hsl(248,15%,61%)
			Info:         colors.FromHex("68aeca"), // hsl(197,48%,60%)
			Success:      colors.FromHex("68aeca"), // hsl(197,48%,60%)
			Notice:       colors.FromHex("f6c279"), // hsl(35,88%,72%)
			Warning:      colors.FromHex("e99895"), // hsl(2,66%,75%)
			Danger:       colors.FromHex("eb6f93"), // hsl(343,76%,68%)
		},
		Components: map[string]theme.Base{
			"sidebar": {
				Surface: colors.FromHex("2a273f"), // hsl(247,24%,20%)
			},
			"menu": {
				Surface:      colors.FromHex("383450"), // hsl(248,21%,26%)
				TextSubtle:   colors.FromHex("918daa"), // hsl(248,15%,61%)
				TextSubtlest: colors.FromHex("837e9a"), // hsl(249,12%,55%)
				Border:       colors.FromHex("4c476c"), // hsl(248,21%,35%)
			},
		},
	}
}
