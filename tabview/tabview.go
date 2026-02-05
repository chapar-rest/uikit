// Tabview a is a copy of github.com/oligo/gioview tabview with some changes to make it more compatible with the uikit.
package tabview

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/chapar-rest/uikit/divider"
	"github.com/chapar-rest/uikit/theme"
)

var (
	horizontalInset = layout.Inset{Left: unit.Dp(2)}
	verticalInset   = layout.Inset{Top: unit.Dp(2)}
	horizontalFlex  = layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}
	verticalFlex    = layout.Flex{Axis: layout.Horizontal, Alignment: layout.Start}
)

type TabView struct {
	Axis        layout.Axis
	list        layout.List
	tabItems    []*TabItem
	currentView int
	headerSize  int
	bodySize    int
}

type TabItem struct {
	// Title of the tab.
	Title func(gtx layout.Context, th *theme.Theme) layout.Dimensions
	// Main part of the tab content.
	Widget func(gtx layout.Context, th *theme.Theme) layout.Dimensions
	// Title padding of the tab item.
	Inset     layout.Inset
	alignment layout.Direction
	click     gesture.Click
	hovering  bool
	selected  bool
}

func (item *TabItem) Update(gtx layout.Context) bool {
	for {
		event, ok := gtx.Event(
			pointer.Filter{Target: item, Kinds: pointer.Enter | pointer.Leave},
		)
		if !ok {
			break
		}

		switch event := event.(type) {
		case pointer.Event:
			switch event.Kind {
			case pointer.Enter:
				item.hovering = true
			case pointer.Leave:
				item.hovering = false
			case pointer.Cancel:
				item.hovering = false
			}
		}
	}

	var clicked bool
	for {
		e, ok := item.click.Update(gtx.Source)
		if !ok {
			break
		}
		if e.Kind == gesture.KindClick {
			clicked = true
			item.selected = true
		}
	}

	return clicked
}

func (item *TabItem) LayoutTab(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	item.Update(gtx)

	macro := op.Record(gtx.Ops)
	dims := item.layoutTab(gtx, th)
	call := macro.Stop()

	rect := clip.Rect(image.Rectangle{Max: dims.Size})
	defer rect.Push(gtx.Ops).Pop()

	item.click.Add(gtx.Ops)
	// register tag
	event.Op(gtx.Ops, item)
	call.Add(gtx.Ops)

	return dims
}

func (item *TabItem) layoutTab(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return item.layoutTabBackground(gtx, th)
		},
		func(gtx layout.Context) layout.Dimensions {
			return item.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return item.alignment.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return item.Title(gtx, th)
				})
			})
		},
	)
}

func (item *TabItem) layoutTabBackground(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	var fill color.NRGBA
	if item.hovering {
		fill = th.Base.SurfaceHighlight
	} else if item.selected {
		fill = th.Base.Primary
	}

	rr := gtx.Dp(unit.Dp(4))
	rect := clip.RRect{
		Rect: image.Rectangle{
			Max: image.Point{X: gtx.Constraints.Min.X, Y: gtx.Constraints.Min.Y},
		},
		NE: rr,
		SE: rr,
		NW: rr,
		SW: rr,
	}
	paint.FillShape(gtx.Ops, fill, rect.Op(gtx.Ops))
	return layout.Dimensions{Size: gtx.Constraints.Min}
}

func (item *TabItem) LayoutWidget(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	return item.Widget(gtx, th)
}

func (tv *TabView) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	tv.Update(gtx)

	if len(tv.tabItems) <= 0 {
		return layout.Dimensions{}
	}

	maxTabSize := tv.calculateWidth(gtx, th)
	var direction layout.Direction
	var flex layout.Flex
	var tabAlign layout.Direction
	if tv.Axis == layout.Horizontal {
		direction = layout.Center
		flex = horizontalFlex
		tabAlign = layout.Center
	} else {
		direction = layout.N
		flex = verticalFlex
		tabAlign = layout.W
	}

	return flex.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return direction.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				tv.list.Axis = tv.Axis
				tv.list.Alignment = layout.Start
				listDims := tv.list.Layout(gtx, len(tv.tabItems), func(gtx layout.Context, index int) layout.Dimensions {
					gtx.Constraints.Min = maxTabSize
					item := tv.tabItems[index]
					item.alignment = tabAlign

					if index == 0 {
						return item.LayoutTab(gtx, th)
					}

					if tv.Axis == layout.Horizontal {
						return horizontalInset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return item.LayoutTab(gtx, th)
						})
					} else {
						return verticalInset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return item.LayoutTab(gtx, th)
						})
					}

				})

				if tv.Axis == layout.Horizontal {
					tv.headerSize = listDims.Size.X
				} else {
					tv.headerSize = listDims.Size.Y
				}
				return listDims
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if tv.Axis == layout.Horizontal {
				return layout.Spacer{Height: unit.Dp(2)}.Layout(gtx)
			} else {
				return layout.Spacer{Width: unit.Dp(24)}.Layout(gtx)
			}
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if tv.Axis == layout.Horizontal {
				gtx.Constraints.Min.X = tv.headerSize
			} else {
				gtx.Constraints.Min.Y = max(tv.headerSize, tv.bodySize)
			}
			return divider.NewDivider(tv.Axis, unit.Dp(0.5)).Layout(gtx, th)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if tv.Axis == layout.Horizontal {
				return layout.Spacer{Height: unit.Dp(24)}.Layout(gtx)
			} else {
				return layout.Spacer{Width: unit.Dp(24)}.Layout(gtx)
			}
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			dims := tv.tabItems[tv.currentView].LayoutWidget(gtx, th)
			if tv.Axis == layout.Vertical {
				tv.bodySize = dims.Size.Y
				gtx.Execute(op.InvalidateCmd{})
			}

			return dims
		}),
	)
}

func (tv *TabView) Update(gtx layout.Context) {
	for idx, item := range tv.tabItems {
		if item.Update(gtx) {
			// unselect last item
			lastItem := tv.tabItems[tv.currentView]
			if lastItem != nil && idx != tv.currentView {
				lastItem.selected = false
			}

			tv.currentView = idx
		}

		if tv.currentView == idx && !item.selected {
			item.selected = true
		}
	}
}

func (tv *TabView) CurrentTab() int {
	return tv.currentView
}

func (tv *TabView) calculateWidth(gtx layout.Context, th *theme.Theme) image.Point {
	fakeOps := new(op.Ops)
	current := gtx.Ops
	gtx.Ops = fakeOps
	maxSize := image.Point{}

	gtx.Constraints.Min = image.Point{}
	for _, item := range tv.tabItems {
		dims := item.layoutTab(gtx, th)
		if dims.Size.X > maxSize.X {
			maxSize.X = dims.Size.X
		}
	}

	gtx.Ops = current
	return maxSize
}

func NewTabView(axis layout.Axis, item ...*TabItem) *TabView {
	return &TabView{
		Axis:     axis,
		tabItems: item,
	}
}

func NewTabItem(inset layout.Inset, title, wgt func(gtx layout.Context, th *theme.Theme) layout.Dimensions) *TabItem {
	return &TabItem{
		Title:  title,
		Widget: wgt,
		Inset:  inset,
	}
}

func SimpleTabItem(inset layout.Inset, title string, wgt func(gtx layout.Context, th *theme.Theme) layout.Dimensions) *TabItem {
	return &TabItem{
		Title: func(gtx layout.Context, th *theme.Theme) layout.Dimensions {
			label := material.Label(th.Material(), unit.Sp(14), title)
			label.Font.Weight = font.Medium
			return label.Layout(gtx)
		},
		Widget: wgt,
		Inset:  inset,
	}
}
