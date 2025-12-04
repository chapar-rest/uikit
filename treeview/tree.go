package treeview

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Tree struct {
	nodes        []*Node
	childrenList *widget.List

	NodeStyle NodeStyle
}

type NodeStyle struct {
	BorderOnHover bool
	BorderColor   color.NRGBA
	BorderWidth   unit.Dp
	BorderRadius  unit.Dp
}

func NewTree() *Tree {
	return &Tree{
		nodes: make([]*Node, 0),
		childrenList: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
		NodeStyle: NodeStyle{
			BorderOnHover: true,
			BorderColor:   color.NRGBA{0, 0, 0, 0},
			BorderWidth:   unit.Dp(1),
			BorderRadius:  unit.Dp(2),
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

func (t *Tree) Layout(gtx layout.Context) layout.Dimensions {
	theme := material.NewTheme()
	return material.List(theme, t.childrenList).Layout(gtx, len(t.nodes), func(gtx layout.Context, i int) layout.Dimensions {
		return t.nodes[i].Layout(gtx)
	})
}
