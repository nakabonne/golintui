package app

import (
	"io"

	"github.com/sirupsen/logrus"

	"github.com/nakabonne/golintui/pkg/config"
	"github.com/nakabonne/golintui/pkg/gui"
)

type App struct {
	closers []io.Closer

	Config *config.Config
	Log    *logrus.Entry
	Gui    *gui.Gui
	// Tr            *i18n.Localizer
	// Updater       *updates.Updater // may only need this on the Gui
	// ClientContext string
}

func New(conf *config.Config) (*App, error) {
	return &App{
		closers: []io.Closer{},
		Config:  conf,
		Log:     newLogger(conf),
		Gui:     gui.New(),
	}, nil
}

func (a *App) Run() error {
	return a.Gui.Run()
}
