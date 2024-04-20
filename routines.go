package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Metoda pro simulaci tankování u stojanu
func fuelStandRoutine(fs *FuelStand) {
	defer fuelStandWaiter.Done()
	fuelStandWaiter.Add(1)

	//fmt.Printf("Fuel stand %d is open\n", fs.Id)
	for car := range fs.Queue {
		car.QueueTime = time.Duration(time.Since(car.QueueEnter).Milliseconds())
		simulateFueling(car)
		car.carSync.Add(1)
		payStart := time.Now()
		buildingQueue <- car.carSync
		car.carSync.Wait()
		car.PayTime = time.Duration(time.Since(payStart).Milliseconds())
		exit <- car
	}
	//fmt.Printf("Fuel stand %d is closed\n", fs.Id)
}

// Metoda pro simulaci průběhu u kasy (placení)
func cashRegisterRoutine(cs *CashRegister) {
	defer cashRegisterWaiter.Done()
	cashRegisterWaiter.Add(1)

	//fmt.Printf("Cash register %d is open\n", cs.Id)
	for car := range cs.Queue {
		simulatePayment()
		car.Done()
	}
	//fmt.Printf("Cash register %d is closed\n", cs.Id)
}

// Metoda pro simulaci příjezdu aut
func carArrivalRoutine(config Config) {
	for i := 0; i < config.Cars.Count; i++ {
		var arrivalTimeMax = int(config.Cars.ArrivalTimeMax.Milliseconds())
		var arrivalTimeMin = int(config.Cars.ArrivalTimeMin.Milliseconds())
		arrivals <- &Car{ID: i, Fuel: generateRandomFuelType(), carSync: &sync.WaitGroup{}, QueueEnter: time.Now()}
		stagger := time.Duration(rand.Intn(arrivalTimeMax-arrivalTimeMin) + arrivalTimeMin)
		time.Sleep(stagger * time.Millisecond)
	}
	close(arrivals)
}

// Metoda pro agregaci statisti aut co natankují a zaplatí
func aggregationRoutine() {
	var totalCars int
	var printers Printers
	stats := Statistics{}
	var totalRegisterTime time.Duration
	var maxRegisterQueue time.Duration

	var totalGasTime time.Duration
	maxGasQueue := 0
	gasCount := 0

	var totalDieselTime time.Duration
	maxDieselQueue := 0
	dieselCount := 0

	var totalLPGTime time.Duration
	maxLPGQueue := 0
	lpgCount := 0

	var totalElectricTime time.Duration
	maxElectricQueue := 0
	electricCount := 0

	for car := range exit {
		totalCars++
		totalRegisterTime += car.PayTime
		switch car.Fuel {
		case Gas:
			totalGasTime += car.FuelTime
			gasCount++
			if car.QueueTime > 0 {
				maxGasQueue = int(car.QueueTime)
			}
		case Diesel:
			totalDieselTime += car.FuelTime
			dieselCount++
			if car.QueueTime > 0 {
				maxDieselQueue = int(car.QueueTime)
			}
		case LPG:
			totalLPGTime += car.FuelTime
			lpgCount++
			if car.QueueTime > 0 {
				maxLPGQueue = int(car.QueueTime)
			}
		case Electric:
			totalElectricTime += car.FuelTime
			electricCount++
			if car.QueueTime > 0 {
				maxElectricQueue = int(car.QueueTime)
			}
		}
		if car.QueueTime > maxRegisterQueue {
			maxRegisterQueue = car.QueueTime
		}
		fmt.Println("Car", car.ID, "Fuel", car.Fuel, "QueueTime", car.QueueTime, "FuelTime", car.FuelTime, "PayTime", car.PayTime)
	}
	averageGasTime := int(totalGasTime) / gasCount
	averageDieselTime := int(totalDieselTime) / dieselCount
	averageLPGTime := int(totalLPGTime) / lpgCount
	averageElectricTime := int(totalElectricTime) / electricCount
	averageRegisterTime := int(totalRegisterTime) / totalCars

	printers.printStatistics(totalCars, totalRegisterTime, maxRegisterQueue, totalGasTime, maxGasQueue, gasCount, averageGasTime, totalDieselTime, maxDieselQueue, dieselCount, averageDieselTime, totalLPGTime, maxLPGQueue, lpgCount, averageLPGTime, totalElectricTime, maxElectricQueue, electricCount, averageElectricTime, averageRegisterTime)

	stats.printStatsInOutputYaml(totalCars, totalRegisterTime, maxRegisterQueue, totalGasTime, time.Duration(maxGasQueue), gasCount, averageGasTime, totalDieselTime, time.Duration(maxDieselQueue), dieselCount, averageDieselTime, totalLPGTime, time.Duration(maxLPGQueue), lpgCount, averageLPGTime, totalElectricTime, time.Duration(maxElectricQueue), electricCount, averageElectricTime, averageRegisterTime)
	end.Done()
}
