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

	// Specifies the working directory of golangci-lint
	workingDir string
	logger     *logrus.Entry
	cfg        config.Config
}

func NewRunner(executable string, args []string, logger *logrus.Entry) (*Runner, error) {
	r := &Runner{
		Executable: executable,
		Args:       args,
		workingDir: ".",
		logger:     logger,
		cfg:        config.Config{},
	}
	err := r.initConfig()
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
func (r *Runner) ListLinters() []config.Linter {
	res := make([]config.Linter, 0, len(r.cfg.Linters))
	for _, linter := range r.cfg.Linters {
		res = append(res, linter)
	}
	return res
}

func (r *Runner) EnableLinter(linterName string) {
	linter, ok := r.cfg.Linters[linterName]
	if !ok {
		r.logger.WithField("linter", linterName).Error("linter not found")
		return
	}
	linter.Enable()
	r.cfg.Linters[linterName] = linter
}

func (r *Runner) DisableLinter(linterName string) error {
	linter, ok := r.cfg.Linters[linterName]
	if !ok {
		r.logger.WithField("linter", linterName).Error("linter not found")
		return nil
	}
	if linter.EnabledByConfig() && r.cfg.DisableAll {
		return fmt.Errorf("can't disable '%s' linter because 'disable-all' is specified in the config file", linter.Name())
	}
	linter.Disable()
	r.cfg.Linters[linterName] = linter
	return nil
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
	for _, l := range r.cfg.Linters {
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

// initConfig sets config of linters applied for the current directory.
// First up, run golangci-lint against a temporary Go project to see which linter is enabled or not.
// Then start to read the configuration file.
func (r *Runner) initConfig() error {
	tmpDir, cleaner, err := tmpProject()
	if err != nil {
		return err
	}
	defer cleaner()

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
	linters := config.NewLinters(res.Report.Linters)
	r.cfg.Linters = make(map[string]config.Linter, len(linters))
	for _, l := range linters {
		r.cfg.Linters[l.Name()] = l
	}

	reader := config.NewReader(&r.cfg, r.logger)
	return reader.Read()
}
