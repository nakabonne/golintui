package golangcilint

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"

	"github.com/golangci/golangci-lint/pkg/printers"

	"github.com/nakabonne/golintui/pkg/config"
)

const globOperator = "/..."

type Runner struct {
	// Path to a golangci-lint executable.
	Executable string
	// Args given to `golangci-lint run`.
	// An arg can be a file name, a dir, and in addition,
	// `...` to analyze them recursively.
	Args   []string
	Config *config.Config

	// dir specifies the working directory.
	dir    string
	logger *logrus.Entry
}

func NewRunner(executable string, args []string, logger *logrus.Entry) *Runner {
	// TODO: Automatically read config from golangci settings file.
	return &Runner{
		Executable: executable,
		Args:       args,
		Config:     &config.Config{},
		dir:        ".",
		logger:     logger,
	}
}

func (r *Runner) AddArgs(arg string) {
	r.Args = append(r.Args, arg+globOperator)
}
func (r *Runner) RemoveArgs(arg string) {
	args := make([]string, 0, len(r.Args)-1)
	for _, a := range r.Args {
		if a != arg+globOperator {
			args = append(args, a)
		}
	}
	r.Args = args
}

// Run executes `golangci-lint run` with its own args and configuration.
func (r *Runner) Run() ([]Issue, error) {
	outJSON, err := r.execute(append([]string{"run", "--out-format=json", "--issues-exit-code=0"}, r.Args...)...)
	if err != nil {
		r.logger.WithError(err).
			WithField("stderr", string(outJSON)).
			Error("failed to run golangci-lint run")
		return nil, fmt.Errorf("%s: %w", string(outJSON), err)
	}

	var res printers.JSONResult
	if err := json.Unmarshal(outJSON, &res); err != nil {
		return nil, err
	}
	return NewIssues(res.Issues), nil
}

func (r *Runner) ListLinters() []Linter {
	// TODO: First up, run `golangci-lint run --out-format=json` against safety dir.
	//   And then fetch linters from Report.Linters.
	return []Linter{}
}

func (r *Runner) GetVersion() string {
	version, err := r.execute("version")
	if err != nil {
		r.logger.Error(err)
		return ""
	}
	return string(version)
}

func (r *Runner) execute(args ...string) ([]byte, error) {
	cmd := exec.Command(r.Executable, args...)
	cmd.Dir = r.dir
	return cmd.CombinedOutput()
}
