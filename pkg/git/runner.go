package git

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/nakabonne/golintui/pkg/str"
)

const defaultGitExecutable = "git"

type Runner struct {
	// Path to the git executable.
	Executable string

	// Specifies the working directory of git
	workingDir string
	logger     *logrus.Entry
}

func NewRunner(executable string, logger *logrus.Entry) *Runner {
	if executable == "" {
		executable = defaultGitExecutable
	}
	return &Runner{Executable: executable, logger: logger}
}

func (r *Runner) ListCommits(limit int) ([]*Commit, error) {
	args := fmt.Sprintf("log --pretty=format:'%s' -%d", prettyCommitFormat, limit)
	cmd := r.makeCmd(args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		r.logger.Error(err.Error())
		return nil, fmt.Errorf(string(out))
	}
	// Clean the dirty output to be JSON format.
	outStr := string(out)
	outStr = strings.ReplaceAll(outStr, `"`, `\"`)
	outStr = strings.ReplaceAll(outStr, "^^^^", `"`)
	arrayJSON := fmt.Sprintf("[%s]", strings.TrimRight(outStr, ","))

	var commits []*Commit
	err = json.Unmarshal([]byte(arrayJSON), &commits)
	if err != nil {
		return nil, err
	}
	return commits, nil
}

func (r *Runner) makeCmd(argsStr string) *exec.Cmd {
	args := str.ToArgv(argsStr)
	cmd := exec.Command(r.Executable, args...)
	cmd.Env = append(os.Environ(), "GIT_OPTIONAL_LOCKS=0")
	cmd.Dir = r.workingDir
	return cmd
}
