package editor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type Editor struct {
	OpenCommandEnvKey string
	Logger            *logrus.Entry
}

func NewEditor(openCommandEnv string, logger *logrus.Entry) *Editor {
	return &Editor{OpenCommandEnvKey: openCommandEnv, Logger: logger}
}

// OpenFileAtLineColumn opens a file at a specific line and column.
func (e *Editor) OpenFileAtLineColumn(filename string, line, column int) error {
	command := specifyLineColumn(e.openCommand(), filename, line, column)
	return e.run(command[0], command[1:]...)
}

// openCommand returns an executable editor command.
// Falling back to environment variable for golintui, EDITOR then vi.
func (e *Editor) openCommand() string {
	executable := os.Getenv(e.OpenCommandEnvKey)
	if executable == "" {
		executable = os.Getenv("EDITOR")
	}
	if executable == "" {
		vi, err := exec.LookPath("vi")
		if err != nil {
			e.Logger.Error("failed to get path to vi", err)
		}
		executable = vi
	}
	if executable == "" {
		// TODO: Populate platform defaults,
		//   win: 'cmd /c "start "" {{filename}}"'
		//   osx: 'open {{filename}}'
		//   linux: 'sh -c "xdg-open {{filename}} >/dev/null"'
	}
	return executable
}

// specifyLineColumn makes a command that specify line and column number.
func specifyLineColumn(command, filename string, line, column int) []string {
	res := strings.Split(command, " ")
	switch res[0] {
	case "vi":
		args := fmt.Sprintf("+%d %s", line, filename)
		res = append(res, strings.Split(args, " ")...)
	case "vim":
		args := fmt.Sprintf("+%d %s", line, filename)
		res = append(res, strings.Split(args, " ")...)
	case "nvim":
		args := fmt.Sprintf("+%d %s", line, filename)
		res = append(res, strings.Split(args, " ")...)
	case "emacs":
		args := fmt.Sprintf("+%d:%d %s", line, column, filename)
		res = append(res, strings.Split(args, " ")...)
	case "code":
		args := fmt.Sprintf("--goto %s:%d:%d", filename, line, column)
		res = append(res, strings.Split(args, " ")...)
	default:
		// Don't specify when using unsupported editor.
		res = append(res, filename)
	}
	return res
}

func (e *Editor) run(executable string, args ...string) error {
	cmd := exec.Command(executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	e.Logger.Debug("now start running editor")
	if err := cmd.Run(); err != nil {
		// not handling the error explicitly because usually we're going to see it in the output anyway
		e.Logger.Error(err)
	}
	e.Logger.Debug("finish to edit")

	// fmt.Fprintf(os.Stdout, "\n%s", utils.ColoredString("Press Enter", color.FgGreen))
	// fmt.Scanln() // wait for enter press
	return nil
}
