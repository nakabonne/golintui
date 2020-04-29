package item

import "github.com/rivo/tview"

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
