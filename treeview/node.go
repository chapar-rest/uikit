package treeview

import (
	"image"

	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/icons"
)

type Node struct {
	ID string

	Widget    layout.Widget
	Parent    *Node
	Collapsed bool

	Children     []*Node
	childrenList *widget.List

	discloser DiscloserState

	click gesture.Click
}

func NewNode(id string, w layout.Widget) *Node {
	return &Node{
		ID:        id,
		Widget:    w,
		Collapsed: false,
		childrenList: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}
}

func (n *Node) Layout(gtx layout.Context) layout.Dimensions {
	theme := material.NewTheme()

	var dims layout.Dimensions
	if len(n.Children) == 0 {
		dims = n.rootLayout(gtx, false)
	} else {
		dims = Discloser(theme, &n.discloser).Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				return n.rootLayout(gtx, true)
			},
			func(gtx layout.Context) layout.Dimensions {
				return n.detailLayout(gtx, theme)
			},
		)
	}

	return dims
}

func (n *Node) hasParent() bool {
	return n.Parent != nil
}

func (n *Node) rootLayout(gtx layout.Context, withDiscloser bool) layout.Dimensions {
	for {
		ev, ok := n.click.Update(gtx.Source)
		if !ok {
			break
		}
		switch ev.Kind {
		case gesture.KindClick:
			n.discloser.ToggleVisibility(gtx.Now)
		}
	}

	borderColor := colors.White
	if n.click.Hovered() {
		borderColor = colors.LightGray
	}

	border := widget.Border{
		Color:        borderColor,
		Width:        unit.Dp(0),
		CornerRadius: unit.Dp(1),
	}

	paddingLeft := unit.Dp(0)
	if n.hasParent() {
		paddingLeft = unit.Dp(16)
	}

	c := op.Record(gtx.Ops)
	dims := layout.Inset{
		Left:   paddingLeft,
		Top:    unit.Dp(8),
		Bottom: unit.Dp(8),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if withDiscloser {
					return n.discloser.Clickable.Layout(gtx, n.controlLayout)
				}
				return layout.Dimensions{Size: image.Point{X: gtx.Dp(16), Y: gtx.Dp(16)}}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return n.Widget(gtx)
			}),
		)
	})
	call := c.Stop()

	defer pointer.PassOp{}.Push(gtx.Ops).Pop()
	defer clip.Rect(image.Rectangle{
		Max: dims.Size,
	}).Push(gtx.Ops).Pop()
	event.Op(gtx.Ops, n)
	n.click.Add(gtx.Ops)
	paint.Fill(gtx.Ops, borderColor)

	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		call.Add(gtx.Ops)
		return dims
	})
}

func (n *Node) controlLayout(gtx layout.Context) layout.Dimensions {
	if len(n.Children) == 0 {
		return layout.Dimensions{}
	}
	return layout.Inset{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Max.X = gtx.Dp(16)
		if !n.discloser.Visible() {
			return icons.ChevronRight.Layout(gtx, colors.DarkGray)
		}
		return icons.ChevronDown.Layout(gtx, colors.DarkGray)
	})
}

func (n *Node) detailLayout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	listStyle := material.List(theme, n.childrenList)
	listStyle.ScrollbarStyle.Indicator.MinorWidth = 0
	listStyle.ScrollbarStyle.Track.MinorPadding = 0
	listStyle.ScrollbarStyle.Track.MajorPadding = 0
	return listStyle.Layout(gtx, len(n.Children), func(gtx layout.Context, i int) layout.Dimensions {
		return n.Children[i].Layout(gtx)
	})
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func (n *Node) RemoveChild(child *Node) {
	for i, c := range n.Children {
		if c == child {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			return
		}
	}
}

func (n *Node) Traverse(callback func(node *Node)) {
	callback(n)
	for _, child := range n.Children {
		child.Traverse(callback)
	}
}

func (n *Node) Find(id string) *Node {
	if n.ID == id {
		return n
	}
	for _, child := range n.Children {
		if node := child.Find(id); node != nil {
			return node
		}
	}
	return nil
}
