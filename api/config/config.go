package config

import "fmt"

type Config struct {
	Host     string       `mapstructure:"host"`
	HttpPort string       `mapstructure:"http_port"`
	Services ServicesConf `mapstructure:"Services"`
}

type ServicesConf struct {
	SpacexApi *SpacexApiConf `mapstructure:"spacex_api"`
}

type SpacexApiConf struct {
	Host string `mapstructure:"Host"`
}

func (c *Config) GetAppAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.HttpPort)
}

func (c *Config) GetSpacexAddress() string {
	return fmt.Sprintf("%s/graphql", c.Services.SpacexApi.Host)
}
