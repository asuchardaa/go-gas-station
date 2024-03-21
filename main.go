package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

// Constants for fuel types
const (
	Gas      = "gas"
	Diesel   = "diesel"
	LPG      = "LPG"
	Electric = "electric"
)

type Config struct {
	Cars      CarConfig       `yaml:"cars"`
	Stations  StationsConfig  `yaml:"stations"`
	Registers RegistersConfig `yaml:"registers"`
}

type CarConfig struct {
	Count          int           `yaml:"count"`
	ArrivalTimeMin time.Duration `yaml:"arrival_time_min"`
	ArrivalTimeMax time.Duration `yaml:"arrival_time_max"`
}

type StationsConfig struct {
	Gas      Station `yaml:"gas"`
	Diesel   Station `yaml:"diesel"`
	LPG      Station `yaml:"lpg"`
	Electric Station `yaml:"electric"`
}

type Station struct {
	Count        int           `yaml:"count"`
	ServeTimeMin time.Duration `yaml:"serve_time_min"`
	ServeTimeMax time.Duration `yaml:"serve_time_max"`
}

type RegistersConfig struct {
	Count         int           `yaml:"count"`
	HandleTimeMin time.Duration `yaml:"handle_time_min"`
	HandleTimeMax time.Duration `yaml:"handle_time_max"`
}

func main() {
	config := loadConfig("config.yaml")

	fmt.Println("Cars Count:", config.Cars.Count)
	fmt.Println("Gas Stations Count:", config.Stations.Gas.Count)
	fmt.Println("Gas Stations Serve Time Min:", config.Stations.Gas.ServeTimeMin)
	fmt.Println("Gas Stations Serve Time Max:", config.Stations.Gas.ServeTimeMax)
	fmt.Println("Registers Count:", config.Registers.Count)
	fmt.Println("Registers Handle Time Min:", config.Registers.HandleTimeMin)
	fmt.Println("Registers Handle Time Max:", config.Registers.HandleTimeMax)

}

func loadConfig(filename string) Config {
	data, err := ioutil.ReadFile(filename)
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
