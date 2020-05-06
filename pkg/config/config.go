package config

const (
	defaultExecutable     = "golangci-lint"
	defaultOpenCommandEnv = "GOLINTUI_OPEN_COMMAND"
)

// Config includes the base configuration fields required for golintui.
type Config struct {
	Name           string
	Debug          bool
	Version        string
	Commit         string
	Date           string
	BuildDate      string
	BuildSource    string
	OpenCommandEnv string

	// Path to a golangci-lint executable.
	Executable string
}

func New(name, version, commit, date, buildSource, executable, openCommandEnv string, debuggingFlag bool) (*Config, error) {
	if executable == "" {
		executable = defaultExecutable
	}
	if openCommandEnv == "" {
		openCommandEnv = defaultOpenCommandEnv
	}
	return &Config{
		Name:           name,
		Debug:          debuggingFlag,
		Version:        version,
		Commit:         commit,
		Date:           date,
		BuildDate:      "",
		BuildSource:    buildSource,
		Executable:     executable,
		OpenCommandEnv: openCommandEnv,
	}, nil
}

func (c *Config) GetDebug() bool {
	return c.Debug
}

func (c *Config) GetVersion() string {
	return c.Version
}

func (c *Config) GetCommit() string {
	return c.Commit
}

func (c *Config) GetBuildDate() string {
	return c.BuildDate
}
