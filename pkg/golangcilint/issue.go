package golangcilint

import (
	"fmt"

	"github.com/golangci/golangci-lint/pkg/result"
)

// Issue is a wrapper of the issue model represented golangci-lint internally.
type Issue struct {
	issue result.Issue
}

func NewIssues(issues []result.Issue) []Issue {
	res := make([]Issue, 0, len(issues))
	for _, i := range issues {
		res = append(res, Issue{issue: i})
	}
	return res
}

func (i *Issue) FromLinter() string {
	return i.issue.FromLinter
}

func (i *Issue) Message() string {
	return fmt.Sprintf("%s:%d:%d: %s",
		i.issue.FilePath(),
		i.issue.Line(),
		i.issue.Column(),
		i.issue.Text,
	)
}
