package main

import (
	"fmt"
	"time"
)

type ConfigPrinter struct {
	Config Config
}

type Printers struct {
}

// printConfig Metoda pro výtisk configu
func (cp *ConfigPrinter) printConfig() {
	fmt.Println("---------------------------------------")
	fmt.Println("	  GAS STATION CONFIG")
	fmt.Println("----------------------------------------")
	fmt.Println("Cars Count:", cp.Config.Cars.Count)
	fmt.Println(" Cars min arrival", cp.Config.Cars.ArrivalTimeMin)
	fmt.Println(" Cars max arrival", cp.Config.Cars.ArrivalTimeMax)
	fmt.Println("Gas Stations Count:", cp.Config.Stations.Gas.Count)
	fmt.Println(" Gas Stations Serve Time Min:", cp.Config.Stations.Gas.ServeTimeMin)
	fmt.Println(" Gas Stations Serve Time Max:", cp.Config.Stations.Gas.ServeTimeMax)
	fmt.Println("Registers Count:", cp.Config.Registers.Count)
	fmt.Println(" Registers Handle Time Min:", cp.Config.Registers.HandleTimeMin)
	fmt.Println(" Registers Handle Time Max:", cp.Config.Registers.HandleTimeMax)
}

// printStatistics Metoda pro výtisk statistik
func (p *Printers) printStatistics(totalCars int, totalRegisterTime time.Duration, maxRegisterQueue time.Duration, totalGasTime time.Duration, maxGasQueue int, gasCount int, averageGasTime int, totalDieselTime time.Duration, maxDieselQueue int, dieselCount int, averageDieselTime int, totalLPGTime time.Duration, maxLPGQueue int, lpgCount int, averageLPGTime int, totalElectricTime time.Duration, maxElectricQueue int, electricCount int, averageElectricTime int, averageRegisterTime int) {
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

	fmt.Println("registers:")
	fmt.Printf("  total_cars: %d\n", totalCars)
	fmt.Printf("  total_time: %ds\n", int(totalRegisterTime)/1000)
	fmt.Printf("  avg_queue_time: %dms\n", averageRegisterTime)
	fmt.Printf("  max_queue_time: %dms\n", maxRegisterQueue)
}