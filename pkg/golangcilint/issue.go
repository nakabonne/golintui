package golangcilint

import (
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
