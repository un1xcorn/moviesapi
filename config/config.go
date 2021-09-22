package config

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Host string
	Port string
}

func LoadOrDefault(name string) *Config {

	file, err := ioutil.ReadFile(name)
	if err != nil {
		log.Warn("Requested config file does not exist.")
		return &Config{Host: "127.0.0.1", Port: "8080"}
	}

	var c Config
	err = json.Unmarshal(file, &c)
	if err != nil {
		log.Warn("An error happened while parsing the requested config file.")
		return &Config{Host: "127.0.0.1", Port: "8080"}
	}
	return &c

}
