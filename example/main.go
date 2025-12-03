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
	}

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
}

func (s *appState) appLayout(gtx layout.Context, th *material.Theme) {
	s.split.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, unit.Sp(16), "Left").Layout(gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, unit.Sp(16), "Right").Layout(gtx)
		},
	)
}
