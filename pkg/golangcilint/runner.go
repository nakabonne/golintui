package golangcilint

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	"github.com/golangci/golangci-lint/pkg/printers"
	"github.com/sirupsen/logrus"

	"github.com/nakabonne/golintui/pkg/golangcilint/config"
)

const globOperator = "/..."

type Runner struct {
	// Path to a golangci-lint executable.
	Executable string
	// Args given to `golangci-lint run`.
	// An arg can be a file name, a workingDir, and in addition,
	// `...` to analyze them recursively.
	Args []string
	// Path to config file for golangci-lint.
	// The Supported formats are yaml, json and toml.
	//ConfigPath string
	// A map to indicate which linters are enabled.
	Linters map[string]Linter

	// Specifies the working directory of golangci-lint
	workingDir string
	logger     *logrus.Entry
	cfg        *config.Config
}

func NewRunner(executable string, args []string, logger *logrus.Entry) (*Runner, error) {
	r := &Runner{
		Executable: executable,
		Args:       args,
		workingDir: ".",
		logger:     logger,
		cfg:        config.NewConfig(),
	}
	err := r.initLinters()
	return r, err
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
	if err := r.cfg.ReadConfig(); err != nil {
		return nil, err
	}
	b, err := r.cfg.ToYAML()
	if err != nil {
		return nil, err
	}
	confPath, clean, err := tmpConfigFile(b)
	if err != nil {
		return nil, err
	}
	defer clean()
	r.Args = append(r.Args, "--config", confPath)
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
func (r *Runner) ListLinters() []Linter {
	res := make([]Linter, 0, len(r.Linters))
	for _, linter := range r.Linters {
		res = append(res, linter)
	}
	return res
}

func (r *Runner) EnableLinter(linterName string) {
	linter, ok := r.Linters[linterName]
	if !ok {
		r.logger.WithField("linter", linterName).Error("linter not found")
		return
	}
	linter.Enable()
	r.Linters[linterName] = linter
}

func (r *Runner) DisableLinter(linterName string) {
	linter, ok := r.Linters[linterName]
	if !ok {
		r.logger.WithField("linter", linterName).Error("linter not found")
		return
	}
	linter.Disable()
	r.Linters[linterName] = linter
}

func (r *Runner) GetVersion() string {
	version, err := r.execute("version")
	if err != nil {
		r.logger.Error(err)
		return ""
	}
	return string(version)
}

func (r *Runner) run(targets []string) ([]byte, error) {
	args := []string{"run", "--out-format=json", "--issues-exit-code=0"}

	// Specify enabled linters
	linters := []string{}
	for _, l := range r.Linters {
		if l.Enabled() {
			linters = append(linters, "-E", l.Name())
		}
	}
	if len(linters) != 0 {
		linters = append(linters, "--disable-all")
	}
	args = append(args, linters...)

	out, err := r.execute(append(args, targets...)...)
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
	r.logger.WithField("executable", r.Executable).WithField("args", args).Debug("start running golangci-lint")
	return cmd.CombinedOutput()
}

// initLinters sets linters applied for the current directory.
func (r *Runner) initLinters() error {
	tmpDir, cleaner, err := tmpProject()
	if err != nil {
		return err
	}
	defer cleaner()

	// Run against tmp project to fetch linters information.
	outJSON, err := r.run([]string{fmt.Sprintf("./%s/%s", tmpDir, tmpGoFileName)})
	if err != nil {
		return err
	}

	var res printers.JSONResult
	if err := json.Unmarshal(outJSON, &res); err != nil {
		return err
	}

	if res.Report == nil {
		return errors.New("wrong result was returned from golangci-lint")
	}
	linters := NewLinters(res.Report.Linters)
	r.Linters = make(map[string]Linter, len(linters))
	for _, l := range linters {
		r.Linters[l.Name()] = l
	}
	return nil
}
