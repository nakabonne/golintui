package item

import "github.com/rivo/tview"

type Linters struct {
	*tview.Checkbox
}

func NewLinters() *Linters {
	b := &Linters{
		Checkbox: tview.NewCheckbox(),
	}
	b.SetLabel("golint").
		SetBorder(true).SetTitle("Linters").SetTitleAlign(tview.AlignLeft)
	return b
}
