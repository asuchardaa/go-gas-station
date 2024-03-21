package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type FuelType string

var config Config

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

var Exit = make(chan *Car)

// stand setups
var standWaiter sync.WaitGroup
var standBuffer = 2
var registerWaiter sync.WaitGroup
var registerBuffer = 3

// stand numbers
var numGas = config.Stations.Gas.Count
var numDiesel = config.Stations.Diesel.Count
var numLPG = config.Stations.LPG.Count
var numElectric = config.Stations.Electric.Count

var numRegisters = config.Registers.Count

func main() {
	loader := ConfigLoader{Filename: "config.yaml"}
	config = loader.Load()

	configPrinter := ConfigPrinter{Config: config}
	configPrinter.PrintConfig()

	numGas = config.Stations.Gas.Count
	numDiesel = config.Stations.Diesel.Count
	numLPG = config.Stations.LPG.Count
	numElectric = config.Stations.Electric.Count
	numRegisters = config.Registers.Count

	var stands []*FuelStand
	standCount := 0

	for i := 0; i < numGas; i++ {
		stands = append(stands, newFuelStand(standCount, Gas, standBuffer))
		standCount++
	}

	for i := 0; i < numDiesel; i++ {
		stands = append(stands, newFuelStand(standCount, Diesel, standBuffer))
		standCount++
	}

	for i := 0; i < numLPG; i++ {
		stands = append(stands, newFuelStand(standCount, LPG, standBuffer))
		standCount++
	}

	for i := 0; i < numElectric; i++ {
		stands = append(stands, newFuelStand(standCount, Electric, standBuffer))
		standCount++
	}

	var registers []*CashRegister
	for i := 0; i < numRegisters; i++ {
		registers = append(registers, newCashRegister(i, registerBuffer))
	}

	end.Add(1)
	go createCarsRoutine(config)
	for _, stand := range stands {
		go standRoutine(stand)
	}
	for _, register := range registers {
		go registerRoutine(register)
	}
	go findStandRoutine(stands)
	go findRegister(registers)
	go aggregationRoutine()

	standWaiter.Wait()
	close(buildingQueue)

	registerWaiter.Wait()
	close(Exit)

	end.Wait()
}

var arrivals = make(chan *Car, 20)

func createCarsRoutine(config Config) {
	for i := 0; i < config.Cars.Count; i++ {
		var arrivalTimeMax = int(config.Cars.ArrivalTimeMax.Milliseconds())
		var arrivalTimeMin = int(config.Cars.ArrivalTimeMin.Milliseconds())
		arrivals <- &Car{ID: i, Fuel: genFuelType(), carSync: &sync.WaitGroup{}, QueueEnter: time.Now()}
		stagger := time.Duration(rand.Intn(arrivalTimeMax-arrivalTimeMin) + arrivalTimeMin)
		time.Sleep(stagger * time.Millisecond)
	}
	close(arrivals)
}

func findStandRoutine(stands []*FuelStand) {
	for car := range arrivals {

		var bestStand *FuelStand
		bestQueueLength := -1

		for _, stand := range stands {
			if stand.Type == car.Fuel {
				queueLength := len(stand.Queue)
				if bestQueueLength == -1 || queueLength < bestQueueLength {
					bestStand = stand
					bestQueueLength = queueLength
				}
			}
		}
		bestStand.Queue <- car
	}
	for _, stand := range stands {
		close(stand.Queue)
	}
}

var buildingQueue = make(chan *sync.WaitGroup, 10)

func findRegister(registers []*CashRegister) {
	for car := range buildingQueue {
		var bestRegister *CashRegister
		bestQueueLength := -1
		for _, register := range registers {
			queueLength := len(register.Queue)
			if bestQueueLength == -1 || queueLength < bestQueueLength {
				bestRegister = register
				bestQueueLength = queueLength
			}
		}
		bestRegister.Queue <- car
	}
	for _, register := range registers {
		close(register.Queue)
	}
}

func newFuelStand(id int, fuel FuelType, bufferSize int) *FuelStand {
	return &FuelStand{
		Id:    id,
		Type:  fuel,
		Queue: make(chan *Car, bufferSize),
	}
}

func standRoutine(fs *FuelStand) {
	defer standWaiter.Done()
	standWaiter.Add(1)
	//fmt.Printf("Fuel stand %d is open\n", fs.Id)
	for car := range fs.Queue {
		car.QueueTime = time.Duration(time.Since(car.QueueEnter).Milliseconds())
		doFueling(car)
		car.carSync.Add(1)
		payStart := time.Now()
		buildingQueue <- car.carSync
		car.carSync.Wait()
		car.PayTime = time.Duration(time.Since(payStart).Milliseconds())
		Exit <- car
	}
	//fmt.Printf("Fuel stand %d is closed\n", fs.Id)
}

func doFueling(car *Car) {
	var gasMinT = int(config.Stations.Gas.ServeTimeMin.Milliseconds())
	var gasMaxT = int(config.Stations.Gas.ServeTimeMax.Milliseconds())
	var dieselMinT = int(config.Stations.Diesel.ServeTimeMin.Milliseconds())
	var dieselMaxT = int(config.Stations.Diesel.ServeTimeMax.Milliseconds())
	var lpgMinT = int(config.Stations.LPG.ServeTimeMin.Milliseconds())
	var lpgMaxT = int(config.Stations.LPG.ServeTimeMax.Milliseconds())
	var electricMinT = int(config.Stations.Electric.ServeTimeMin.Milliseconds())
	var electricMaxT = int(config.Stations.Electric.ServeTimeMax.Milliseconds())
	switch car.Fuel {
	case Gas:
		car.FuelTime = randomTime(gasMinT, gasMaxT)
	case Diesel:
		car.FuelTime = randomTime(dieselMinT, dieselMaxT)
	case LPG:
		car.FuelTime = randomTime(lpgMinT, lpgMaxT)
	case Electric:
		car.FuelTime = randomTime(electricMinT, electricMaxT)
	}
	doSleeping(car.FuelTime)
}

func newCashRegister(id, bufferSize int) *CashRegister {
	return &CashRegister{
		Id:    id,
		Queue: make(chan *sync.WaitGroup, bufferSize),
	}
}

func registerRoutine(cs *CashRegister) {
	defer registerWaiter.Done()
	registerWaiter.Add(1)
	//fmt.Printf("Cash register %d is open\n", cs.Id)
	for car := range cs.Queue {
		doPayment()
		car.Done()
	}
	//fmt.Printf("Cash register %d is closed\n", cs.Id)
}

func doPayment() {
	var minPaymentT = int(config.Registers.HandleTimeMin.Milliseconds())
	var maxPaymentT = int(config.Registers.HandleTimeMax.Milliseconds())
	doSleeping(randomTime(minPaymentT, maxPaymentT))
}

var end sync.WaitGroup

func aggregationRoutine() {
	var totalCars int
	var totalRegisterTime time.Duration
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
	maxRegisterQueue := 0

	for car := range Exit {
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
		if car.QueueTime > 0 {
			maxRegisterQueue = int(car.QueueTime)
		}
	}
	averageGasTime := int(totalGasTime) / gasCount
	averageDieselTime := int(totalDieselTime) / dieselCount
	averageLPGTime := int(totalLPGTime) / lpgCount
	averageElectricTime := int(totalElectricTime) / electricCount

	fmt.Println("---------------------------------------")
	fmt.Println("	     STATISTICS")
	fmt.Println("----------------------------------------")
	fmt.Println("Stations:")
	fmt.Println(" gas:")
	fmt.Printf("   total_cars: %d\n", gasCount)
	fmt.Printf("   total_time: %dms\n", int(totalGasTime))
	fmt.Printf("   avg_queue_time: %dms\n", averageGasTime)
	fmt.Printf("   max_queue_time: %dms\n", maxGasQueue)

	fmt.Println(" diesel:")
	fmt.Printf("   total_cars: %d\n", dieselCount)
	fmt.Printf("   total_time: %dms\n", int(totalDieselTime))
	fmt.Printf("   avg_queue_time: %dms\n", averageDieselTime)
	fmt.Printf("   max_queue_time: %dms\n", maxDieselQueue)

	fmt.Println(" lpg:")
	fmt.Printf("   total_cars: %d\n", lpgCount)
	fmt.Printf("   total_time: %dms\n", int(totalLPGTime))
	fmt.Printf("   avg_queue_time: %dms\n", averageLPGTime)
	fmt.Printf("   max_queue_time: %dms\n", maxLPGQueue)

	fmt.Println(" electric:")
	fmt.Printf("   total_cars: %d\n", electricCount)
	fmt.Printf("   total_time: %dms\n", int(totalElectricTime))
	fmt.Printf("   avg_queue_time: %dms\n", averageElectricTime)
	fmt.Printf("   max_queue_time: %dms\n", maxElectricQueue)

	averageRegisterTime := int(totalRegisterTime) / totalCars
	fmt.Println("registers:")
	fmt.Printf("  total_cars: %d\n", totalCars)
	fmt.Printf("  total_time: %ds\n", int(totalRegisterTime))
	fmt.Printf("  avg_queue_time: %ds\n", averageRegisterTime)
	fmt.Printf("  max_queue_time: %ds\n", maxRegisterQueue)
	end.Done()
}

func genFuelType() FuelType {
	fuelTypes := []FuelType{Gas, Diesel, Electric, LPG}
	randomIndex := rand.Intn(len(fuelTypes))
	return fuelTypes[randomIndex]
}

func randomTime(min, max int) time.Duration {
	generatedTime := time.Duration(rand.Intn(max-min) + min)
	return generatedTime
}

func doSleeping(delay time.Duration) {
	time.Sleep(delay * time.Millisecond)
}
