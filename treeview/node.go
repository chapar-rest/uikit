package treeview

import (
	"strings"

	"gioui.org/gesture"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

const (
	mimeText = "application/json"
	NodeMIME = "chapar-uikit/treeview/node"
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

	paddingLeft unit.Dp

	draggable widget.Draggable
	// entered and dnsInited are for Drag and Drop op.
	entered   bool
	dndInited bool
	reader    *strings.Reader

	OnDropConfirmFunc func(source *Node, target *Node) bool
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
		paddingLeft: unit.Dp(8),
	}
}

func (n *Node) AddChild(child *Node) {
	child.Parent = n
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
