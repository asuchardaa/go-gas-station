package main

import (
	"fmt"
)

type ConfigPrinter struct {
	Config Config
}

func (cp *ConfigPrinter) PrintConfig() {
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
