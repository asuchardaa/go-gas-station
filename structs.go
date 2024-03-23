package main

import (
	"sync"
	"time"
)

const (
	Gas      = "gas"
	Diesel   = "diesel"
	LPG      = "LPG"
	Electric = "electric"
)

type Car struct {
	ID         int
	Fuel       FuelType
	QueueEnter time.Time
	QueueTime  time.Duration
	FuelTime   time.Duration
	PayTime    time.Duration
	carSync    *sync.WaitGroup
}

type FuelStand struct {
	Id    int
	Type  FuelType
	Queue chan *Car
}

type CashRegister struct {
	Id    int
	Queue chan *sync.WaitGroup
}

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
