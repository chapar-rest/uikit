package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func Nord() *theme.Theme {
	return &theme.Theme{
		Id:     "nord",
		Name:   "Nord",
		IsDark: true,
		Base: theme.Base{
			Surface: colors.FromHex("2f3541"), // hsl(220,16%,22%)
			SurfaceHighlight: colors.FromHex("3d4451"), // hsl(220,14%,28%)
			Text: colors.FromHex("e8ebf2"), // hsl(220,28%,93%)
			TextSubtle: colors.FromHex("dfe3ec"), // hsl(220,26%,90%)
			TextSubtlest: colors.FromHex("d3d8e4"), // hsl(220,24%,86%)
			Primary: colors.FromHex("8ebfcc"), // hsl(193,38%,68%)
			Secondary: colors.FromHex("81a1c1"), // hsl(210,34%,63%)
			Info: colors.FromHex("9cc4c0"), // hsl(174,25%,69%)
			Success: colors.FromHex("a9bf92"), // hsl(89,26%,66%)
			Notice: colors.FromHex("e8c98d"), // hsl(40,66%,73%)
			Warning: colors.FromHex("cf9077"), // hsl(17,48%,64%)
			Danger: colors.FromHex("bf5f6a"), // hsl(353,43%,56%)
		},
	}
}