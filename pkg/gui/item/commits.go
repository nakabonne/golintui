package item

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/nakabonne/golintui/pkg/git"
)

const (
	DefaultCommitColor  = tcell.ColorSilver
	SelectedCommitColor = tcell.ColorLime
)

type Commits struct {
	*tview.TreeView
}

func NewCommits(commits []*git.Commit) *Commits {
	c := &Commits{
		TreeView: tview.NewTreeView(),
	}
	root := tview.NewTreeNode("").
		SetColor(tcell.ColorWhite)

	c.SetRoot(root).SetCurrentNode(root).
		SetBorder(true).SetTitle("Commits").SetTitleAlign(tview.AlignLeft)

	c.addChildren(root, commits)
	return c
}

func (l *Commits) addChildren(node *tview.TreeNode, commits []*git.Commit) {
	for _, commit := range commits {
		child := tview.NewTreeNode(fmt.Sprintf("[red]%s[white] %s", commit.ShortSha(), commit.Message)).
			SetReference(commit).
			SetSelectable(true)

		node.AddChild(child)
	}
}
