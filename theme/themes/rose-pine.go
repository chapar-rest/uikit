package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func RosePine() *theme.Theme {
	return &theme.Theme{
		Id:     "rose-pine",
		Name:   "Rosé Pine",
		IsDark: true,
		Base: theme.Base{
			Surface:      colors.FromHex("1a1825"), // hsl(249,22%,12%)
			Text:         colors.FromHex("deddf4"), // hsl(245,50%,91%)
			TextSubtle:   colors.FromHex("918daa"), // hsl(248,15%,61%)
			TextSubtlest: colors.FromHex("6e6986"), // hsl(249,12%,47%)
			Primary:      colors.FromHex("c4a7e7"), // hsl(267,57%,78%)
			Secondary:    colors.FromHex("6e6986"), // hsl(249,12%,47%)
			Info:         colors.FromHex("67abcb"), // hsl(199,49%,60%)
			Success:      colors.FromHex("9dd8d8"), // hsl(180,43%,73%)
			Notice:       colors.FromHex("f6c279"), // hsl(35,88%,72%)
			Warning:      colors.FromHex("f1a3a2"), // hsl(1,74%,79%)
			Danger:       colors.FromHex("eb6f93"), // hsl(343,76%,68%)
		},
		Components: map[string]theme.Base{
			"sidebar": {
				Surface: colors.FromHex("201d2f"), // hsl(247,23%,15%)
			},
			"menu": {
				Surface:      colors.FromHex("383450"), // hsl(248,21%,26%)
				TextSubtle:   colors.FromHex("9f9bb5"), // hsl(248,15%,66%)
				TextSubtlest: colors.FromHex("7a7693"), // hsl(249,12%,52%)
				Border:       colors.FromHex("4c476c"), // hsl(248,21%,35%)
			},
		},
	}
}
