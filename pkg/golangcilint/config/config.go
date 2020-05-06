package config

// Config is part of configuration for golangci-lint, especially regarding linters.
type Config struct {
	DisableAll bool
	// NOTE: Deprecated, soon be removed from golangci-lint.
	EnableAll bool

	// A map to indicate which linters are enabled.
	Linters map[string]Linter
}
