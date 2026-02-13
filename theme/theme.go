package theme

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
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

	Radii  Radii
	Insets Insets
	Sizes  Sizes
}

type Radii struct {
	None            unit.Dp
	Small           unit.Dp
	Medium          unit.Dp
	Large           unit.Dp
	ExtraLarge      unit.Dp
	ExtraExtraLarge unit.Dp
}

type Insets struct {
	None            layout.Inset
	Small           layout.Inset
	Medium          layout.Inset
	Large           layout.Inset
	ExtraLarge      layout.Inset
	ExtraExtraLarge layout.Inset
}

type Sizes struct {
	None            Size
	Small           Size
	Medium          Size
	Large           Size
	ExtraLarge      Size
	ExtraExtraLarge Size
}

type Size struct {
	Height unit.Dp
	Width  unit.Dp
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

func (b *Theme) GetComponent(id string) Base {
	cm, ok := b.Components[id]
	if !ok {
		return b.Base
	}

	// return a copy of the base theme with the components colors merged
	merged := b.Base
	merged.Surface = getOrDefault(cm.Surface, b.Base.Surface)
	merged.SurfaceHighlight = getOrDefault(cm.SurfaceHighlight, b.Base.SurfaceHighlight)
	merged.Text = getOrDefault(cm.Text, b.Base.Text)
	merged.TextSubtle = getOrDefault(cm.TextSubtle, b.Base.TextSubtle)
	merged.TextSubtlest = getOrDefault(cm.TextSubtlest, b.Base.TextSubtlest)
	merged.Border = getOrDefault(cm.Border, b.Base.Border)
	merged.Primary = getOrDefault(cm.Primary, b.Base.Primary)
	merged.Secondary = getOrDefault(cm.Secondary, b.Base.Secondary)
	merged.Info = getOrDefault(cm.Info, b.Base.Info)
	merged.Success = getOrDefault(cm.Success, b.Base.Success)
	merged.Notice = getOrDefault(cm.Notice, b.Base.Notice)
	merged.Warning = getOrDefault(cm.Warning, b.Base.Warning)
	merged.Danger = getOrDefault(cm.Danger, b.Base.Danger)
	return merged
}

func getOrDefault(c color.NRGBA, d color.NRGBA) color.NRGBA {
	if c == (color.NRGBA{}) {
		return d
	}
	return c
}
