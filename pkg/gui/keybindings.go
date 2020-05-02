package gui

import (
	"github.com/gdamore/tcell"
)

func (g *Gui) setKeybind() {
	g.sourceFilesItem.SetKeybinds(g.grobalKeybind, g.registerPath, g.unregisterPath)
	g.resultsItem.SetKeybinds(g.grobalKeybind, g.openFile)
}

func (g *Gui) grobalKeybind(event *tcell.EventKey) {
	switch event.Rune() {
	case 'q':
		g.application.Stop()
	case 'r':
		issues, err := g.runner.Run()
		if err != nil {
			g.logger.Error(err.Error())
			return
		}
		g.resultsItem.ShowLatest(issues)
		g.switchPanel(g.resultsItem)
	}
}
