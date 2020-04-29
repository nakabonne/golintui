package box

import "github.com/rivo/tview"

type InfoBox struct {
	*tview.TextView
}

func NewInfoBox() *InfoBox {
	b := &InfoBox{
		TextView: tview.NewTextView(),
	}
	b.SetText("golangci-lint version: 1.25.0").
		SetBorder(false).
		SetTitle("Info").SetTitleAlign(tview.AlignLeft)
	return b
}
