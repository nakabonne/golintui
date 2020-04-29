package box

import "github.com/rivo/tview"

type TargetsBox struct {
	*tview.TreeView
}

func NewTargetsBox() *TargetsBox {
	b := &TargetsBox{
		TreeView: tview.NewTreeView(),
	}
	b.SetBorder(true).SetTitle("Targets").SetTitleAlign(tview.AlignLeft)
	return b
}
