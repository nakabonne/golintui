package golangcilint

import "github.com/nakabonne/golintui/pkg/config"

const globOperator = "/..."

type Runner struct {
	// Args given to `golangci-lint run`.
	// An arg can be a file name, a dir, and in addition,
	// `...` to analyze them recursively.
	Args   []string
	Config *config.Config
}

func NewRunner(args []string) *Runner {
	// TODO: Automatically read config from golangci settings file.
	return &Runner{Args: args, Config: &config.Config{}}
}

func (r *Runner) AddArgs(arg string) {
	r.Args = append(r.Args, arg+globOperator)
}
