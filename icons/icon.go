package icons

import (
	"image/color"

	"gioui.org/layout"
)

type Icon interface {
	Layout(gtx layout.Context, color color.NRGBA) layout.Dimensions
}

var (
	ChevronRight = loadSvgIcon("chevron-right")
	ChevronDown  = loadSvgIcon("chevron-down")
)
