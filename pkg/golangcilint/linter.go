package golangcilint

import (
	"github.com/golangci/golangci-lint/pkg/report"
)

// Linter represents a linter available on golangci-lint.
type Linter struct {
	Name    string
	Enabled bool
}

// NewLinters converts LinterData represented internally into the opinionated Linter.
func NewLinters(linters []report.LinterData) []Linter {
	res := make([]Linter, 0, len(linters))
	for _, l := range linters {
		res = append(res, Linter{Name: l.Name, Enabled: l.Enabled})
	}
	return res
}
