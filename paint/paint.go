package paint

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func PaintRect(gtx layout.Context, size image.Point, fill color.NRGBA) {
	Rect{
		Color: fill,
		Size:  size,
	}.Layout(gtx)
}

type Rect struct {
	Color color.NRGBA
	Size  image.Point
	Radii int
}

func (r Rect) Layout(gtx layout.Context) layout.Dimensions {
	paint.FillShape(
		gtx.Ops,
		r.Color,
		clip.UniformRRect(
			image.Rectangle{
				Max: r.Size,
			},
			r.Radii,
		).Op(gtx.Ops))
	return layout.Dimensions{Size: r.Size}
}
