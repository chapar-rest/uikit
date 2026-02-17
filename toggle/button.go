package toggle

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
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/icons"
	"github.com/chapar-rest/uikit/theme"
)

const ToggleButtonComponent = "button"

// State defines one state of the toggle: icon only, text only, or both.
// Leave Icon nil for text-only; leave Label empty for icon-only.
// Color and Background are optional; zero value uses theme defaults when using ThemeStyle.
type State struct {
	Tag        string
	Icon       icons.Icon
	Label      string
	Color      color.NRGBA
	Background color.NRGBA
}

type ToggleButton struct {
	Kind         theme.Kind
	Button       *widget.Clickable
	CornerRadius unit.Dp

	states       []*State
	currentIndex int

	// Shared layout (can be overridden per-state via State.Color/Background)
	IconSize   unit.Sp
	IconInset  layout.Inset
	Background color.NRGBA
	Color      color.NRGBA
	Font       font.Font
	TextSize   unit.Sp
	Inset      layout.Inset
}

type ToggleButtonLayoutStyle struct {
	Background   color.NRGBA
	CornerRadius unit.Dp
	Button       *widget.Clickable
}

// NewToggleButton creates a toggle button with two or more states.
// States are cycled on each click. Caller must not pass empty states.
func NewToggleButton(th *theme.Theme, kind theme.Kind, states []*State) *ToggleButton {
	if len(states) == 0 {
		panic("toggle: at least one state required")
	}
	return &ToggleButton{
		Kind:         kind,
		states:       states,
		currentIndex: 0,
		Button:       &widget.Clickable{},
		CornerRadius: theme.RadiusSmall,

		IconSize:  unit.Sp(14),
		IconInset: theme.InsetSmall,
		Font:      font.Font{Typeface: th.Material().Face},
		TextSize:  unit.Sp(14),
		Inset:     theme.InsetNone,
	}
}

// State returns the current state (for immediate-mode reporting after Layout).
func (b *ToggleButton) State() *State {
	if len(b.states) == 0 {
		return nil
	}
	return b.states[b.currentIndex]
}

// StateIndex returns the current state index (0-based).
func (b *ToggleButton) StateIndex() int {
	return b.currentIndex
}

// StateTag returns the current state's Tag.
func (b *ToggleButton) StateTag() string {
	s := b.State()
	if s == nil {
		return ""
	}
	return s.Tag
}

// SetState sets the current state by index. No-op if index is out of range.
func (b *ToggleButton) SetState(index int) {
	if index < 0 || index >= len(b.states) {
		return
	}
	b.currentIndex = index
}

func (b *ToggleButton) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	_, bg, txt := th.FgBgTxt(b.Kind, ToggleButtonComponent)
	// Immediate mode: on click, cycle to next state
	if b.Button.Clicked(gtx) {
		b.currentIndex = (b.currentIndex + 1) % len(b.states)
	}

	cur := b.State()
	if cur == nil {
		return layout.Dimensions{}
	}

	return ToggleButtonLayoutStyle{
		Background:   bg,
		CornerRadius: b.CornerRadius,
		Button:       b.Button,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		items := []layout.FlexChild{}

		if cur.Icon != nil {
			items = append(items, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return b.IconInset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					sz := gtx.Sp(b.IconSize)
					gtx.Constraints.Min.X = sz
					gtx.Constraints.Max.X = sz
					return cur.Icon.Layout(gtx, txt)
				})
			}))
		}

		if cur.Label != "" {
			items = append(items, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lb := material.Label(th.Material(), b.TextSize, cur.Label)
				lb.Font = b.Font
				lb.Color = txt
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

func (t ToggleButtonLayoutStyle) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	min := gtx.Constraints.Min
	return t.Button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				rr := gtx.Dp(t.CornerRadius)
				defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, rr).Push(gtx.Ops).Pop()
				background := t.Background
				switch {
				case !gtx.Enabled():
					background = colors.Disabled(t.Background)
				case t.Button.Hovered() || gtx.Focused(t.Button):
					background = colors.Hovered(t.Background)
				}
				paint.Fill(gtx.Ops, background)
				for _, c := range t.Button.History() {
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
	const (
		expandDuration = float32(0.5)
		fadeDuration   = float32(0.9)
	)

	now := gtx.Now
	t := float32(now.Sub(c.Start).Seconds())
	end := c.End
	if end.IsZero() {
		end = now
	}
	endt := float32(end.Sub(c.Start).Seconds())

	var alphat float32
	{
		var haste float32
		if c.Cancelled {
			if h := 0.5 - endt/fadeDuration; h > 0 {
				haste = h
			}
		}
		half1 := t/fadeDuration + haste
		if half1 > 0.5 {
			half1 = 0.5
		}
		half2 := float32(now.Sub(end).Seconds())
		half2 /= fadeDuration
		half2 += haste
		if half2 > 0.5 {
			return
		}
		alphat = half1 + half2
	}

	sizet := t
	if c.Cancelled {
		sizet = endt
	}
	sizet /= expandDuration

	if !c.End.IsZero() || sizet <= 1.0 {
		gtx.Execute(op.InvalidateCmd{})
	}

	if sizet > 1.0 {
		sizet = 1.0
	}

	if alphat > .5 {
		alphat = 1.0 - alphat
	}
	t2 := alphat * 2
	alphaBezier := t2 * t2 * (3.0 - 2.0*t2)
	sizeBezier := sizet * sizet * (3.0 - 2.0*sizet)
	size := gtx.Constraints.Min.X
	if h := gtx.Constraints.Min.Y; h > size {
		size = h
	}
	size = int(float32(size) * 2 * float32(math.Sqrt(2)) * sizeBezier)
	alpha := 0.7 * alphaBezier
	const col = 0.8
	ba, bc := byte(alpha*0xff), byte(col*0xff)
	rgba := colors.MulAlpha(color.NRGBA{A: 0xff, R: bc, G: bc, B: bc}, ba)
	ink := paint.ColorOp{Color: rgba}
	ink.Add(gtx.Ops)
	rr := size / 2
	defer op.Offset(c.Position.Add(image.Point{X: -rr, Y: -rr})).Push(gtx.Ops).Pop()
	defer clip.UniformRRect(image.Rectangle{Max: image.Pt(size, size)}, rr).Push(gtx.Ops).Pop()
	paint.PaintOp{}.Add(gtx.Ops)
}
