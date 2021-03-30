package config

import "fmt"

// Config - holds the configuration of the application
type Config struct {
	Host     string       `mapstructure:"host"`
	HttpPort string       `mapstructure:"http_port"`
	Services ServicesConf `mapstructure:"Services"`
}

// Config - holds the http services used in the application
type ServicesConf struct {
	SpacexApi *SpacexApiConf `mapstructure:"spacex_api"`
}

// Config - holds the http information of SpaceX API
type SpacexApiConf struct {
	Host string `mapstructure:"Host"`
}

// GetAppAddress - provides the application host with the listening port
func (c *Config) GetAppAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.HttpPort)
}

// GetSpacexAddress - provides the graphql address of spacex
func (c *Config) GetSpacexAddress() string {
	return fmt.Sprintf("%s/graphql", c.Services.SpacexApi.Host)
}
