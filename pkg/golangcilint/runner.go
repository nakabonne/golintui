package golangcilint

import (
	"encoding/json"
	"os/exec"

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
	dir string
}

func NewRunner(executable string, args []string) *Runner {
	// TODO: Automatically read config from golangci settings file.
	return &Runner{
		Executable: executable,
		Args:       args,
		Config:     &config.Config{},
		dir:        ".",
	}
}

func (r *Runner) AddArgs(arg string) {
	r.Args = append(r.Args, arg+globOperator)
}

// Run executes `golangci-lint run` with its own args and configuration.
func (r *Runner) Run(arg string) ([]Issue, error) {
	outJSON, err := r.execute("run", true, r.Args...)
	if err != nil {
		return nil, err
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

func (r *Runner) execute(subCommand string, outJSON bool, args ...string) ([]byte, error) {
	if outJSON {
		args = append(args, "--out-format=json")
	}
	cmd := exec.Command(r.Executable, args...)
	cmd.Dir = r.dir
	return cmd.CombinedOutput()
}
