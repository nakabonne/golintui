package item

import "github.com/rivo/tview"

type Results struct {
	*tview.TreeView
}

func NewResults() *Results {
	b := &Results{
		TreeView: tview.NewTreeView(),
	}
	b.SetBorder(true).SetTitle("Results").SetTitleAlign(tview.AlignLeft)
	return b
}
