package main

import (
	"fmt"
	"io"
	"os"
	"runtime"

	flag "github.com/spf13/pflag"

	"github.com/nakabonne/golintui/pkg/config"
	"github.com/nakabonne/golintui/pkg/editor"
	"github.com/nakabonne/golintui/pkg/golangcilint"
	"github.com/nakabonne/golintui/pkg/gui"
	"github.com/nakabonne/golintui/pkg/logger"
)

var (
	flagSet = flag.NewFlagSet("golintui", flag.ContinueOnError)

	usage = func() {
		fmt.Fprintln(os.Stderr, "usage: golintui [<flag> ...]")
		flagSet.PrintDefaults()
	}
	// Automatically populated by goreleaser during build
	version = "unversioned"
	commit  = ""
	date    = ""
)

type cli struct {
	debugFlag   bool
	versionFlag bool
	executable  string
	stdout      io.Writer
	stderr      io.Writer
}

func main() {
	c := &cli{
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	flagSet.BoolVarP(&c.versionFlag, "version", "v", false, "print the current version")
	flagSet.BoolVar(&c.debugFlag, "debug", false, "run in debug mode")
	flagSet.StringVarP(&c.executable, "executable", "e", "", "path to golangci-lint executable")
	flagSet.Usage = usage
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintln(c.stderr, err)
		}
		return
	}

	os.Exit(c.run())
}

func (c *cli) run() int {
	if c.versionFlag {
		fmt.Fprintf(c.stderr, "version=%s, os=%s, arch=%s\n", version, runtime.GOOS, runtime.GOARCH)
		return 0
	}

	conf := config.New("golintui", version, commit, date, "", c.executable, "", c.debugFlag)
	logger := logger.NewLogger(conf)
	runner, err := golangcilint.NewRunner(conf.Executable, []string{}, logger)
	if err != nil {
		fmt.Fprintln(c.stderr, err.Error())
		return 1
	}
	editor := editor.NewEditor(conf.OpenCommandEnv, logger)
	g, err := gui.New(logger, runner, editor)
	if err != nil {
		fmt.Fprintln(c.stderr, err.Error())
		return 1
	}

	if err := g.Run(); err != nil {
		fmt.Fprintln(c.stderr, err.Error())
		return 1
	}

	return 0
}
