package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfigLoader struct {
	Filename string
}

// Metoda pro načítání konfigurace (.yaml)
func (cl *ConfigLoader) load() Config {
	data, err := ioutil.ReadFile(cl.Filename)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config data: %v", err)
	}

	return config
}
