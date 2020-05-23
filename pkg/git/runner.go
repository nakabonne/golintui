package git

const defaultGitExecutable = "git"

type Runner struct {
	// Path to the git executable.
	Executable string

	// Specifies the working directory of git
	workingDir string
}

func NewRunner(executable string) *Runner {
	if executable == "" {
		executable = defaultGitExecutable
	}
	return &Runner{Executable: executable}
}

func (r *Runner) ListCommits(limit int) []*Commit {
	return []*Commit{
		{
			SHA:     "aaabbbcccddd",
			Message: "foo",
		},
		{
			SHA:     "dddcccbbbaaa",
			Message: "bar",
		},
	}
}
