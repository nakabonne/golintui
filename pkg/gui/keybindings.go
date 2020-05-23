package gui

import (
	"fmt"

	"github.com/gdamore/tcell"
)

func (g *Gui) setKeybind() {
	g.lintersItem.SetKeybinds(g.grobalKeybind, g.enableLinter, g.disableLinter)
	g.sourceFilesItem.SetKeybinds(g.grobalKeybind, g.registerPath, g.unregisterPath)
	g.commitsItem.SetKeybinds(g.grobalKeybind, g.registerRevision, g.unregisterRevision)
	g.resultsItem.SetKeybinds(g.grobalKeybind, g.openFile)
}

func (g *Gui) grobalKeybind(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyTab:
		g.nextPanel()
	case tcell.KeyBacktab:
		g.prevPanel()
	}

	switch event.Rune() {
	case 'l':
		g.nextPanel()
	case 'h':
		g.prevPanel()
	case 'q':
		g.application.Stop()
	case 'r':
		paths := g.runner.ArgsString()
		close := g.showLoading("running linters...", fmt.Sprintf("given paths: %s", paths))

		go func() {
			issues, err := g.runner.Run()
			close()
			if err != nil {
				g.showWarn(err.Error(), "Press Enter to close")
				g.application.Draw()
				return
			}

			if len(issues) == 0 {
				g.resultsItem.ShowMessage("no issues found")
			} else {
				g.resultsItem.SetLatestIssues(issues)
				g.resultsItem.ShowLatestIssues()
				g.switchPanel(g.resultsItem)
			}
			g.application.Draw()
		}()

	}
}
