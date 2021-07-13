package healthcheck

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Services []Service
}

type Service struct {
	Name string
	Url  string

	Request struct {
		Method string
		Header []struct {
			Key   string
			Value string
		}
	}
}

var MainConfig = &Config{}

func ReadConfig() *Config {
	data, _ := ioutil.ReadFile("application.yaml")
	_ = yaml.Unmarshal(data, MainConfig)

	return MainConfig
}
