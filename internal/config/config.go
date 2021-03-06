// Package config deals with configuration
package config

import (
	"encoding/json"
	"fmt"
	"github.com/ivost/nixug/internal/test"
	"io/ioutil"
)

const (
	DefaultConfigFile = "config.json"
	DefaultUserFile   = "/etc/passwd"
	DefaultGroupFile  = "/etc/group"
)

type Listener struct {
}

type Config struct {
	Host      string `json:"Host"`
	Port      int    `json:"Port"`
	Auth      bool   `json:"Auth"`
	UserFile  string `json:"UserFile"`
	GroupFile string `json:"GroupFile"`
}

func NewConfig(configFile string) (*Config, error) {
	if len(configFile) == 0 {
		return NewDefaultConfig(), nil
	}
	return ReadConfig(configFile)
}

func NewDefaultConfig() *Config {
	return &Config{Host: "0.0.0.0", Port: 8080, UserFile: DefaultUserFile, GroupFile: DefaultGroupFile}
}

func ReadConfig(configFile string) (*Config, error) {
	d, err := ioutil.ReadFile(configFile)
	if err != nil {
		test.PrintCurrentDir()
		return nil, err
	}
	c := new(Config)
	err = json.Unmarshal(d, c)
	return c, err
}

func (c Config) GetHostPort() string {
	addr := fmt.Sprintf("%v:%v", c.Host, c.Port)
	return addr
}

func (c Config) GetEndpoint() string {
	s := fmt.Sprintf("http://%v", c.GetHostPort())
	return s
}
