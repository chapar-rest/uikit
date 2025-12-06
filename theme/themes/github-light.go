package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func GithubLight() *theme.Theme {
	return &theme.Theme{
		Id:     "github-light",
		Name:   "GitHub",
		IsDark: false,
		Base: theme.Base{
			Surface: colors.FromHex("ffffff"), // hsl(0,0%,100%)
			SurfaceHighlight: colors.FromHex("ebf0f4"), // hsl(210,29%,94%)
			Text: colors.FromHex("1f2328"), // hsl(213,13%,14%)
			TextSubtle: colors.FromHex("646d78"), // hsl(212,9%,43%)
			TextSubtlest: colors.FromHex("838e95"), // hsl(203,8%,55%)
			Border: colors.FromHex("e8ebee"), // hsl(210,15%,92%)
			Primary: colors.FromHex("814edf"), // hsl(261,69%,59%)
			Secondary: colors.FromHex("6e7781"), // hsl(212,8%,47%)
			Info: colors.FromHex("0a73eb"), // hsl(212,92%,48%)
			Success: colors.FromHex("1c873a"), // hsl(137,66%,32%)
			Notice: colors.FromHex("cc8800"), // hsl(40,100%,40%)
			Warning: colors.FromHex("e05a00"), // hsl(24,100%,44%)
			Danger: colors.FromHex("d1232f"), // hsl(356,71%,48%)
		},
		Components: map[string]theme.Base{},
	}
}