package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	ConfigFile = "config.json"
	DefaultUserFile = "/etc/passwd"
	DefaultGroupFile = "/etc/group"
)


type Listener struct {
	Port int    `json:"port"`
	Addr string `json:"addr"`
}

type Config struct {
	Listener      Listener    `json:"Listener"`

	UserFile string  `json:"UserFile"`
	GroupFile string `json:"GroupFile"`
}

func InitConfig() (*Config, error) {
	return ReadConfig(ConfigFile)
}

func ReadConfig(configFile string) (*Config, error) {
	// default config file
	if len(configFile) == 0 {
		configFile = ConfigFile
	}
	d, err := ioutil.ReadFile(configFile)
	if err != nil {
		configFile = ConfigFile
	}
	log.Printf("Reading config file %v", configFile)
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
	addr := fmt.Sprintf("%v:%v", c.Listener.Addr, c.Listener.Port)
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
