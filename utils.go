package main

import (
	"math/rand"
	"sync"
	"time"
)

// Metoda pro náhodný výběr paiva pro příchozí auto
func generateRandomFuelType() FuelType {
	fuelTypes := []FuelType{Gas, Diesel, Electric, LPG}
	randomIndex := rand.Intn(len(fuelTypes))
	return fuelTypes[randomIndex]
}

// Metoda pro generování random času -> používám při tankování a placení
func generateRandomTime(min, max int) time.Duration {
	return time.Duration(rand.Intn(max-min) + min)
}

// Metoda pro simulaci placení u pokladny
func simulatePayment() {
	minPayment := int(config.Registers.HandleTimeMin.Milliseconds())
	maxPayment := int(config.Registers.HandleTimeMax.Milliseconds())
	delay := generateRandomTime(minPayment, maxPayment)
	sleepFor(delay)
}

// metoda pro spinkání :)
func sleepFor(duration time.Duration) {
	time.Sleep(duration * time.Millisecond)
}

// Metoda pro vytvoření nového stojanu
func createNewFuelStand(id int, fuel FuelType, bufferSize int) *FuelStand {
	return &FuelStand{
		Id:    id,
		Type:  fuel,
		Queue: make(chan *Car, bufferSize),
	}
}

// Metoda pro simulaci tankování auta
func simulateFueling(car *Car) {
	var minServeTime, maxServeTime int

	switch car.Fuel {
	case Gas:
		minServeTime = int(config.Stations.Gas.ServeTimeMin.Milliseconds())
		maxServeTime = int(config.Stations.Gas.ServeTimeMax.Milliseconds())
	case Diesel:
		minServeTime = int(config.Stations.Diesel.ServeTimeMin.Milliseconds())
		maxServeTime = int(config.Stations.Diesel.ServeTimeMax.Milliseconds())
	case LPG:
		minServeTime = int(config.Stations.LPG.ServeTimeMin.Milliseconds())
		maxServeTime = int(config.Stations.LPG.ServeTimeMax.Milliseconds())
	case Electric:
		minServeTime = int(config.Stations.Electric.ServeTimeMin.Milliseconds())
		maxServeTime = int(config.Stations.Electric.ServeTimeMax.Milliseconds())
	}

	car.FuelTime = generateRandomTime(minServeTime, maxServeTime)
	sleepFor(car.FuelTime)
}

// metoda pro vytvoření nové kasy
func createNewCashRegister(id, bufferSize int) *CashRegister {
	return &CashRegister{
		Id:    id,
		Queue: make(chan *sync.WaitGroup, bufferSize),
	}
}
