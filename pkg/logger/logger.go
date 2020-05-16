package logger

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"

	"github.com/nakabonne/golintui/pkg/config"
)

func NewLogger(cfg *config.Config, w io.Writer) *logrus.Entry {
	var log *logrus.Logger
	if cfg.GetDebug() {
		log = newLogger(cfg.CfgDir, w)
	} else {
		log = newNopLogger()
	}

	log.Formatter = &logrus.JSONFormatter{}
	l := log.WithFields(logrus.Fields{
		"debug":     cfg.GetDebug(),
		"version":   cfg.GetVersion(),
		"commit":    cfg.GetCommit(),
		"buildDate": cfg.GetBuildDate(),
	})
	l.Trace("successfully generated logger")
	return l
}

func newLogger(dir string, w io.Writer) *logrus.Logger {
	if w == nil {
		err := os.MkdirAll(dir, 0766)
		if err != nil {
			panic(err)
		}
		w, err = os.OpenFile(filepath.Join(dir, "debug.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	}

	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)
	log.SetOutput(w)
	pp.SetDefaultOutput(w)
	return log
}

// Generates a logger that doesn't anything.
func newNopLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.ErrorLevel)
	log.SetOutput(ioutil.Discard)
	return log
}
