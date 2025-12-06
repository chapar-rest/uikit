package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

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
		tree:  buildFileTree(th),
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

func buildFileTree(th *theme.Theme) *treeview.Tree {
	tree := treeview.NewTree()

	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return tree
	}

	var ignoreList = []string{".git", ".idea", ".vscode", ".DS_Store", ".env"}

	for _, entry := range entries {
		if slices.Contains(ignoreList, entry.Name()) {
			continue
		}

		node := buildFileNode(th, entry, ".")
		if node != nil {
			tree.Insert(node)
		}
	}

	return tree
}

func buildFileNode(th *theme.Theme, entry os.DirEntry, parentPath string) *treeview.Node {
	name := entry.Name()
	fullPath := filepath.Join(parentPath, name)

	node := treeview.NewNode(fullPath, func(gtx layout.Context) layout.Dimensions {
		return material.Label(th.Material(), unit.Sp(16), name).Layout(gtx)
	})

	// If it's a directory, recursively add its contents as children
	if entry.IsDir() {
		dirEntries, err := os.ReadDir(fullPath)
		if err == nil {
			for _, childEntry := range dirEntries {
				childNode := buildFileNode(th, childEntry, fullPath)
				if childNode != nil {
					node.AddChild(childNode)
				}
			}
		}
	}

	return node
}
