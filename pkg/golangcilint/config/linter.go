package config

import (
	"github.com/golangci/golangci-lint/pkg/report"
)

// Linter represents a linter available on golangci-lint.
type Linter struct {
	name            string
	enabled         bool
	enabledByConfig bool
}

// NewLinters converts LinterData represented internally into the opinionated Linter.
func NewLinters(linters []report.LinterData) []Linter {
	res := make([]Linter, 0, len(linters))
	for _, l := range linters {
		res = append(res, Linter{
			name:            l.Name,
			enabled:         l.Enabled,
			enabledByConfig: l.Enabled,
		})
	}
	return res
}

func (l *Linter) Name() string {
	return l.name
}

func (l *Linter) Enabled() bool {
	return l.enabled
}

func (l *Linter) EnabledByConfig() bool {
	return l.enabledByConfig
}

// Enable makes itself enabled.
func (l *Linter) Enable() {
	l.enabled = true
}

// Disable makes itself disabled.
func (l *Linter) Disable() {
	l.enabled = false
}
