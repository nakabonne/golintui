package box

import "github.com/rivo/tview"

type LintersBox struct {
	*tview.Checkbox
}

func NewLintersBox() *LintersBox {
	l := &LintersBox{
		Checkbox: tview.NewCheckbox(),
	}
	l.SetBorder(true).SetTitle("Linters").SetTitleAlign(tview.AlignLeft)
	return l
}
