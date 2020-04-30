package item

import (
	"fmt"
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
		SetColor(tcell.ColorRed)
	s.SetRoot(root).
		SetCurrentNode(root).
		SetBorder(true).
		SetTitle("Source Files").
		SetTitleAlign(tview.AlignLeft)

	s.addChildren(root, rootDir)
	return s
}

func (s *SourceFiles) SetKeybinds(globalKeybind func(event *tcell.EventKey)) {
	// TODO: Be sure to buffer the selected directories instead of switching toggle.
	s.SetSelectedFunc(s.SwitchToggle)

	s.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		node := s.GetCurrentNode()
		switch event.Rune() {
		case 'r':
			// TODO: Run linters against the directories marked as selected.
			if ref := node.GetReference(); ref != nil {
				fmt.Println(ref)
			}
		case 'l':
			// TODO: Expand toggle
		case 'h':
			// TODO: Collapse toggle
		}
		globalKeybind(event)
		return event
	})
}

// AddChildren adds child nodes to the given node which represents a directory.
func (s *SourceFiles) addChildren(target *tview.TreeNode, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		node := tview.NewTreeNode(file.Name()).
			SetReference(filepath.Join(path, file.Name())).
			SetSelectable(file.IsDir())
		if file.IsDir() {
			node.SetColor(tcell.ColorGreen)
		}
		target.AddChild(node)
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
		s.addChildren(node, path)
	} else {
		// Collapse if visible, expand if collapsed.
		node.SetExpanded(!node.IsExpanded())
	}
}
