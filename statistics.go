package main

import (
	"fmt"
	"os"
	"time"
)

type Statistics struct {
}

// printStatsInOutputYaml Metoda pro zápis statistik do souboru output.yaml
func (p *Statistics) printStatsInOutputYaml(totalCars int, totalRegisterTime time.Duration, maxRegisterQueue time.Duration, totalGasTime time.Duration, maxGasQueue time.Duration, gasCount int, averageGasTime int, totalDieselTime time.Duration, maxDieselQueue time.Duration, dieselCount int, averageDieselTime int, totalLPGTime time.Duration, maxLPGQueue time.Duration, lpgCount int, averageLPGTime int, totalElectricTime time.Duration, maxElectricQueue time.Duration, electricCount int, averageElectricTime int, averageRegisterTime int) {
	file, err := os.Create("output.yaml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Writing to the file
	_, _ = fmt.Fprintf(file, "Stations:\n")
	_, _ = fmt.Fprintf(file, " gas:\n")
	_, _ = fmt.Fprintf(file, "   total_cars: %d\n", gasCount)
	_, _ = fmt.Fprintf(file, "   total_time: %dms\n", int(totalGasTime)/1000000)
	_, _ = fmt.Fprintf(file, "   avg_queue_time: %dms\n", averageGasTime/1000000)
	_, _ = fmt.Fprintf(file, "   max_queue_time: %dms\n", maxGasQueue)

	_, _ = fmt.Fprintf(file, " diesel:\n")
	_, _ = fmt.Fprintf(file, "   total_cars: %d\n", dieselCount)
	_, _ = fmt.Fprintf(file, "   total_time: %dms\n", int(totalDieselTime)/1000000)
	_, _ = fmt.Fprintf(file, "   avg_queue_time: %dms\n", averageDieselTime/1000000)
	_, _ = fmt.Fprintf(file, "   max_queue_time: %dms\n", maxDieselQueue)

	_, _ = fmt.Fprintf(file, " lpg:\n")
	_, _ = fmt.Fprintf(file, "   total_cars: %d\n", lpgCount)
	_, _ = fmt.Fprintf(file, "   total_time: %dms\n", int(totalLPGTime)/1000000)
	_, _ = fmt.Fprintf(file, "   avg_queue_time: %dms\n", averageLPGTime/1000000)
	_, _ = fmt.Fprintf(file, "   max_queue_time: %dms\n", maxLPGQueue)

	_, _ = fmt.Fprintf(file, " electric:\n")
	_, _ = fmt.Fprintf(file, "   total_cars: %d\n", electricCount)
	_, _ = fmt.Fprintf(file, "   total_time: %dms\n", int(totalElectricTime)/1000000)
	_, _ = fmt.Fprintf(file, "   avg_queue_time: %dms\n", averageElectricTime/1000000)
	_, _ = fmt.Fprintf(file, "   max_queue_time: %dms\n", maxElectricQueue)

	_, _ = fmt.Fprintf(file, "Registers:\n")
	_, _ = fmt.Fprintf(file, "  total_cars: %d\n", totalCars)
	_, _ = fmt.Fprintf(file, "  total_time: %ds\n", int(totalRegisterTime)/1000)
	_, _ = fmt.Fprintf(file, "  avg_queue_time: %dms\n", averageRegisterTime)
	_, _ = fmt.Fprintf(file, "  max_queue_time: %dms\n", maxRegisterQueue)
}
