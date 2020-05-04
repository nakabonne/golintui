package golangcilint

import (
	"github.com/golangci/golangci-lint/pkg/report"
)

// Linter represents a linter available on golangci-lint.
type Linter struct {
	name    string
	enabled bool
}

// NewLinters converts LinterData represented internally into the opinionated Linter.
func NewLinters(linters []report.LinterData) []Linter {
	res := make([]Linter, 0, len(linters))
	for _, l := range linters {
		res = append(res, Linter{name: l.Name, enabled: l.Enabled})
	}
	return res
}

func (l *Linter) Name() string {
	return l.name
}

func (l *Linter) Enabled() bool {
	return l.enabled
}
