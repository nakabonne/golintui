package item

import (
	"github.com/rivo/tview"
)

type Info struct {
	*tview.TextView
}

func NewInfo(version string) *Info {
	b := &Info{
		TextView: tview.NewTextView(),
	}
	b.SetText(version).
		SetBorder(false).
		SetTitle("Info").SetTitleAlign(tview.AlignLeft)
	return b
}
