package treeview

import (
	"image"
	"image/color"
	"io"
	"log"
	"strings"

	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/io/transfer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/icons"
	"github.com/chapar-rest/uikit/theme"
)

func (n *Node) Update(gtx layout.Context) {
	if n.reader == nil || n.reader.Len() <= 0 {
		n.reader = strings.NewReader(n.ID)
	}

	filters := []event.Filter{
		transfer.TargetFilter{Target: n, Type: mimeText},
	}
	if len(n.Children) > 0 {
		filters = append(filters,
			// For DnD. This ensures only dir can be dragged and dropped to.
			transfer.TargetFilter{Target: n, Type: NodeMIME},
			// Detect if pointer is inside of the dir item, so we can highlight it when dropping items to it.
			pointer.Filter{Target: n, Kinds: pointer.Enter | pointer.Leave},
		)
	}

	for {
		ke, ok := gtx.Event(filters...)
		if !ok {
			break
		}
		switch event := ke.(type) {
		case pointer.Event:
			switch event.Kind {
			case pointer.Enter:
				n.entered = true
			case pointer.Leave:
				n.entered = false
			}
		case transfer.InitiateEvent:
			n.dndInited = true
		case transfer.CancelEvent:
			n.dndInited = false
			n.entered = false
		case transfer.DataEvent:
			// read the clipboard content:
			reader := event.Open()
			defer reader.Close()
			_, err := io.ReadAll(reader)
			if err != nil {
				log.Println("error reading clipboard content:", err)
				continue
			}
			defer gtx.Execute(op.InvalidateCmd{})
			switch event.Type {
			case NodeMIME:
				source, isFromNode := reader.(*Node)
				if !isFromNode {
					break
				}
				if source == n || source.Parent == n {
					break
				}
				if n.OnDropConfirmFunc != nil {
					if !n.OnDropConfirmFunc(source, n) {
						break
					}
				}
			}
		}
	}

	for {
		ev, ok := n.click.Update(gtx.Source)
		if !ok {
			break
		}
		switch ev.Kind {
		case gesture.KindClick:
			// TODO: make it a configurable option on treeview component.
			if len(n.Children) > 0 {
				n.discloser.ToggleVisibility(gtx.Now)
			} else {
				if n.OnClickFunc != nil {
					n.OnClickFunc(n)
				}
			}
		}
	}

	if n.draggable.Type == "" {
		n.draggable.Type = NodeMIME
	}

	if m, ok := n.draggable.Update(gtx); ok {
		n.draggable.Offer(gtx, m, n)
	}
}

func (n *Node) Read(p []byte) (int, error) {
	return n.reader.Read(p)
}

func (n *Node) Close() error {
	return nil
}

func (n *Node) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	n.Update(gtx)
	c := op.Record(gtx.Ops)
	dims := n.layout(gtx, th)
	call := c.Stop()
	defer pointer.PassOp{}.Push(gtx.Ops).Pop()

	call.Add(gtx.Ops)
	return dims
}

func (n *Node) droppable() bool {
	return n.entered && n.dndInited && !n.draggable.Dragging()
}

func (n *Node) layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if len(n.Children) == 0 {
		return n.rootLayout(gtx, false, th)
	} else {
		return Discloser(&n.discloser).Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				return n.rootLayout(gtx, true, th)
			},
			func(gtx layout.Context) layout.Dimensions {
				return n.detailLayout(gtx, th)
			},
		)
	}
}

func (n *Node) draggeBox(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	if !n.draggable.Dragging() {
		return layout.Dimensions{}
	}

	offset := n.draggable.Pos()
	if offset.Round().X == 0 && offset.Round().Y == 0 {
		return layout.Dimensions{}
	}

	macro := op.Record(gtx.Ops)
	dims := func(gtx layout.Context) layout.Dimensions {
		return widget.Border{
			Color:        th.Base.SurfaceHighlight,
			Width:        unit.Dp(1),
			CornerRadius: unit.Dp(8),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top:    unit.Dp(4),
				Bottom: unit.Dp(4),
				Left:   unit.Dp(8),
				Right:  unit.Dp(8),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return n.Widget(gtx, th)
			})
		})
	}(gtx)
	call := macro.Stop()

	defer clip.UniformRRect(image.Rectangle{Max: dims.Size}, gtx.Dp(unit.Dp(4))).Push(gtx.Ops).Pop()
	paint.ColorOp{Color: th.Base.SurfaceHighlight}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	defer paint.PushOpacity(gtx.Ops, 0.8).Pop()
	call.Add(gtx.Ops)

	return dims
}

func (n *Node) hasParent() bool {
	return n.Parent != nil
}

func (n *Node) rootLayout(gtx layout.Context, withControl bool, th *theme.Theme) layout.Dimensions {
	_, bg, _ := th.FgBgTxt(theme.KindPrimary, TreeComponent)
	hoverBgColor := colors.Hovered(bg)
	var bgColor color.NRGBA
	if n.click.Hovered() || n.droppable() {
		bgColor = hoverBgColor
	} else {
		bgColor = color.NRGBA{}
	}

	if n.droppable() && withControl && !n.discloser.Visible() {
		n.discloser.ToggleVisibility(gtx.Now)
	}

	paddingLeft := unit.Dp(8)
	if n.hasParent() {
		paddingLeft = unit.Dp(14) + n.Parent.paddingLeft
	}

	n.paddingLeft = paddingLeft

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
				if withControl {
					return n.discloser.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return n.controlLayout(gtx, th)
					})
				}
				return layout.Dimensions{Size: image.Point{X: gtx.Dp(16), Y: gtx.Dp(16)}}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return n.Widget(gtx, th)
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

	if bgColor != (color.NRGBA{}) {
		// only paint the hover color
		paint.Fill(gtx.Ops, bgColor)
	}

	return n.draggable.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			call.Add(gtx.Ops)
			return dims
		},
		func(gtx layout.Context) layout.Dimensions {
			return n.draggeBox(gtx, th)
		},
	)
}

func (n *Node) controlLayout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	_, _, txt := th.FgBgTxt(theme.KindPrimary, TreeComponent)
	if len(n.Children) == 0 {
		return layout.Dimensions{}
	}
	return layout.Inset{
		Right: unit.Dp(4),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Max.X = gtx.Dp(14)
		if !n.discloser.Visible() {
			return icons.ChevronRight.Layout(gtx, txt)
		}
		return icons.ChevronDown.Layout(gtx, txt)
	})
}

func (n *Node) detailLayout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	listStyle := material.List(th.Material(), n.childrenList)
	// hide the detail component scrollbar
	listStyle.ScrollbarStyle.Indicator.MinorWidth = 0
	listStyle.ScrollbarStyle.Track.MinorPadding = 0
	listStyle.ScrollbarStyle.Track.MajorPadding = 0
	return listStyle.Layout(gtx, len(n.Children), func(gtx layout.Context, i int) layout.Dimensions {
		return n.Children[i].Layout(gtx, th)
	})
}
