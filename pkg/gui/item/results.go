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
		SetColor(tcell.ColorWhite)

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
		// Add a reporting linter to root as children.
		child := tview.NewTreeNode("reported by " + linter).
			SetReference(linter).
			SetSelectable(true).
			SetColor(tcell.ColorAqua)
		node.AddChild(child)

		// Add issues to reporting linters as children.
		issues := linterIssues[linter]
		for _, i := range issues {
			grandchild := tview.NewTreeNode(i.Message()).
				SetReference(i).
				SetSelectable(true).
				SetColor(tcell.ColorWhite)
			child.AddChild(grandchild)
		}
	}
}
