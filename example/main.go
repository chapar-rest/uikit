package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/chapar-rest/uikit/split"
	"github.com/chapar-rest/uikit/theme"
	"github.com/chapar-rest/uikit/theme/themes"
	"github.com/chapar-rest/uikit/treeview"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Size(unit.Dp(800), unit.Dp(700)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := themes.Dark()
	fonts := theme.BuiltinFonts()
	th.WithFonts(fonts)

	state := &appState{
		split: &split.Split{
			Axis:  layout.Horizontal,
			Ratio: 0.3,
			HandleStyle: split.HandleStyle{
				Color:      th.Base.Border,
				Width:      unit.Dp(3),
				HoverColor: th.Base.Secondary,
			},
		},
		tree:  buildMockTreeData(th),
		theme: th,
	}

	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			state.appLayout(gtx, th.Material())
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

type appState struct {
	split *split.Split

	tree *treeview.Tree

	theme *theme.Theme
}

func (s *appState) appLayout(gtx layout.Context, th *material.Theme) {
	// paint the background of the window with the theme's surface color
	paint.Fill(gtx.Ops, s.theme.Base.Surface)

	s.split.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return s.tree.Layout(gtx, s.theme)
		},
		func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, unit.Sp(16), "Right").Layout(gtx)
		},
	)
}

func buildMockTreeData(th *theme.Theme) *treeview.Tree {
	fmt.Println("building mock tree data")
	tree := treeview.NewTree()

	// Allow enough nodes for 10 root nodes plus their children
	maxNodes := 50
	nodeCount := 0

	// Always create exactly 10 root nodes
	numTopLevel := 10

	for i := 0; i < numTopLevel; i++ {
		node := buildNode(th, fmt.Sprintf("node%d", i), 0, &nodeCount, maxNodes)
		if node != nil {
			tree.Insert(node)
		}
	}

	return tree
}

func buildNode(th *theme.Theme, name string, depth int, nodeCount *int, maxNodes int) *treeview.Node {
	// Check if we've reached max nodes or max depth
	if *nodeCount >= maxNodes || depth >= 5 {
		return nil
	}

	*nodeCount++
	node := treeview.NewNode(name, func(gtx layout.Context) layout.Dimensions {
		return material.Label(th.Material(), unit.Sp(16), name).Layout(gtx)
	})

	// Randomly decide if this node should have children (70% chance)
	shouldHaveChildren := rand.Intn(100) < 70 && depth < 5 && *nodeCount < maxNodes

	if shouldHaveChildren {
		// Randomly decide how many children (1-3)
		numChildren := rand.Intn(3) + 1

		for i := 0; i < numChildren && *nodeCount < maxNodes; i++ {
			childName := fmt.Sprintf("%s-child%d", name, i)
			child := buildNode(th, childName, depth+1, nodeCount, maxNodes)
			if child != nil {
				node.AddChild(child)
			}
		}
	}

	return node
}
