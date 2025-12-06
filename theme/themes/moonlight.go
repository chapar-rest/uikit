package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func Moonlight() *theme.Theme {
	return &theme.Theme{
		Id:     "moonlight",
		Name:   "Moonlight",
		IsDark: true,
		Base: theme.Base{
			Surface: colors.FromHex("212335"), // hsl(234,23%,17%)
			Text: colors.FromHex("d3dcf8"), // hsl(225,71%,90%)
			TextSubtle: colors.FromHex("838cb9"), // hsl(230,28%,62%)
			TextSubtlest: colors.FromHex("51598a"), // hsl(232,26%,43%)
			Primary: colors.FromHex("c5a3ff"), // hsl(262,100%,82%)
			Secondary: colors.FromHex("969ab6"), // hsl(232,18%,65%)
			Info: colors.FromHex("7aadff"), // hsl(217,100%,74%)
			Success: colors.FromHex("3cd7c8"), // hsl(174,66%,54%)
			Notice: colors.FromHex("ffc675"), // hsl(35,100%,73%)
			Warning: colors.FromHex("ff956b"), // hsl(17,100%,71%)
			Danger: colors.FromHex("ff757e"), // hsl(356,100%,73%)
		},
		Components: map[string]theme.Base{
			"appHeader": {
				Surface: colors.FromHex("1d202f"), // hsl(233,23%,15%)
			},
			"sidebar": {
				Surface: colors.FromHex("1d202f"), // hsl(233,23%,15%)
			},
		},
	}
}