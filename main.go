// Autor: Adam Sucharda @ 2024
package main

import (
	"sync"
)

type FuelType string

var (
	config             Config
	exit               = make(chan *Car) // ukončení programu
	end                sync.WaitGroup
	buildingQueue      = make(chan *sync.WaitGroup, 10)
	fuelStandWaiter    sync.WaitGroup // synchronizace stanic
	standBuffer        = 2
	cashRegisterWaiter sync.WaitGroup // synchronizace registrů
	registerBuffer     = 3
	arrivals           = make(chan *Car, 20) // příjezdy aut
)

func main() {
	loader := ConfigLoader{Filename: "config.yaml"}
	config = loader.load()

	configPrinter := ConfigPrinter{Config: config}
	configPrinter.printConfig()

	// setup
	numGas := config.Stations.Gas.Count
	numDiesel := config.Stations.Diesel.Count
	numLPG := config.Stations.LPG.Count
	numElectric := config.Stations.Electric.Count
	numRegisters := config.Registers.Count

	var stands []*FuelStand
	standCount := 0

	// Inicializace
	setupStands(&stands, Gas, numGas, &standCount)
	setupStands(&stands, Diesel, numDiesel, &standCount)
	setupStands(&stands, LPG, numLPG, &standCount)
	setupStands(&stands, Electric, numElectric, &standCount)

	var registers []*CashRegister
	setupRegisters(&registers, numRegisters)

	end.Add(1)

	// Spuštění
	go carArrivalRoutine(config)
	for _, stand := range stands {
		go fuelStandRoutine(stand)
	}
	for _, register := range registers {
		go cashRegisterRoutine(register)
	}
	go findStandRoutine(stands)
	go findRegister(registers)
	go aggregationRoutine()

	fuelStandWaiter.Wait()
	close(buildingQueue)

	cashRegisterWaiter.Wait()
	close(exit)
	end.Wait()
}

// Funkce pro inicializaci stanic
func setupStands(stands *[]*FuelStand, fuelType FuelType, count int, standCount *int) {
	for i := 0; i < count; i++ {
		*stands = append(*stands, createNewFuelStand(*standCount, fuelType, standBuffer))
		*standCount++
	}
}

// Funkce pro inicializaci registrů
func setupRegisters(registers *[]*CashRegister, count int) {
	for i := 0; i < count; i++ {
		*registers = append(*registers, createNewCashRegister(i, registerBuffer))
	}
}
