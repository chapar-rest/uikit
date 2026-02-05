package tabs

import (
	"image"
	"image/color"
	"slices"

	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/chapar-rest/uikit/icons"
	"github.com/chapar-rest/uikit/theme"
)

type Tabs struct {
	tabs []*Tab

	list layout.List

	currentView int
}

type TabState string

const (
	TabStateDirty TabState = "dirty"
	TabStateClean TabState = "clean"
)

func NewTabs() *Tabs {
	return &Tabs{
		list: layout.List{
			Axis: layout.Horizontal,
		},
		tabs:        make([]*Tab, 0),
		currentView: 0,
	}
}

type TabWidget func(gtx layout.Context, th *theme.Theme) layout.Dimensions

type Tab struct {
	Widget TabWidget

	closed   bool
	selected bool
	hovering bool
	click    gesture.Click

	State TabState

	closeClickable widget.Clickable

	// OnCloseFunc is called when the close button is clicked.
	// It should return true if the tab should be closed, false otherwise.
	OnCloseFunc func(tab *Tab) bool

	// OnSelectFunc is called when the tab is selected.
	OnSelectFunc func(tab *Tab)
}

func NewTab(widget TabWidget) *Tab {
	return &Tab{
		Widget: widget,
	}
}

func (t *Tabs) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	t.Update(gtx)

	gtx.Constraints.Max.Y = gtx.Dp(46)
	return t.list.Layout(gtx, len(t.tabs), func(gtx layout.Context, i int) layout.Dimensions {
		return t.tabs[i].Layout(gtx, th)
	})
}

func (t *Tabs) AddTab(tab *Tab) {
	// Unselect the current tab if any
	if t.currentView >= 0 && t.currentView < len(t.tabs) {
		t.tabs[t.currentView].selected = false
	}
	t.tabs = append(t.tabs, tab)
	t.currentView = len(t.tabs) - 1
	tab.selected = true
}

func (t *Tab) Update(gtx layout.Context) bool {
	if t.closed {
		return false
	}

	// handle close button click
	for {
		e, ok := t.closeClickable.Update(gtx)
		if !ok {
			break
		}
		if e.NumClicks > 0 {
			t.handleOnClose()
			return false
		}
	}

	for {
		event, ok := gtx.Event(
			pointer.Filter{Target: t, Kinds: pointer.Enter | pointer.Leave | pointer.Cancel},
		)
		if !ok {
			break
		}
		switch event := event.(type) {
		case pointer.Event:
			switch event.Kind {
			case pointer.Enter:
				t.hovering = true
			case pointer.Leave:
				t.hovering = false
			case pointer.Cancel:
				t.hovering = false
			}
		}
	}

	var clicked bool
	for {
		e, ok := t.click.Update(gtx.Source)
		if !ok {
			break
		}
		if e.Kind == gesture.KindClick {
			t.selected = true
			clicked = true
			t.handleOnSelect()
		}
	}

	return clicked
}

func (t *Tab) handleOnSelect() {
	if t.OnSelectFunc != nil {
		t.OnSelectFunc(t)
		t.selected = true
	}
}

func (t *Tab) handleOnClose() {
	if t.OnCloseFunc != nil {
		t.closed = t.OnCloseFunc(t)
	}
}

func (t *Tab) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	t.Update(gtx)

	macro := op.Record(gtx.Ops)
	dims := t.layout(gtx, th)
	call := macro.Stop()

	defer pointer.PassOp{}.Push(gtx.Ops).Pop()
	defer clip.Rect(image.Rectangle{Max: dims.Size}).Push(gtx.Ops).Pop()

	t.click.Add(gtx.Ops)
	// register tag
	event.Op(gtx.Ops, t)
	call.Add(gtx.Ops)

	return dims
}

func (t *Tab) layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	macro := op.Record(gtx.Ops)
	dims := layout.Inset{
		Left:   unit.Dp(8),
		Right:  unit.Dp(8),
		Top:    unit.Dp(12),
		Bottom: unit.Dp(8),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:      layout.Horizontal,
			Alignment: layout.Start,
			Spacing:   layout.SpaceBetween,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return t.Widget(gtx, th)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Left: unit.Dp(8),
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return t.layoutCloseButton(gtx, th)
				})
			}),
		)
	})
	call := macro.Stop()

	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return t.layoutBackground(gtx, th)
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.Stack{
				Alignment: layout.S,
			}.Layout(gtx,
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					call.Add(gtx.Ops)
					return dims
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					if !t.selected {
						return layout.Dimensions{}
					}
					indicatorHeight := gtx.Dp(unit.Dp(2))
					indicatorRect := image.Rect(0, 0, dims.Size.X, indicatorHeight)
					paint.FillShape(gtx.Ops, th.Base.Primary, clip.Rect(indicatorRect).Op())
					return layout.Dimensions{
						Size: image.Point{X: dims.Size.X, Y: indicatorHeight},
					}
				}),
			)
		},
	)
}

func (t *Tab) layoutCloseButton(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	var icon icons.Icon = icons.CloseIcon
	var fill color.NRGBA = th.Base.Surface
	var padding int = 0
	if t.State == TabStateDirty {
		padding = 4
	}

	return t.closeClickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		hovering := t.closeClickable.Hovered()
		if hovering {
			fill = th.Base.Primary
		} else if t.selected {
			fill = th.Base.SurfaceHighlight
		} else {
			fill = th.Base.Surface
		}

		if hovering {
			padding = 0
			icon = icons.CloseIcon
		} else if t.State == TabStateDirty {
			padding = 4
			icon = icons.CircleIcon
		} else if t.State == TabStateClean {
			padding = 0
			icon = icons.CloseIcon
		}

		rr := gtx.Dp(unit.Dp(4))
		rect := clip.RRect{
			Rect: image.Rectangle{
				Max: image.Point{X: gtx.Dp(16), Y: gtx.Dp(16)},
			},
			NE: rr,
			SE: rr,
			NW: rr,
			SW: rr,
		}
		gtx.Constraints.Max.X = gtx.Dp(unit.Dp(16))
		gtx.Constraints.Max.Y = gtx.Dp(unit.Dp(16))
		paint.FillShape(gtx.Ops, fill, rect.Op(gtx.Ops))
		return layout.UniformInset(unit.Dp(padding)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return icon.Layout(gtx, th.Base.Text)
		})
	})
}

func (t *Tab) layoutBackground(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	var fill color.NRGBA
	if t.hovering {
		fill = th.Base.SurfaceHighlight
	} else {
		fill = th.Base.Surface
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

func (t *Tabs) Update(gtx layout.Context) {
	// First pass: process events and collect indices of closed tabs.
	// We must not delete while ranging; deleting shifts indices and skips items.
	var closedIndices []int
	for idx, item := range t.tabs {
		clicked := item.Update(gtx)

		if item.closed {
			closedIndices = append(closedIndices, idx)
			continue
		}

		if clicked {
			// unselect last item
			if t.currentView >= 0 && t.currentView < len(t.tabs) {
				lastItem := t.tabs[t.currentView]
				if lastItem != nil && idx != t.currentView {
					lastItem.selected = false
				}
			}
			t.currentView = idx
		}

		if t.currentView == idx && !item.selected {
			item.selected = true
		}
	}

	// Delete closed tabs in reverse order so indices remain valid.
	for i := len(closedIndices) - 1; i >= 0; i-- {
		idx := closedIndices[i]
		t.tabs = slices.Delete(t.tabs, idx, 1)
		// Adjust currentView: if we removed the selected tab or one before it, clamp or decrement.
		if idx < t.currentView {
			t.currentView--
		} else if idx == t.currentView {
			// Selected tab was closed; stay on the same index (next tab slides in) or clamp to end.
			if t.currentView >= len(t.tabs) {
				t.currentView = len(t.tabs) - 1
			}
		}
		if t.currentView < 0 {
			t.currentView = 0
		}
	}
}
