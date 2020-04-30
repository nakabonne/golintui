package golangcilint

import "github.com/nakabonne/golintui/pkg/config"

type Runner struct {
	Args   []string
	Config *config.Config
}

func NewRunner(args []string) *Runner {
	// TODO: Automatically read config from golangci settings file.
	return &Runner{Args: args, Config: &config.Config{}}
}
