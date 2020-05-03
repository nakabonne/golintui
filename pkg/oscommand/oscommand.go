package oscommand

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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
	//_, err := o.runCommand(command[0], command[1:]...)
	return o.runSubprocess(command[0], command[1:]...)
	//return err
}

// openCommand returns an executable editor command.
// Falling back to environment variable for golintui, EDITOR then vi.
func (o *OSCommand) openCommand() string {
	executable := os.Getenv(o.OpenCommandEnv)
	if executable == "" {
		executable = os.Getenv("EDITOR")
	}
	if executable == "" {
		vi, err := exec.LookPath("vi")
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

func (o *OSCommand) runCommand(executable string, args ...string) ([]byte, error) {
	cmd := exec.Command(executable, args...)
	return cmd.CombinedOutput()
}

func (o *OSCommand) runSubprocess(executable string, args ...string) error {
	cmd := exec.Command(executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	//fmt.Fprintf(os.Stdout, "\n%s\n\n", utils.ColoredString("+ "+strings.Join(cmd.Args, " "), color.FgBlue))

	o.Logger.Info("now start running!")
	if err := cmd.Run(); err != nil {
		// not handling the error explicitly because usually we're going to see it in the output anyway
		o.Logger.Error(err)
	}
	o.Logger.Info("now finish!")

	/*	cmd.Stdout = ioutil.Discard
		cmd.Stderr = ioutil.Discard
		cmd.Stdin = nil
		cmd = nil

		fmt.Fprintf(os.Stdout, "\n%s", utils.ColoredString("Press Enter", color.FgGreen))
		fmt.Scanln() // wait for enter press
	*/
	return nil
}
