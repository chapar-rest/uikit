package actionbar

import (
	"gioui.org/layout"
	"github.com/chapar-rest/uikit/theme"
)

type ActionBarItem interface {
	Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions
}

type ActionBar struct {
	Items []ActionBarItem

	Axis      layout.Axis
	Alignment layout.Alignment
	Spacing   layout.Spacing
}

func NewActionBar(axis layout.Axis, alignment layout.Alignment, spacing layout.Spacing) *ActionBar {
	return &ActionBar{
		Items:     make([]ActionBarItem, 0),
		Axis:      axis,
		Alignment: alignment,
		Spacing:   spacing,
	}
}

func (a *ActionBar) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	items := []layout.FlexChild{}
	for _, item := range a.Items {
		items = append(items, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return item.Layout(gtx, th)
		}))
	}

	return layout.Flex{
		Axis:      a.Axis,
		Alignment: a.Alignment,
		Spacing:   a.Spacing,
	}.Layout(gtx, items...)
}

func (a *ActionBar) AddItem(item ActionBarItem) {
	a.Items = append(a.Items, item)
}

func ActionBarItemFunc(fn func(gtx layout.Context, th *theme.Theme) layout.Dimensions) ActionBarItem {
	return actionBarItemFunc{Fn: fn}
}

type actionBarItemFunc struct {
	Fn func(gtx layout.Context, th *theme.Theme) layout.Dimensions
}

func (f actionBarItemFunc) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return f.Fn(gtx, th)
}
