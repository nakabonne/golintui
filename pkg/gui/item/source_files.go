package item

import "github.com/rivo/tview"

type SourceFiles struct {
	*tview.TreeView
}

func NewSourceFiles() *SourceFiles {
	b := &SourceFiles{
		TreeView: tview.NewTreeView(),
	}
	b.SetBorder(true).SetTitle("Source Files").SetTitleAlign(tview.AlignLeft)
	return b
}
