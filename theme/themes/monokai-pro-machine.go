package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func MonokaiProMachine() *theme.Theme {
	return &theme.Theme{
		Id:     "monokai-pro-machine",
		Name:   "Monokai Pro Machine",
		IsDark: true,
		Base: theme.Base{
			Surface: colors.FromHex("273035"), // hsl(200,16%,18%)
			Text: colors.FromHex("e9f1f0"), // hsl(173,24%,93%)
			TextSubtle: colors.FromHex("8b9798"), // hsl(185,6%,57%)
			TextSubtlest: colors.FromHex("6c787a"), // hsl(189,6%,45%)
			Primary: colors.FromHex("baa0f8"), // hsl(258,86%,80%)
			Secondary: colors.FromHex("bac5c4"), // hsl(175,9%,75%)
			Info: colors.FromHex("7ed6f1"), // hsl(194,81%,72%)
			Success: colors.FromHex("a2e57b"), // hsl(98,67%,69%)
			Notice: colors.FromHex("ffec70"), // hsl(52,100%,72%)
			Warning: colors.FromHex("ffb370"), // hsl(28,100%,72%)
			Danger: colors.FromHex("ff6b7c"), // hsl(353,100%,71%)
		},
		Components: map[string]theme.Base{
			"appHeader": {
				Surface: colors.FromHex("1e2629"), // hsl(196,16%,14%)
				Text: colors.FromHex("b2b9bd"), // hsl(202,8%,72%)
				TextSubtle: colors.FromHex("767a7f"), // hsl(213,4%,48%)
				TextSubtlest: colors.FromHex("696d77"), // hsl(223,6%,44%)
			},
			"button": {
				Primary: colors.FromHex("9f7af5"), // hsl(258,86%,72%)
				Secondary: colors.FromHex("a6b5b4"), // hsl(175,9%,68%)
				Info: colors.FromHex("5ecced"), // hsl(194,80%,65%)
				Success: colors.FromHex("8ddf5d"), // hsl(98,67%,62%)
				Notice: colors.FromHex("ffe74d"), // hsl(52,100%,65%)
				Warning: colors.FromHex("ffa04d"), // hsl(28,100%,65%)
				Danger: colors.FromHex("ff475d"), // hsl(353,100%,64%)
			},
		},
	}
}