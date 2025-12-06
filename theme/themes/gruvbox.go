package themes

import (
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

func Gruvbox() *theme.Theme {
	return &theme.Theme{
		Id:     "gruvbox",
		Name:   "Gruvbox",
		IsDark: true,
		Base: theme.Base{
			Surface:          colors.FromHex("292929"), // hsl(0,0%,16%)
			SurfaceHighlight: colors.FromHex("32302f"), // hsl(20,3%,19%)
			Text:             colors.FromHex("f9f5d7"), // hsl(53,74%,91%)
			TextSubtle:       colors.FromHex("bdaf93"), // hsl(39,24%,66%)
			TextSubtlest:     colors.FromHex("918273"), // hsl(30,12%,51%)
			Primary:          colors.FromHex("d4879c"), // hsl(344,47%,68%)
			Secondary:        colors.FromHex("83a598"), // hsl(157,16%,58%)
			Info:             colors.FromHex("8ec07c"), // hsl(104,35%,62%)
			Success:          colors.FromHex("b8ba26"), // hsl(61,66%,44%)
			Notice:           colors.FromHex("fabd2e"), // hsl(42,95%,58%)
			Warning:          colors.FromHex("fe811b"), // hsl(27,99%,55%)
			Danger:           colors.FromHex("fb4632"), // hsl(6,96%,59%)
		},
	}
}
