package config

import (
	"encoding/json"
	"fmt"
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
		return nil, err
	}
	//log.Printf("Reading config file %v", configFile)
	d, err = ioutil.ReadFile(configFile)
	if err != nil {
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

func (c Config) GetPath(typ string, id string) string {
	s := fmt.Sprintf("%v", c.GetEndpoint())
	p := "/v1"
	return fmt.Sprintf("%v%v/%v", s, p, id)
}
