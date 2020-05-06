package gui

import (
	"github.com/gdamore/tcell"
	"github.com/sirupsen/logrus"

	"github.com/nakabonne/golintui/pkg/golangcilint/config"

	"github.com/rivo/tview"

	"github.com/nakabonne/golintui/pkg/editor"
	"github.com/nakabonne/golintui/pkg/golangcilint"
	"github.com/nakabonne/golintui/pkg/gui/item"
)

const (
	mainPageName  = "main"
	modalPageName = "modal"
)

// Gui wraps the tview application which handles rendering and events.
type Gui struct {
	application *tview.Application
	pages       *tview.Pages

	lintersItem     *item.Linters
	sourceFilesItem *item.SourceFiles
	resultsItem     *item.Results
	infoItem        *item.Info
	naviItem        *item.Navi

	runner *golangcilint.Runner
	editor *editor.Editor
	logger *logrus.Entry
}

func New(logger *logrus.Entry, runner *golangcilint.Runner, command *editor.Editor) (*Gui, error) {
	linters := runner.ListLinters()
	return &Gui{
		application:     tview.NewApplication(),
		lintersItem:     item.NewLinters(linters),
		sourceFilesItem: item.NewSourceFiles("."),
		resultsItem:     item.NewResults(),
		infoItem:        item.NewInfo(runner.GetVersion()), // TODO: Run GetVersion() concurrency
		naviItem:        item.NewNavi(),
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
		SetRows(1, 0, 1).
		SetColumns(30, 40, 0, 0).
		SetBorders(true).
		AddItem(g.infoItem, 0, 0, 1, 2, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(g.lintersItem, 1, 0, 1, 1, 0, 100, true).
		AddItem(g.sourceFilesItem, 1, 1, 1, 1, 0, 100, false).
		AddItem(g.resultsItem, 0, 2, 2, 2, 0, 100, false).
		AddItem(g.naviItem, 2, 0, 1, 4, 0, 0, false)

	g.naviItem.Update(g.lintersItem)

	g.pages = tview.NewPages().
		AddAndSwitchToPage(mainPageName, grid, true)
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

// switchPanel switches to focus on the given Primitive.
func (g *Gui) switchPanel(p tview.Primitive) {
	g.application.SetFocus(p)
	g.naviItem.Update(p)
}

func (g *Gui) switchPage(prev, next string) {
	g.pages.RemovePage(prev).ShowPage(next)
}

func (g *Gui) modal(p tview.Primitive, width, height int) *tview.Grid {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}

// message shows the given message as a modal.
func (g *Gui) message(message, doneLabel string, doneFunc func()) {
	modal := tview.NewModal().
		SetText(message).
		SetBackgroundColor(tcell.ColorBlack).
		AddButtons([]string{doneLabel}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == doneLabel {
				doneFunc()
			}
			g.switchPage(modalPageName, mainPageName)
		})

	g.pages.AddAndSwitchToPage("modal", g.modal(modal, 150, 60), true).ShowPage(mainPageName)
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

func (g *Gui) enableLinter(node *tview.TreeNode, linter *config.Linter) {
	g.runner.EnableLinter(linter.Name())
	node.SetColor(item.EnabledLinterColor)
}

func (g *Gui) disableLinter(node *tview.TreeNode, linter *config.Linter) {
	if err := g.runner.DisableLinter(linter.Name()); err != nil {
		g.message(err.Error(), "Enter", func() {})
		return
	}
	node.SetColor(item.DefaultLinterColor)
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
