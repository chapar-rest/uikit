package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func RosePineDawn() *theme.Theme {
	return &theme.Theme{
		Id:     "rose-pine-dawn",
		Name:   "Rosé Pine Dawn",
		IsDark: false,
		Base: theme.Base{
			Surface:          colors.FromHex("faf3eb"), // hsl(32,57%,95%)
			SurfaceHighlight: colors.FromHex("f3ece7"), // hsl(25,35%,93%)
			Text:             colors.FromHex("585379"), // hsl(248,19%,40%)
			TextSubtle:       colors.FromHex("7a7693"), // hsl(248,12%,52%)
			TextSubtlest:     colors.FromHex("9893a5"), // hsl(257,9%,61%)
			Border:           colors.FromHex("dfd9d8"), // hsl(10,9%,86%)
			Primary:          colors.FromHex("9071ad"), // hsl(271,27%,56%)
			Secondary:        colors.FromHex("6e6986"), // hsl(249,12%,47%)
			Info:             colors.FromHex("2c708c"), // hsl(197,52%,36%)
			Success:          colors.FromHex("4f8d96"), // hsl(188,31%,45%)
			Notice:           colors.FromHex("cd882d"), // hsl(34,64%,49%)
			Warning:          colors.FromHex("ce7b78"), // hsl(2,47%,64%)
			Danger:           colors.FromHex("b4647b"), // hsl(343,35%,55%)
		},
		Components: map[string]theme.Base{
			"sidebar": {
				Border: colors.FromHex("e9e4e2"), // hsl(20,12%,90%)
			},
			"appHeader": {
				Border: colors.FromHex("e9e4e2"), // hsl(20,12%,90%)
			},
			"input": {
				Border: colors.FromHex("dfd9d8"), // hsl(10,9%,86%)
			},
			"dialog": {
				Border: colors.FromHex("e9e4e2"), // hsl(20,12%,90%)
			},
			"menu": {
				Surface: colors.FromHex("f3eae2"), // hsl(28,40%,92%)
				Border:  colors.FromHex("dfd9d8"), // hsl(10,9%,86%)
			},
		},
	}
}
