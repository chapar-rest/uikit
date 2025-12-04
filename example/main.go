package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/chapar-rest/uikit/colors"
	"github.com/chapar-rest/uikit/split"
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
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	state := &appState{
		split: &split.Split{
			Axis:  layout.Horizontal,
			Ratio: 0.3,
			HandleStyle: split.HandleStyle{
				Color:      colors.Gray,
				Width:      unit.Dp(3),
				HoverColor: colors.Blue,
			},
		},
		tree: treeview.NewTree(),
	}

	node1 := treeview.NewNode("1", func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(16), "1 node with children").Layout(gtx)
	})
	node1.Parent = nil

	child1 := treeview.NewNode("1.1", func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(16), "1.1 foo bar baz").Layout(gtx)
	})
	child1.Parent = node1

	node111 := treeview.NewNode("1.1.1", func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(16), "1.1.1").Layout(gtx)
	})
	node111.Parent = child1

	child1.AddChild(node111)

	node112 := treeview.NewNode("1.1.2", func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(16), "1.1.2").Layout(gtx)
	})
	node112.Parent = child1
	child1.AddChild(node112)

	node113 := treeview.NewNode("1.1.3", func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(16), "1.1.3").Layout(gtx)
	})
	node113.Parent = child1
	child1.AddChild(node113)

	child2 := treeview.NewNode("1.2", func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(16), "1.2").Layout(gtx)
	})

	child3 := treeview.NewNode("1.3", func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(16), "1.3").Layout(gtx)
	})

	node1.AddChild(child1)
	node1.AddChild(child2)
	node1.AddChild(child3)

	state.tree.Insert(node1)

	node2 := treeview.NewNode("2", func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(16), "2").Layout(gtx)
	})
	state.tree.Insert(node2)

	node3 := treeview.NewNode("3", func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(16), "3").Layout(gtx)
	})
	state.tree.Insert(node3)

	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			state.appLayout(gtx, th)
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

type appState struct {
	split *split.Split

	tree *treeview.Tree
}

func (s *appState) appLayout(gtx layout.Context, th *material.Theme) {
	s.split.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return s.tree.Layout(gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, unit.Sp(16), "Right").Layout(gtx)
		},
	)
}
