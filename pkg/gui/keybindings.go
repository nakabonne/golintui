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
		issues, err := g.runner.Run()
		if err != nil {
			g.resultsItem.ShowMessage(err.Error())
			return
		}
		if len(issues) == 0 {
			g.resultsItem.ShowMessage("no issues found")
		} else {
			g.resultsItem.SetLatestIssues(issues)
			g.resultsItem.ShowLatestIssues()
			g.switchPanel(g.resultsItem)
		}
	}
}
