package config

import (
	"github.com/spf13/viper"
)

// Load - extract configuration values from the provided file.
// It should have an extension supported by viper
func Load(configName string) (*Config, error) {
	viper.SetConfigFile(configName)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	C := &Config{}
	if err := viper.Unmarshal(C); err != nil {
		return nil, err
	}

	return C, nil
}
