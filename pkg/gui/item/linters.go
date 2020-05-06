package item

import (
	"sort"

	"github.com/k0kubun/pp"

	"github.com/nakabonne/golintui/pkg/golangcilint/config"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	DefaultLinterColor = tcell.ColorRed
	EnabledLinterColor = tcell.ColorLime
)

type Linters struct {
	*tview.TreeView
}

func NewLinters(linters []config.Linter) *Linters {
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

func (l *Linters) SetKeybinds(globalKeybind func(event *tcell.EventKey), selectAction func(*tview.TreeNode, *config.Linter), unselectAction func(*tview.TreeNode, *config.Linter) error) {
	l.SetSelectedFunc(func(node *tview.TreeNode) {
		ref, ok := node.GetReference().(config.Linter)
		if !ok {
			return
		}
		// TODO: Use ref.Enable() to check if enabled.
		//    For that, be sure to populate a pointer to linter to the reference.
		if node.GetColor() == EnabledLinterColor {
			if err := unselectAction(node, &ref); err != nil {
				pp.Println(err) // TODO: Replace popup
			}
		} else {
			selectAction(node, &ref)
		}
	})

	l.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'a':
			// TODO: Enable/Disable all linters
			return event
		}
		globalKeybind(event)
		return event
	})
}

func (l *Linters) addChildren(node *tview.TreeNode, linters []config.Linter) {
	sort.SliceStable(linters, func(i, j int) bool { return linters[i].Name() < linters[j].Name() })
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
