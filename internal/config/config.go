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
	Host string `json:"Host"`
	Port int    `json:"Port"`
	UserFile string  `json:"UserFile"`
	GroupFile string `json:"GroupFile"`
}

func InitConfig() (*Config, error) {
	return ReadConfig(DefaultConfigFile)
}

func ReadConfig(configFile string) (*Config, error) {
	// default config file
	if len(configFile) == 0 {
		configFile = DefaultConfigFile
	}
	d, err := ioutil.ReadFile(configFile)
	if err != nil {
		configFile = DefaultConfigFile
	}
	//log.Printf("Reading config file %v", configFile)
	d, err = ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	c := new(Config)
	err = json.Unmarshal(d, c)
	if err != nil {
		return nil, err
	}
	return c, nil
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
