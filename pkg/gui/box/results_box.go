package box

import "github.com/rivo/tview"

type ResultsBox struct {
	*tview.TreeView
}

func NewResultsBox() *ResultsBox {
	b := &ResultsBox{
		TreeView: tview.NewTreeView(),
	}
	b.SetBorder(true).SetTitle("Results").SetTitleAlign(tview.AlignLeft)
	return b
}
