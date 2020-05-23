package gui

import (
	"github.com/gdamore/tcell"
	"github.com/sirupsen/logrus"

	"github.com/rivo/tview"

	"github.com/nakabonne/golintui/pkg/editor"
	"github.com/nakabonne/golintui/pkg/git"
	"github.com/nakabonne/golintui/pkg/golangcilint"
	"github.com/nakabonne/golintui/pkg/golangcilint/config"
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
	commitsItem     *item.Commits
	//infoItem        *item.Info
	naviItem *item.Navi

	runner *golangcilint.Runner
	editor *editor.Editor
	logger *logrus.Entry
}

func New(logger *logrus.Entry, runner *golangcilint.Runner, gitrunner *git.Runner, command *editor.Editor) *Gui {
	linters := runner.ListLinters()
	// TODO: Make limit changeable.
	commits, err := gitrunner.ListCommits(20)
	if err != nil {
		logger.Error(err.Error())
		commits = []*git.Commit{}
	}
	return &Gui{
		application:     tview.NewApplication(),
		lintersItem:     item.NewLinters(linters),
		sourceFilesItem: item.NewSourceFiles(logger, "."),
		resultsItem:     item.NewResults(logger),
		commitsItem:     item.NewCommits(commits),
		naviItem:        item.NewNavi(),
		runner:          runner,
		logger:          logger,
		editor:          command,
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
		SetRows(0, 10, 1).
		SetColumns(30, 30, 0, 0).
		SetBorders(true)

	// Layout for screens wider than 100 cells.
	grid.AddItem(g.lintersItem, 0, 0, 1, 1, 0, 100, true).
		AddItem(g.sourceFilesItem, 0, 1, 1, 1, 0, 100, false).
		AddItem(g.resultsItem, 0, 2, 2, 2, 0, 100, false).
		AddItem(g.commitsItem, 1, 0, 1, 2, 0, 100, false).
		AddItem(g.naviItem, 2, 0, 1, 4, 0, 0, false)

	// Layout for screens narrower than 100 cells.
	grid.AddItem(g.lintersItem, 0, 0, 1, 1, 0, 0, true).
		AddItem(g.sourceFilesItem, 0, 1, 1, 1, 0, 0, false).
		AddItem(g.resultsItem, 0, 2, 2, 2, 0, 0, false).
		AddItem(g.commitsItem, 1, 0, 1, 2, 0, 0, false).
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

// showWarn shows the given message as a modal.
// doneLabel param is used for the name of button to close.
func (g *Gui) showWarn(message, doneLabel string) {
	g.showModal(message, doneLabel, func() { g.switchPage(modalPageName, mainPageName) }, tcell.ColorBlack, tcell.ColorLime, tcell.ColorBlack)
}

// showLoading shows the given message as a modal.
// It doesn't provide any close buttons, instead it returns the function to close itself.
func (g *Gui) showLoading(message, doneLabel string) func() {
	g.showModal(message, doneLabel, func() {}, tcell.ColorBlack, tcell.ColorWhite, tcell.ColorBlack)
	close := func() {
		g.switchPage(modalPageName, mainPageName)
	}
	return close
}

func (g *Gui) showModal(message, doneLabel string, doneFunc func(), backGroundColor, buttonTextColor, buttonBackGroundColor tcell.Color) {
	modal := tview.NewModal().
		SetText(message).
		SetBackgroundColor(backGroundColor).
		SetButtonBackgroundColor(buttonTextColor).
		SetButtonTextColor(buttonBackGroundColor) // NOTE: The text color and background color are reversed, which is tview's issue.

	if doneLabel != "" {
		modal.AddButtons([]string{doneLabel}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == doneLabel {
					doneFunc()
				}
			})
	}

	colWidth, rowWidth := 1, 1
	grid := tview.NewGrid().
		SetColumns(0, colWidth, 0).
		SetRows(0, rowWidth, 0).
		AddItem(modal, 1, 1, 1, 1, 0, 0, true)
	g.pages.AddAndSwitchToPage(modalPageName, grid, true).ShowPage(mainPageName)
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
		g.showWarn(err.Error(), "Press Enter to close")
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
