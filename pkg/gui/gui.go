package gui

import (
	"github.com/rivo/tview"

	"github.com/nakabonne/golintui/pkg/gui/item"
)

type Gui struct {
	application *tview.Application
	pages       *tview.Pages
}

func New() *Gui {
	return &Gui{
		application: tview.NewApplication(),
	}
}

func (g *Gui) Run() error {
	g.initPrimitive()
	if err := g.application.Run(); err != nil {
		g.application.Stop()
		return err
	}
	return nil
}

func (g *Gui) initPrimitive() {
	grid := tview.NewGrid().
		SetRows(2, 0, 0).
		SetColumns(0, 0).
		SetBorders(true).
		AddItem(item.NewInfo(), 0, 0, 1, 1, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(item.NewLinters(), 1, 0, 1, 1, 0, 100, true).
		AddItem(item.NewSourceFiles(), 2, 0, 1, 1, 0, 100, false).
		AddItem(item.NewResults(), 0, 1, 3, 1, 0, 100, false)

	g.pages = tview.NewPages().
		AddAndSwitchToPage("main", grid, true)
	g.application.SetRoot(g.pages, true)
}
