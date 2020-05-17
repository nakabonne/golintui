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

func (i *Issue) Message() string {
	// TODO: Append Replacement as well.
	return fmt.Sprintf("%s:%d:%d: %s",
		i.issue.FilePath(),
		i.issue.Line(),
		i.issue.Column(),
		i.issue.Text,
	)
}

func (i *Issue) FromLinter() string {
	return i.issue.FromLinter
}

func (i *Issue) FilePath() string {
	return i.issue.FilePath()
}

func (i *Issue) Line() int {
	return i.issue.Line()
}

func (i *Issue) Column() int {
	return i.issue.Column()
}

func (i *Issue) SourceLine() string {
	if len(i.issue.SourceLines) < 1 {
		return ""
	}
	return i.issue.SourceLines[0]
}
