package item

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/nakabonne/golintui/pkg/golangcilint"
)

const (
	DefaultLinterColor = tcell.ColorAqua
	EnabledLinterColor = tcell.ColorYellow
)

type Linters struct {
	*tview.TreeView
}

func NewLinters(linters []golangcilint.Linter) *Linters {
	l := &Linters{
		TreeView: tview.NewTreeView(),
	}
	root := tview.NewTreeNode("").
		SetColor(tcell.ColorWhite)

	l.SetRoot(root).SetCurrentNode(root).
		SetBorder(true).SetTitle("Linters").SetTitleAlign(tview.AlignLeft)

	l.addChildren(root, linters)
	return l
}

func (l *Linters) SetKeybinds(globalKeybind func(event *tcell.EventKey), selectAction, unselectAction func(*tview.TreeNode, string)) {
	l.SetSelectedFunc(func(node *tview.TreeNode) {
		ref, ok := node.GetReference().(golangcilint.Linter)
		if !ok {
			return
		}
		if ref.Enabled() {
			unselectAction(node, ref.Name())
		} else {
			selectAction(node, ref.Name())
		}
	})

	l.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		globalKeybind(event)
		return event
	})
}

func (l *Linters) addChildren(node *tview.TreeNode, linters []golangcilint.Linter) {
	// TODO: Sort linters by alphabet
	for _, linter := range linters {
		child := tview.NewTreeNode(linter.Name()).
			SetReference(linter).
			SetSelectable(true)
		if linter.Enabled() {
			child.SetColor(EnabledLinterColor)
		} else {
			child.SetColor(DefaultLinterColor)
		}
		node.AddChild(child)
	}

}
