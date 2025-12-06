package theme

import (
	"image/color"

	"gioui.org/text"
	"gioui.org/widget/material"
)

type Theme struct {
	Id         string
	Name       string
	IsDark     bool
	Base       Base
	Components map[string]Base

	// material theme is used to create the material components
	// colors will be taken from the base theme
	material *material.Theme
}

type Base struct {
	Surface          color.NRGBA
	SurfaceHighlight color.NRGBA
	Text             color.NRGBA
	TextSubtle       color.NRGBA
	TextSubtlest     color.NRGBA
	Border           color.NRGBA
	Primary          color.NRGBA
	Secondary        color.NRGBA
	Info             color.NRGBA
	Success          color.NRGBA
	Notice           color.NRGBA
	Warning          color.NRGBA
	Danger           color.NRGBA
}

func (b *Theme) WithFonts(fonts []text.FontFace) *Theme {
	if b.material == nil {
		b.material = material.NewTheme()
	}
	b.material.Shaper = text.NewShaper(text.WithCollection(fonts))
	return b
}

func (b *Theme) Material() *material.Theme {
	if b.material == nil {
		b.material = material.NewTheme()
	}
	b.material.Palette.Bg = b.Base.Surface
	b.material.Palette.Fg = b.Base.Text
	b.material.Palette.ContrastBg = b.Base.SurfaceHighlight
	b.material.Palette.ContrastFg = b.Base.Text
	return b.material
}

func (b *Theme) RegisterComponent(id string, component Base) {
	b.Components[id] = component
}
