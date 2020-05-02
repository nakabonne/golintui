package item

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/nakabonne/golintui/pkg/golangcilint"
)

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

// ShowLatest updates its own tree view and lists the latest execution results.
func (r *Results) ShowLatest(issues []golangcilint.Issue) {
	root := tview.NewTreeNode("Issues").
		SetColor(tcell.ColorYellow)

	r.SetRoot(root).
		SetCurrentNode(root)

	r.addChildren(issues)
}

func (r *Results) addChildren(issues []golangcilint.Issue) {

}
