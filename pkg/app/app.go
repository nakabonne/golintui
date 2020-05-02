package app

import (
	"io"

	"github.com/sirupsen/logrus"

	"github.com/nakabonne/golintui/pkg/config"
	"github.com/nakabonne/golintui/pkg/golangcilint"
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
	logger := newLogger(conf)
	runner := golangcilint.NewRunner(conf.Executable, []string{}, logger)
	return &App{
		closers: []io.Closer{},
		Config:  conf,
		Log:     logger,
		Gui:     gui.New(logger, runner),
	}, nil
}

func (a *App) Run() error {
	return a.Gui.Run()
}
