package oscommand

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/k0kubun/pp"

	"github.com/sirupsen/logrus"
)

type OSCommand struct {
	OpenCommandEnv string
	Logger         *logrus.Entry
}

func NewOSCommand(openCommandEnv string, logger *logrus.Entry) *OSCommand {
	return &OSCommand{OpenCommandEnv: openCommandEnv, Logger: logger}
}

// OpenFileAtLineColumn opens a file at a specific line and column.
func (o *OSCommand) OpenFileAtLineColumn(filename string, line, column int) error {
	command := specifyLineColumn(o.openCommand(), filename, line, column)
	_, err := o.runCommand(command[0], command[1:]...)
	return err
}

func (o *OSCommand) openCommand() string {
	executable := os.Getenv(o.OpenCommandEnv)
	if executable == "" {
		executable = os.Getenv("EDITOR")
	}
	if executable == "" {
		vi, err := o.runCommand("which", "vi")
		if err != nil {
			o.Logger.Error("failed to get path to vi", err)
		}
		executable = string(vi)
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
		s := fmt.Sprintf("\"+normal %dG%d|\" %s", line, column, filename)
		res = append(res, strings.Split(s, " ")...)
	case "vim":
		s := fmt.Sprintf("\"+normal %dG%d|\" %s", line, column, filename)
		res = append(res, strings.Split(s, " ")...)
	case "nvim":
		s := fmt.Sprintf("\"+normal %dG%d|\" %s", line, column, filename)
		res = append(res, strings.Split(s, " ")...)
	case "emacs":
		s := fmt.Sprintf("+%d:%d %s", line, column, filename)
		res = append(res, strings.Split(s, " ")...)
	case "code":
		s := fmt.Sprintf("--goto %s:%d:%d", filename, line, column)
		res = append(res, strings.Split(s, " ")...)
	default:
		// Don't specify when using unsupported editor.
		res = append(res, filename)
	}
	return res
}

func (o *OSCommand) runCommand(executable string, args ...string) ([]byte, error) {
	pp.Println("command:", executable, "args:", args)
	cmd := exec.Command(executable, args...)
	return cmd.CombinedOutput()
}
