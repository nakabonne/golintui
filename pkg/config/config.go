package config

const defaultExecutable = "golangci-lint"

// Config includes the base configuration fields required for golintui.
type Config struct {
	Name        string
	Debug       bool
	Version     string
	Commit      string
	BuildDate   string
	BuildSource string

	// Path to a golangci-lint executable.
	Executable string
	// UserConfig    *viper.Viper
	// UserConfigDir string
	// IsNewRepo bool
}

func New(name, version, commit, date, buildSource, executable string, debuggingFlag bool) (*Config, error) {
	if executable == "" {
		executable = defaultExecutable
	}
	return &Config{
		Name:        "",
		Debug:       true,
		Version:     "",
		Commit:      "",
		BuildDate:   "",
		BuildSource: "",
		Executable:  executable,
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
