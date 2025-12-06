package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func MonokaiProOctagon() *theme.Theme {
	return &theme.Theme{
		Id:     "monokai-pro-octagon",
		Name:   "Monokai Pro Octagon",
		IsDark: true,
		Base: theme.Base{
			Surface: colors.FromHex("282a39"), // hsl(233,18%,19%)
			Text: colors.FromHex("e9f1f0"), // hsl(173,24%,93%)
			TextSubtle: colors.FromHex("b2b9bd"), // hsl(202,8%,72%)
			TextSubtlest: colors.FromHex("767a7f"), // hsl(213,4%,48%)
			Primary: colors.FromHex("c39cc9"), // hsl(292,30%,70%)
			Secondary: colors.FromHex("b2b9bd"), // hsl(202,8%,72%)
			Info: colors.FromHex("9dd2bc"), // hsl(155,37%,72%)
			Success: colors.FromHex("b9d760"), // hsl(75,60%,61%)
			Notice: colors.FromHex("ffd86b"), // hsl(44,100%,71%)
			Warning: colors.FromHex("ff9a5c"), // hsl(23,100%,68%)
			Danger: colors.FromHex("ff667a"), // hsl(352,100%,70%)
		},
		Components: map[string]theme.Base{
			"appHeader": {
				Surface: colors.FromHex("1d1e2a"), // hsl(235,18%,14%)
				Text: colors.FromHex("b2b9bd"), // hsl(202,8%,72%)
				TextSubtle: colors.FromHex("767a7f"), // hsl(213,4%,48%)
				TextSubtlest: colors.FromHex("696d77"), // hsl(223,6%,44%)
			},
			"button": {
				Primary: colors.FromHex("b388b9"), // hsl(292,26%,63%)
				Secondary: colors.FromHex("a0a8ac"), // hsl(201,7%,65%)
				Info: colors.FromHex("88c3ab"), // hsl(155,33%,65%)
				Success: colors.FromHex("abca4e"), // hsl(75,54%,55%)
				Notice: colors.FromHex("f6ca51"), // hsl(44,90%,64%)
				Warning: colors.FromHex("f58742"), // hsl(23,90%,61%)
				Danger: colors.FromHex("f64c62"), // hsl(352,90%,63%)
			},
		},
	}
}