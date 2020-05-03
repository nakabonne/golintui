package gui

import (
	"github.com/sirupsen/logrus"

	"github.com/rivo/tview"

	"github.com/nakabonne/golintui/pkg/golangcilint"
	"github.com/nakabonne/golintui/pkg/gui/item"
	"github.com/nakabonne/golintui/pkg/oscommand"
)

// Gui wraps the tview application which handles rendering and events.
type Gui struct {
	application *tview.Application
	pages       *tview.Pages

	lintersItem     *item.Linters
	sourceFilesItem *item.SourceFiles
	resultsItem     *item.Results
	infoItem        *item.Info

	runner    *golangcilint.Runner
	osCommand *oscommand.OSCommand
	logger    *logrus.Entry
}

func New(logger *logrus.Entry, runner *golangcilint.Runner, command *oscommand.OSCommand) *Gui {
	return &Gui{
		application:     tview.NewApplication(),
		lintersItem:     item.NewLinters(),
		sourceFilesItem: item.NewSourceFiles("."),
		resultsItem:     item.NewResults(),
		infoItem:        item.NewInfo(runner.GetVersion()), // TODO: Run GetVersion() concurrency
		runner:          runner,
		logger:          logger,
		osCommand:       command,
	}
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
		if err := g.osCommand.OpenFileAtLineColumn(filepath, line, colmun); err != nil {
			g.logger.Error(err)
		}
	})
	return nil
}
