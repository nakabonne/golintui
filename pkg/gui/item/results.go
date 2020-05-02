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

	r.addChildren(root, issues)
}

func (r *Results) addChildren(node *tview.TreeNode, issues []golangcilint.Issue) {
	linterIssues := make(map[string][]golangcilint.Issue)
	for _, issue := range issues {
		l := issue.FromLinter()
		linterIssues[l] = append(linterIssues[l], issue)
	}

	for linter := range linterIssues {
		child := tview.NewTreeNode(linter).
			SetText("from " + linter).
			SetSelectable(true).
			SetColor(tcell.ColorGreen)
		node.AddChild(child)
	}
}
