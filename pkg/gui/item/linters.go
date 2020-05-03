package item

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Linters struct {
	*tview.Table
}

func NewLinters() *Linters {
	l := &Linters{
		Table: tview.NewTable(),
	}

	l.SetBorders(true).
		SetSelectable(true, false).
		SetTitle("Linters").SetTitleAlign(tview.AlignLeft)

	return l
}

func (s *Linters) SetKeybinds(globalKeybind func(event *tcell.EventKey)) {
	s.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		globalKeybind(event)
		return event
	})
}
