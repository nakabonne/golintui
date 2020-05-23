package item

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	DefaultCommitColor  = tcell.ColorSilver
	SelectedCommitColor = tcell.ColorLime
)

type Commits struct {
	*tview.TreeView
}

func NewCommits() *Commits {
	c := &Commits{
		TreeView: tview.NewTreeView(),
	}
	root := tview.NewTreeNode("").
		SetColor(tcell.ColorWhite)

	c.SetRoot(root).SetCurrentNode(root).
		SetBorder(true).SetTitle("Commits").SetTitleAlign(tview.AlignLeft)

	//c.addChildren(root, linters)
	return c
}
