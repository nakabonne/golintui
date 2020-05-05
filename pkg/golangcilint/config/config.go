package config

import (
	"fmt"

	gconfig "github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/logutils"
	"github.com/golangci/golangci-lint/pkg/report"
	yaml "gopkg.in/yaml.v2"
)

// Config wraps golagnci-lint's Config
type Config struct {
	config *gconfig.Config
}

func NewConfig() *Config {
	return &Config{
		config: gconfig.NewDefault(),
	}
}

// ReadConfig parses user's golangci-lint config and set its pointer.
func (c *Config) ReadConfig() error {
	// TODO: Use other logger
	log := report.NewLogWrapper(logutils.NewStderrLog(""), &report.Data{})
	reader := gconfig.NewFileReader(c.config, nil, log)
	if err := reader.Read(); err != nil {
		return fmt.Errorf("can't read config: %w", err)
	}
	return nil
}

func (c *Config) ToYAML() ([]byte, error) {
	return yaml.Marshal(c.config)
}
