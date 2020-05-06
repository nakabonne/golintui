package config

import (
	"fmt"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Reader struct {
	cfg    *Config
	logger *logrus.Entry
}

// Mapper is used for temporary mapping to config file.
type Mapper struct {
	Linters struct {
		EnableAll  bool `mapstructure:"enable-all"`
		DisableAll bool `mapstructure:"disable-all"`
	}
}

func NewReader(cfg *Config, logger *logrus.Entry) *Reader {
	return &Reader{cfg: cfg, logger: logger}
}

// Read parses user's golangci-lint config and populate its own Config pointer.
func (r *Reader) Read() error {
	r.addPathsToSearch()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			r.logger.Debug("config not found")
			return nil
		}
		return fmt.Errorf("can't read viper config: %s", err)
	}

	r.logger.Debug(fmt.Sprintf("used config file is %s", viper.ConfigFileUsed()))
	mapper := &Mapper{}
	if err := viper.Unmarshal(mapper); err != nil {
		return fmt.Errorf("can't unmarshal config by viper: %s", err)
	}
	r.cfg.DisableAll = mapper.Linters.DisableAll
	r.cfg.EnableAll = mapper.Linters.EnableAll
	return nil
}

func (r *Reader) addPathsToSearch() {
	// find all dirs from it up to the root
	configSearchPaths := []string{"./"}
	curDir := "."
	for {
		configSearchPaths = append(configSearchPaths, curDir)
		newCurDir := filepath.Dir(curDir)
		if curDir == newCurDir || newCurDir == "" {
			break
		}
		curDir = newCurDir
	}

	r.logger.WithField("paths", configSearchPaths).Debug("Add paths for viper to search")
	viper.SetConfigName(".golangci")
	for _, p := range configSearchPaths {
		viper.AddConfigPath(p)
	}
}
