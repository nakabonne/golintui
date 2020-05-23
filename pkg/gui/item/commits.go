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
	// Just one node can be selected.
	selectedNode *tview.TreeNode
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

func (c *Commits) SetKeybinds(globalKeybind func(event *tcell.EventKey), registerAction func(*tview.TreeNode, string), unregisterAction func(node *tview.TreeNode)) {
	c.SetSelectedFunc(func(node *tview.TreeNode) {
		ref, ok := node.GetReference().(*git.Commit)
		if !ok {
			return
		}
		if node.GetColor() == SelectedCommitColor {
			unregisterAction(node)
			node.SetColor(DefaultCommitColor)
		} else {
			registerAction(node, ref.SHA)
			node.SetColor(SelectedCommitColor)
			if c.selectedNode != nil {
				c.selectedNode.SetColor(DefaultCommitColor)
			}
			c.selectedNode = node
		}
	})

	c.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		globalKeybind(event)
		return event
	})
}

func (c *Commits) addChildren(node *tview.TreeNode, commits []*git.Commit) {
	for _, commit := range commits {
		child := tview.NewTreeNode(fmt.Sprintf("%s %s", commit.ShortSha(), commit.Message)).
			SetReference(commit).
			SetSelectable(true).
			SetColor(DefaultCommitColor)

		node.AddChild(child)
	}
}
