package box

import "github.com/rivo/tview"

type LintersBox struct {
	*tview.Checkbox
}

func NewLintersBox() *LintersBox {
	b := &LintersBox{
		Checkbox: tview.NewCheckbox(),
	}
	b.SetBorder(true).SetTitle("Linters").SetTitleAlign(tview.AlignLeft)
	return b
}
