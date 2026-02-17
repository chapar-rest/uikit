package treeview

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/theme"
)

const (
	TreeComponent = "treeview"
)

type Tree struct {
	nodes        []*Node
	childrenList *widget.List

	Style TreeStyle
}

type TreeStyle struct {
	BackgroundColor color.NRGBA
	NodeStyle       NodeStyle
}

type NodeStyle struct {
	BackgroundColor      color.NRGBA
	HoverBackgroundColor color.NRGBA
}

func NewTree() *Tree {
	return &Tree{
		nodes: make([]*Node, 0),
		childrenList: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}
}

func (t *Tree) Insert(node *Node) {
	t.nodes = append(t.nodes, node)
}

func (t *Tree) Find(id string) *Node {
	for _, node := range t.nodes {
		if node.ID == id {
			return node
		}
	}
	return nil
}

func (t *Tree) Remove(id string) {
	for i, node := range t.nodes {
		if node.ID == id {
			t.nodes = append(t.nodes[:i], t.nodes[i+1:]...)
			return
		}
	}
}

func (t *Tree) Traverse(callback func(node *Node)) {
	for _, node := range t.nodes {
		callback(node)
		if len(node.Children) > 0 {
			t.Traverse(func(child *Node) {
				callback(child)
			})
		}
	}
}

func (t *Tree) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	//bgColor, _ := getBkColor(theme)
	// paint the background of the tree with the theme's background color
	//paint.Fill(gtx.Ops, bgColor)

	lst := material.List(theme.Material(), t.childrenList)
	lst.ScrollbarStyle = makeScrollbarStyle(theme, lst.ScrollbarStyle.Scrollbar)

	return lst.Layout(gtx, len(t.nodes), func(gtx layout.Context, i int) layout.Dimensions {
		return t.nodes[i].Layout(gtx, theme)
	})
}

func makeScrollbarStyle(th *theme.Theme, scrollbar *widget.Scrollbar) material.ScrollbarStyle {
	fg, bg, _ := th.FgBgTxt(theme.KindSecondary, TreeComponent)
	hoverFg := colors.Hovered(fg)

	return material.ScrollbarStyle{
		Scrollbar: scrollbar,
		Indicator: material.ScrollIndicatorStyle{
			Color:        fg,
			HoverColor:   hoverFg,
			CornerRadius: unit.Dp(0),
			MinorWidth:   unit.Dp(8),
		},
		Track: material.ScrollTrackStyle{
			Color:        bg,
			MajorPadding: unit.Dp(2),
			MinorPadding: unit.Dp(2),
		},
	}
}
