package button

import (
	"image"
	"image/color"
	"math"

	"gioui.org/font"
	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/icons"
	"github.com/chapar-rest/uikit/position"
	"github.com/chapar-rest/uikit/theme"
)

const (
	ButtonComponent = "button"
)

type ButtonStyle struct {
	Text         string
	Icon         icons.Icon
	IconPosition position.Position
	// Color is the text color.
	Color        color.NRGBA
	Font         font.Font
	TextSize     unit.Sp
	Background   color.NRGBA
	CornerRadius unit.Dp
	Inset        layout.Inset
	Button       *widget.Clickable
	shaper       *text.Shaper

	IconSize  unit.Sp
	IconInset layout.Inset
}

type ButtonLayoutStyle struct {
	Background   color.NRGBA
	CornerRadius unit.Dp
	Button       *widget.Clickable
}

func TextButton(th *theme.Theme, button *widget.Clickable, txt string, kind theme.Kind) ButtonStyle {
	fg, bg := th.FgBg(kind, ButtonComponent)
	b := ButtonStyle{
		Text:         txt,
		Icon:         nil,
		Color:        fg,
		CornerRadius: theme.RadiusSmall,
		Background:   bg,
		TextSize:     th.Material().TextSize * 14.0 / 16.0,
		Inset:        theme.InsetMedium,
		Button:       button,
		shaper:       th.Material().Shaper,
	}
	b.Font.Typeface = th.Material().Face
	return b
}

func IconButton(th *theme.Theme, button *widget.Clickable, icon icons.Icon, kind theme.Kind) ButtonStyle {
	fg, bg := th.FgBg(kind, ButtonComponent)
	b := ButtonStyle{
		Text:         "",
		Icon:         icon,
		IconPosition: position.PositionLeft,
		Color:        fg,
		CornerRadius: theme.RadiusSmall,
		Background:   bg,
		TextSize:     th.Material().TextSize * 14.0 / 16.0,
		Inset:        theme.InsetSmall,
		Button:       button,
		shaper:       th.Material().Shaper,
		IconInset:    theme.InsetNone,
		IconSize:     unit.Sp(18),
	}
	b.Font.Typeface = th.Material().Face
	return b
}

func Button(th *theme.Theme, button *widget.Clickable, icon icons.Icon, iconPosition position.Position, txt string, kind theme.Kind) ButtonStyle {
	fg, bg := th.FgBg(kind, ButtonComponent)

	b := ButtonStyle{
		Text:         txt,
		Icon:         icon,
		IconPosition: iconPosition,
		Color:        fg,
		CornerRadius: theme.RadiusSmall,
		Background:   bg,
		TextSize:     th.Material().TextSize * 14.0 / 16.0,
		Inset:        theme.InsetSmall,
		Button:       button,
		shaper:       th.Material().Shaper,
		IconInset:    theme.InsetSmall,
		IconSize:     unit.Sp(18),
	}
	b.Font.Typeface = th.Material().Face
	return b
}

func (b ButtonStyle) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	return ButtonLayoutStyle{
		Background:   b.Background,
		CornerRadius: b.CornerRadius,
		Button:       b.Button,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		items := []layout.FlexChild{}

		if b.Icon != nil {
			items = append(items, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return b.IconInset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = gtx.Dp(unit.Dp(b.IconSize))
					gtx.Constraints.Max.X = gtx.Dp(unit.Dp(b.IconSize))
					return b.Icon.Layout(gtx, b.Color)
				})
			}))
		}

		if b.Text != "" {
			items = append(items, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lb := material.Label(theme.Material(), b.TextSize, b.Text)
				lb.Font = b.Font
				lb.Color = b.Color
				return lb.Layout(gtx)
			}))
		}

		return b.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: layout.SpaceAround}.Layout(gtx,
				items...,
			)
		})
	})
}

func (b ButtonLayoutStyle) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	min := gtx.Constraints.Min
	return b.Button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				rr := gtx.Dp(b.CornerRadius)
				defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, rr).Push(gtx.Ops).Pop()
				background := b.Background
				switch {
				case !gtx.Enabled():
					background = colors.Disabled(b.Background)
				case b.Button.Hovered() || gtx.Focused(b.Button):
					background = colors.Hovered(b.Background)
				}
				paint.Fill(gtx.Ops, background)
				for _, c := range b.Button.History() {
					drawInk(gtx, c)
				}
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min = min
				return layout.Center.Layout(gtx, w)
			},
		)
	})
}

func drawInk(gtx layout.Context, c widget.Press) {
	// duration is the number of seconds for the
	// completed animation: expand while fading in, then
	// out.
	const (
		expandDuration = float32(0.5)
		fadeDuration   = float32(0.9)
	)

	now := gtx.Now

	t := float32(now.Sub(c.Start).Seconds())

	end := c.End
	if end.IsZero() {
		// If the press hasn't ended, don't fade-out.
		end = now
	}

	endt := float32(end.Sub(c.Start).Seconds())

	// Compute the fade-in/out position in [0;1].
	var alphat float32
	{
		var haste float32
		if c.Cancelled {
			// If the press was cancelled before the inkwell
			// was fully faded in, fast forward the animation
			// to match the fade-out.
			if h := 0.5 - endt/fadeDuration; h > 0 {
				haste = h
			}
		}
		// Fade in.
		half1 := t/fadeDuration + haste
		if half1 > 0.5 {
			half1 = 0.5
		}

		// Fade out.
		half2 := float32(now.Sub(end).Seconds())
		half2 /= fadeDuration
		half2 += haste
		if half2 > 0.5 {
			// Too old.
			return
		}

		alphat = half1 + half2
	}

	// Compute the expand position in [0;1].
	sizet := t
	if c.Cancelled {
		// Freeze expansion of cancelled presses.
		sizet = endt
	}
	sizet /= expandDuration

	// Animate only ended presses, and presses that are fading in.
	if !c.End.IsZero() || sizet <= 1.0 {
		gtx.Execute(op.InvalidateCmd{})
	}

	if sizet > 1.0 {
		sizet = 1.0
	}

	if alphat > .5 {
		// Start fadeout after half the animation.
		alphat = 1.0 - alphat
	}
	// Twice the speed to attain fully faded in at 0.5.
	t2 := alphat * 2
	// Beziér ease-in curve.
	alphaBezier := t2 * t2 * (3.0 - 2.0*t2)
	sizeBezier := sizet * sizet * (3.0 - 2.0*sizet)
	size := gtx.Constraints.Min.X
	if h := gtx.Constraints.Min.Y; h > size {
		size = h
	}
	// Cover the entire constraints min rectangle and
	// apply curve values to size and color.
	size = int(float32(size) * 2 * float32(math.Sqrt(2)) * sizeBezier)
	alpha := 0.7 * alphaBezier
	const col = 0.8
	ba, bc := byte(alpha*0xff), byte(col*0xff)
	rgba := colors.MulAlpha(color.NRGBA{A: 0xff, R: bc, G: bc, B: bc}, ba)
	ink := paint.ColorOp{Color: rgba}
	ink.Add(gtx.Ops)
	rr := size / 2
	defer op.Offset(c.Position.Add(image.Point{
		X: -rr,
		Y: -rr,
	})).Push(gtx.Ops).Pop()
	defer clip.UniformRRect(image.Rectangle{Max: image.Pt(size, size)}, rr).Push(gtx.Ops).Pop()
	paint.PaintOp{}.Add(gtx.Ops)
}
