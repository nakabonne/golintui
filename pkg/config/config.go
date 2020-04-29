package config

// Config includes the base configuration fields required for golintui.
type Config struct {
	Name        string
	Debug       bool
	Version     string
	Commit      string
	BuildDate   string
	BuildSource string
	// UserConfig    *viper.Viper
	// UserConfigDir string
	// IsNewRepo bool
}

func New(name, version, commit, date, buildSource string, debuggingFlag bool) (*Config, error) {
	return &Config{
		Name:        "",
		Debug:       false,
		Version:     "",
		Commit:      "",
		BuildDate:   "",
		BuildSource: "",
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
