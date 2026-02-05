package divider

import (
	"image"
	"image/color"

	"github.com/chapar-rest/uikit/theme"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

type DividerStyle struct {
	Thickness unit.Dp
	Fill      color.NRGBA
	Inset     layout.Inset
	Axis      layout.Axis
}

func (d *DividerStyle) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if (d.Axis == layout.Horizontal && gtx.Constraints.Min.X == 0) ||
		(d.Axis == layout.Vertical && gtx.Constraints.Min.Y == 0) {
		return layout.Dimensions{}
	}

	if d.Fill == (color.NRGBA{}) {
		d.Fill = th.Base.SurfaceHighlight
	}

	return d.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		weight := gtx.Dp(d.Thickness)

		var maxDim = image.Point{}
		if d.Axis == layout.Horizontal {
			maxDim = image.Pt(gtx.Constraints.Min.X, weight)
		} else {
			maxDim = image.Pt(weight, gtx.Constraints.Min.Y)
		}

		line := image.Rectangle{Max: maxDim}
		paint.FillShape(gtx.Ops, d.Fill, clip.Rect(line).Op())
		return layout.Dimensions{Size: line.Max}
	})
}

func NewDivider(axis layout.Axis, thickness unit.Dp) *DividerStyle {
	return &DividerStyle{
		Thickness: thickness,
		Axis:      axis,
	}
}
