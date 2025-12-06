package treeview

import (
	"image"
	"math"
	"time"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/widget"
	"gioui.org/x/component"
)

// DiscloserState holds state for a widget that can hide and reveal
// content.
type DiscloserState struct {
	component.VisibilityAnimation
	widget.Clickable
}

// Layout updates the state of the Discloser.
func (d *DiscloserState) Layout(gtx layout.Context) layout.Dimensions {
	if d.Duration == time.Duration(0) {
		d.Duration = time.Millisecond * 100
		d.State = component.Invisible
	}
	if d.Clicked(gtx) {
		d.ToggleVisibility(gtx.Now)
	}
	return layout.Dimensions{}
}

// DiscloserStyle defines the presentation of a discloser widget.
type DiscloserStyle struct {
	*DiscloserState
	// Alignment dictates how the control and summary are aligned relative
	// to one another.
	Alignment layout.Alignment
}

// Discloser configures a discloser from the provided theme and state.
func Discloser(state *DiscloserState) DiscloserStyle {
	return DiscloserStyle{
		DiscloserState: state,
		Alignment:      layout.Middle,
	}
}

// Layout the discloser with the provided toggle control, summary widget, and
// detail widget. The toggle widget will be wrapped in a clickable area
// automatically.
//
// The structure of the resulting discloser is:
//
//	root
//	-----------------
//	detail
//
// If d.ControlSide is set to Right, the control will appear after the summary
// instead of before it.
func (d DiscloserStyle) Layout(gtx layout.Context, root, detail layout.Widget) layout.Dimensions {
	d.DiscloserState.Layout(gtx)
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return root(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if !d.Visible() {
				return layout.Dimensions{}
			}
			if !d.Animating() {
				return detail(gtx)
			}
			progress := d.Revealed(gtx)
			macro := op.Record(gtx.Ops)
			dims := detail(gtx)
			call := macro.Stop()
			height := int(math.Round(float64(float32(dims.Size.Y) * progress)))
			dims.Size.Y = height
			defer clip.Rect(image.Rectangle{
				Max: dims.Size,
			}).Push(gtx.Ops).Pop()
			call.Add(gtx.Ops)
			return dims
		}),
	)
}
