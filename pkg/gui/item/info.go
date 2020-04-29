package item

import "github.com/rivo/tview"

type Info struct {
	*tview.TextView
}

func NewInfo() *Info {
	b := &Info{
		TextView: tview.NewTextView(),
	}
	b.SetText("golangci-lint version: 1.25.0").
		SetBorder(false).
		SetTitle("Info").SetTitleAlign(tview.AlignLeft)
	return b
}
