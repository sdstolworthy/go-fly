package skyscanner

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config contains application secrets
type Config struct {
	MashapeKey  string `yaml:"mashape_key"`
	MashapeHost string `yaml:"mashape_host"`
}

func (c *Config) getConfig() *Config {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
