package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func CatppuccinLatte() *theme.Theme {
	return &theme.Theme{
		Id:     "catppuccin-latte",
		Name:   "Catppuccin Latte",
		IsDark: false,
		Base: theme.Base{
			Surface: colors.FromHex("eff1f5"), // hsl(220,23%,95%)
			Text: colors.FromHex("4b4e68"), // hsl(234,16%,35%)
			TextSubtle: colors.FromHex("6c6f84"), // hsl(233,10%,47%)
			TextSubtlest: colors.FromHex("8c8fa1"), // hsl(231,10%,59%)
			Primary: colors.FromHex("8839ef"), // hsl(266,85%,58%)
			Secondary: colors.FromHex("6c6f84"), // hsl(233,10%,47%)
			Info: colors.FromHex("7287fd"), // hsl(231,97%,72%)
			Success: colors.FromHex("17959b"), // hsl(183,74%,35%)
			Notice: colors.FromHex("dd8d1d"), // hsl(35,77%,49%)
			Warning: colors.FromHex("fe640b"), // hsl(22,99%,52%)
			Danger: colors.FromHex("e64754"), // hsl(355,76%,59%)
		},
		Components: map[string]theme.Base{
			"sidebar": {
				Surface: colors.FromHex("e6e9ef"), // hsl(220,22%,92%)
				Border: colors.FromHex("d7dbe5"), // hsl(220,22%,87%)
			},
			"appHeader": {
				Surface: colors.FromHex("dde1e9"), // hsl(220,21%,89%)
				Border: colors.FromHex("d7dbe5"), // hsl(220,22%,87%)
			},
		},
	}
}