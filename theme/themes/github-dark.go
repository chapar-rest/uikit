package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func GithubDark() *theme.Theme {
	return &theme.Theme{
		Id:     "github-dark",
		Name:   "GitHub",
		IsDark: true,
		Base: theme.Base{
			Surface: colors.FromHex("0c1117"), // hsl(213,30%,7%)
			SurfaceHighlight: colors.FromHex("1c2126"), // hsl(213,16%,13%)
			Text: colors.FromHex("dbe2eb"), // hsl(212,27%,89%)
			TextSubtle: colors.FromHex("87919b"), // hsl(212,9%,57%)
			TextSubtlest: colors.FromHex("6a717c"), // hsl(217,8%,45%)
			Border: colors.FromHex("161b22"), // hsl(215,21%,11%)
			Primary: colors.FromHex("af89f0"), // hsl(262,78%,74%)
			Secondary: colors.FromHex("757d8a"), // hsl(217,8%,50%)
			Info: colors.FromHex("5696f0"), // hsl(215,84%,64%)
			Success: colors.FromHex("4abf5b"), // hsl(129,48%,52%)
			Notice: colors.FromHex("e0ab48"), // hsl(39,71%,58%)
			Warning: colors.FromHex("ee8244"), // hsl(22,83%,60%)
			Danger: colors.FromHex("f0635c"), // hsl(3,83%,65%)
		},
		Components: map[string]theme.Base{
			"button": {
				Primary: colors.FromHex("a57bef"), // hsl(262,79%,71%)
				Secondary: colors.FromHex("6a717c"), // hsl(217,8%,45%)
				Info: colors.FromHex("438bef"), // hsl(215,84%,60%)
				Success: colors.FromHex("3eb150"), // hsl(129,48%,47%)
				Notice: colors.FromHex("dca132"), // hsl(39,71%,53%)
				Warning: colors.FromHex("ec7632"), // hsl(22,83%,56%)
				Danger: colors.FromHex("ee5149"), // hsl(3,83%,61%)
			},
		},
	}
}