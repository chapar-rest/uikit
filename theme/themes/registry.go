package themes

import "github.com/chapar-rest/uikit/theme"

var registry = NewThemeRegistry()

type ThemeRegistry struct {
	themes map[string]*theme.Theme
}

func NewThemeRegistry() *ThemeRegistry {
	return &ThemeRegistry{
		themes: make(map[string]*theme.Theme),
	}
}

func (r *ThemeRegistry) RegisterTheme(name string, theme *theme.Theme) {
	r.themes[name] = theme
}

func (r *ThemeRegistry) GetTheme(name string) *theme.Theme {
	theme, ok := r.themes[name]
	if !ok {
		return nil
	}
	return theme
}

func GetTheme(name string) *theme.Theme {
	return registry.GetTheme(name)
}

func GetThemeById(id string) *theme.Theme {
	for _, theme := range registry.themes {
		if theme.Id == id {
			return theme
		}
	}
	return nil
}

func RegisterTheme(name string, theme *theme.Theme) {
	registry.RegisterTheme(name, theme)
}

func GetAllThemes() map[string]*theme.Theme {
	return registry.themes
}

func init() {
	registry.RegisterTheme("dark", Dark())
	registry.RegisterTheme("light", Light())
	registry.RegisterTheme("dracula", Dracula())
	registry.RegisterTheme("nord", Nord())
	registry.RegisterTheme("gruvbox", Gruvbox())
	registry.RegisterTheme("catppuccin", CatppuccinFrappe())
	registry.RegisterTheme("catppuccin-macchiato", CatppuccinMacchiato())
	registry.RegisterTheme("catppuccin-mocha", CatppuccinMocha())
	registry.RegisterTheme("catppuccin-latte", CatppuccinLatte())
}
