package gui

import (
	"github.com/sirupsen/logrus"

	"github.com/rivo/tview"

	"github.com/nakabonne/golintui/pkg/editor"
	"github.com/nakabonne/golintui/pkg/golangcilint"
	"github.com/nakabonne/golintui/pkg/gui/item"
)

// Gui wraps the tview application which handles rendering and events.
type Gui struct {
	application *tview.Application
	pages       *tview.Pages

	lintersItem     *item.Linters
	sourceFilesItem *item.SourceFiles
	resultsItem     *item.Results
	infoItem        *item.Info

	runner *golangcilint.Runner
	editor *editor.Editor
	logger *logrus.Entry
}

func New(logger *logrus.Entry, runner *golangcilint.Runner, command *editor.Editor) (*Gui, error) {
	linters, err := runner.ListLinters()
	if err != nil {
		return nil, err
	}
	return &Gui{
		application:     tview.NewApplication(),
		lintersItem:     item.NewLinters(linters),
		sourceFilesItem: item.NewSourceFiles("."),
		resultsItem:     item.NewResults(),
		infoItem:        item.NewInfo(runner.GetVersion()), // TODO: Run GetVersion() concurrency
		runner:          runner,
		logger:          logger,
		editor:          command,
	}, nil
}

func (g *Gui) Run() error {
	g.setKeybind()
	g.initGrid()
	if err := g.application.Run(); err != nil {
		g.application.Stop()
		return err
	}
	return nil
}

// initGrid sets a grid based layout as a root primitive for the application.
func (g *Gui) initGrid() {
	grid := tview.NewGrid().
		SetRows(1, 0).
		SetColumns(30, 40, 0, 0).
		SetBorders(true).
		AddItem(g.infoItem, 0, 0, 1, 2, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(g.lintersItem, 1, 0, 1, 1, 0, 100, false).
		AddItem(g.sourceFilesItem, 1, 1, 1, 1, 0, 100, true).
		AddItem(g.resultsItem, 0, 2, 2, 2, 0, 100, false)

	g.pages = tview.NewPages().
		AddAndSwitchToPage("main", grid, true)
	g.application.SetRoot(g.pages, true)
}

// TODO: Be more clear the relationship between each panel.
func (g *Gui) nextPanel() {
	switch g.application.GetFocus().(type) {
	case *item.Linters:
		g.switchPanel(g.sourceFilesItem)
	case *item.SourceFiles:
		g.switchPanel(g.resultsItem)
	case *item.Results:
		g.switchPanel(g.lintersItem)
	}
}

func (g *Gui) prevPanel() {
	switch g.application.GetFocus().(type) {
	case *item.Linters:
		g.switchPanel(g.resultsItem)
	case *item.SourceFiles:
		g.switchPanel(g.lintersItem)
	case *item.Results:
		g.switchPanel(g.sourceFilesItem)
	}
}

// registerPath adds path to golangci-lint runner as an arg.
func (g *Gui) registerPath(node *tview.TreeNode, path string) {
	g.runner.AddArgs(path)
	node.SetColor(item.SelectedDirColor)
}

func (g *Gui) unregisterPath(node *tview.TreeNode, path string) {
	g.runner.RemoveArgs(path)
	node.SetColor(item.DefaultDirColor)
}

// switchPanel switches to focus on the given Primitive.
func (g *Gui) switchPanel(p tview.Primitive) {
	g.application.SetFocus(p)
}

// openFile temporarily suspends this application and open file with the editor as a sub process.
func (g *Gui) openFile(filepath string, line, colmun int) error {
	g.application.Suspend(func() {
		if err := g.editor.OpenFileAtLineColumn(filepath, line, colmun); err != nil {
			g.logger.Error(err)
		}
	})
	return nil
}
