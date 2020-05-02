package item

import (
	"io/ioutil"
	"path/filepath"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type SourceFiles struct {
	*tview.TreeView
}

func NewSourceFiles(rootDir string) *SourceFiles {
	s := &SourceFiles{
		TreeView: tview.NewTreeView(),
	}

	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorWhite)
	s.SetRoot(root).
		SetCurrentNode(root).
		SetBorder(true).
		SetTitle("Source Files").
		SetTitleAlign(tview.AlignLeft)

	if err := s.addChildren(root, rootDir); err != nil {
		panic(err) // TODO: Emit log instead of panic
	}
	return s
}

func (s *SourceFiles) SetKeybinds(globalKeybind func(event *tcell.EventKey), selectedFunc func(node *tview.TreeNode)) {
	s.SetSelectedFunc(selectedFunc)

	s.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		node := s.GetCurrentNode()
		switch event.Rune() {
		case 'o':
			s.SwitchToggle(node)
		}
		globalKeybind(event)
		return event
	})
}

// AddChildren adds child nodes to the given node which represents a directory.
func (s *SourceFiles) addChildren(node *tview.TreeNode, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		child := tview.NewTreeNode(file.Name()).
			SetReference(filepath.Join(path, file.Name())).
			SetSelectable(file.IsDir())
		if file.IsDir() {
			child.SetColor(tcell.ColorAqua)
		}
		node.AddChild(child)
	}
	return nil
}

// SwitchToggle switches the current display state.
func (s *SourceFiles) SwitchToggle(node *tview.TreeNode) {
	reference := node.GetReference()
	if reference == nil {
		return // Selecting the root node does nothing.
	}
	children := node.GetChildren()
	if len(children) == 0 {
		// Load and show files in this directory.
		path := reference.(string)
		if err := s.addChildren(node, path); err != nil {
			panic(err) // TODO: Emit log instead of panic
		}
	} else {
		// Collapse if visible, expand if collapsed.
		node.SetExpanded(!node.IsExpanded())
	}
}
