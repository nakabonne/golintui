package item

import "github.com/rivo/tview"

type Info struct {
	*tview.TextView
}

func NewInfo() *Info {
	b := &Info{
		TextView: tview.NewTextView(),
	}
	// TODO: Populate real version of golangci-lint
	b.SetText("golangci-lint version: 1.25.0").
		SetBorder(false).
		SetTitle("Info").SetTitleAlign(tview.AlignLeft)
	return b
}
