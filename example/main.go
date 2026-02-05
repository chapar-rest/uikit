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
	"github.com/chapar-rest/uikit/tabs"
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
		theme:     th,
		openFiles: make(map[string]fileView),
		openTabs:  make(map[string]*tabs.Tab),
	}
	state.tree = state.buildFileTree(th)
	state.tabitems = tabs.NewTabs()

	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			state.appLayout(gtx)
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

	tabitems *tabs.Tabs

	openFiles map[string]fileView
	openTabs  map[string]*tabs.Tab // path -> tab, so we can select an already-open tab
}

type fileView struct {
	Title  string
	Path   string
	Layout func(gtx layout.Context, th *theme.Theme) layout.Dimensions
}

func (s *appState) onFileNodeClick(node *treeview.Node) {
	// node id is the full path of the file
	path := node.ID
	if _, ok := s.openFiles[path]; ok {
		// File already has a tab; just select it
		if tab := s.openTabs[path]; tab != nil {
			s.tabitems.SelectTab(tab)
		}
		return
	}

	s.openFiles[path] = s.buildFileView(s.theme, path)

	t := tabs.NewTab(func(gtx layout.Context, th *theme.Theme) layout.Dimensions {
		lb := material.Label(th.Material(), unit.Sp(14), s.openFiles[path].Title)
		return lb.Layout(gtx)
	})

	// When the tab is closed, remove the path from openFiles and openTabs so clicking the tree node again can open a new tab.
	t.OnCloseFunc = func(tab *tabs.Tab) bool {
		delete(s.openFiles, path)
		delete(s.openTabs, path)
		return true
	}

	t.State = tabs.TabStateClean
	s.tabitems.AddTab(t)
	s.openTabs[path] = t
}

func (s *appState) appLayout(gtx layout.Context) {
	// paint the background of the window with the theme's surface color
	paint.Fill(gtx.Ops, s.theme.Base.Surface)

	s.split.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return s.tree.Layout(gtx, s.theme)
		},
		func(gtx layout.Context) layout.Dimensions {
			return s.tabitems.Layout(gtx, s.theme)
		},
	)
}

func (s *appState) buildFileTree(th *theme.Theme) *treeview.Tree {
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

		node := s.buildFileNode(th, entry, ".")
		if node != nil {
			tree.Insert(node)
		}
	}

	return tree
}

func (s *appState) buildFileNode(th *theme.Theme, entry os.DirEntry, parentPath string) *treeview.Node {
	name := entry.Name()
	fullPath := filepath.Join(parentPath, name)

	node := treeview.NewNode(fullPath, func(gtx layout.Context) layout.Dimensions {
		return material.Label(th.Material(), unit.Sp(14), name).Layout(gtx)
	})

	node.OnClickFunc = func(node *treeview.Node) {
		s.onFileNodeClick(node)
	}

	// If it's a directory, recursively add its contents as children
	if entry.IsDir() {
		dirEntries, err := os.ReadDir(fullPath)
		if err == nil {
			for _, childEntry := range dirEntries {
				childNode := s.buildFileNode(th, childEntry, fullPath)
				if childNode != nil {
					node.AddChild(childNode)
				}
			}
		}
	}

	return node
}

func (s *appState) buildFileView(th *theme.Theme, path string) fileView {
	return fileView{
		Title: path,
		Path:  path,
		Layout: func(gtx layout.Context, th *theme.Theme) layout.Dimensions {
			return material.Label(th.Material(), unit.Sp(14), path).Layout(gtx)
		},
	}
}
