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

func NewConfig(name, version, commit, date, buildSource string, debuggingFlag bool) (*Config, error) {
	return &Config{
		Name:        "",
		Debug:       false,
		Version:     "",
		Commit:      "",
		BuildDate:   "",
		BuildSource: "",
	}, nil
}
