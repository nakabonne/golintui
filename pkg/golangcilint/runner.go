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
	// An arg can be a file name, a workingDir, and in addition,
	// `...` to analyze them recursively.
	Args   []string
	Config *config.Config

	// specifies the working directory of golangci-lint
	workingDir string
	logger     *logrus.Entry
}

func NewRunner(executable string, args []string, logger *logrus.Entry) *Runner {
	// TODO: Automatically read config from golangci settings file.
	return &Runner{
		Executable: executable,
		Args:       args,
		Config:     &config.Config{},
		workingDir: ".",
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
	outJSON, err := r.run(r.Args)
	if err != nil {
		return nil, err
	}

	var res printers.JSONResult
	if err := json.Unmarshal(outJSON, &res); err != nil {
		return nil, err
	}
	return NewIssues(res.Issues), nil
}

// ListLinters returns all linters, with settings about whether to enable or not.
func (r *Runner) ListLinters() ([]Linter, error) {
	tmpDir, cleaner, err := tmpProject()
	if err != nil {
		return nil, err
	}
	defer cleaner()

	outJSON, err := r.run([]string{fmt.Sprintf("./%s/%s", tmpDir, tmpGoFileName)})
	if err != nil {
		return nil, err
	}

	var res printers.JSONResult
	if err := json.Unmarshal(outJSON, &res); err != nil {
		return nil, err
	}

	if res.Report == nil {
		return nil, err
	}
	return NewLinters(res.Report.Linters), nil
}

func (r *Runner) GetVersion() string {
	version, err := r.execute("version")
	if err != nil {
		r.logger.Error(err)
		return ""
	}
	return string(version)
}

func (r *Runner) run(args []string) ([]byte, error) {
	out, err := r.execute(append([]string{"run", "--out-format=json", "--issues-exit-code=0"}, args...)...)
	if err != nil {
		r.logger.WithError(err).
			WithField("stderr", string(out)).
			Error("failed to run golangci-lint run")
		return nil, fmt.Errorf("%s: %w", string(out), err)
	}
	return out, err
}

func (r *Runner) execute(args ...string) ([]byte, error) {
	cmd := exec.Command(r.Executable, args...)
	cmd.Dir = r.workingDir
	r.logger.WithField("executable", r.Executable).WithField("args", args).Debug("run golangci-lint")
	return cmd.CombinedOutput()
}
