package gui

import (
	"github.com/gdamore/tcell"
)

func (g *Gui) setKeybind() {
	g.lintersItem.SetKeybinds(g.grobalKeybind, g.enableLinter, g.disableLinter)
	g.sourceFilesItem.SetKeybinds(g.grobalKeybind, g.registerPath, g.unregisterPath)
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
	case 'q':
		g.application.Stop()
	case 'r':
		g.message("running linters...", "")

		go func() {
			issues, err := g.runner.Run()
			g.switchPage(modalPageName, mainPageName)
			if err != nil {
				g.resultsItem.ShowMessage(err.Error())
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
