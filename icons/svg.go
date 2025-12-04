package icons

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"github.com/inkeliz/giosvg"
)

type SvgIcon struct {
	icon *giosvg.Icon
}

func (s *SvgIcon) Layout(gtx layout.Context, color color.NRGBA) layout.Dimensions {
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	return s.icon.Layout(gtx)
}

func loadSvgIcon(name string) *SvgIcon {
	iconFile, err := assetsDir.ReadFile(fmt.Sprintf("assets/%s.svg", name))
	if err != nil {
		panic(err)
	}

	vector, err := giosvg.NewVector(iconFile)
	if err != nil {
		panic(err)
	}

	return &SvgIcon{
		icon: giosvg.NewIcon(vector),
	}
}
