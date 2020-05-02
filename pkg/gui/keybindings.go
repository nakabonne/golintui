package gui

import (
	"github.com/gdamore/tcell"
	"github.com/k0kubun/pp"
)

func (g *Gui) setKeybind() {
	g.sourceFilesItem.SetKeybinds(g.grobalKeybind, g.registerPath)
}

func (g *Gui) grobalKeybind(event *tcell.EventKey) {
	switch event.Rune() {
	case 'q':
		g.application.Stop()
	case 'r':
		// TODO: Run golangci-lint against the directories marked as selected.
		issues, err := g.runner.Run()
		if err != nil {
			g.logger.Error(err.Error())
			return
		}
		pp.Println(issues)
	}
}
