package icons

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type Icon interface {
	Layout(gtx layout.Context, color color.NRGBA) layout.Dimensions
}

var (
	ChevronRight = loadSvgIcon("chevron-right")
	ChevronDown  = loadSvgIcon("chevron-down")

	CircleIcon *widget.Icon = func() *widget.Icon { icon, _ := widget.NewIcon(icons.ImageLens); return icon }()
	CloseIcon  *widget.Icon = func() *widget.Icon { icon, _ := widget.NewIcon(icons.NavigationClose); return icon }()
)
