package gui

import (
	"github.com/rivo/tview"

	"github.com/nakabonne/golintui/pkg/gui/box"
)

type Gui struct {
	application *tview.Application
}

func New() *Gui {
	return &Gui{
		application: tview.NewApplication(),
	}
}

func (g *Gui) Run() error {
	g.initPrimitive()
	if err := g.application.Run(); err != nil {
		return err
	}
	return nil
}

func (g *Gui) initPrimitive() {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTitleAlign(tview.AlignLeft).
			SetTitle(text)
	}
	targets := newPrimitive("Targets")
	results := newPrimitive("Results")

	grid := tview.NewGrid().
		SetRows(2, 0, 0).
		SetColumns(0, 0).
		SetBorders(true).
		AddItem(newPrimitive("Info"), 0, 0, 1, 1, 0, 0, true)

	// Layout for screens wider than 100 cells.
	grid.AddItem(box.NewLintersBox(), 1, 0, 1, 1, 0, 100, true).
		AddItem(targets, 2, 0, 1, 1, 0, 100, true).
		AddItem(results, 0, 1, 3, 1, 0, 100, true)

	g.application.SetRoot(grid, true)
}
