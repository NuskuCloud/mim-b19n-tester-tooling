package main

import (
	"flag"
	"fmt"
	"github.com/nuskucloud/samsung_mimb19n"
	"math"
	"os"
)

func main() {
	flowTemperature := flag.Float64("flow_temperature", -1000.0, "Flow temperature")
	targetTemperature := flag.Float64("target_temperature", -1000.0, "Target temperature")
	flag.Parse()
	if *flowTemperature == -1000.0 || *targetTemperature == -1000 {
		// dump
		fmt.Printf("Usage: %s -flow_temperature <flow_temperature> -target_temperature <target_temperature>\n", os.Args[0])
		os.Exit(1)
	}

	targetTemperatureUint16, _ := float64PtrToUint16(targetTemperature)
	flowTemperatureUint16, _ := float64PtrToUint16(flowTemperature)

	heatpumpModbus := samsung_mimb19n.NewClient("COM6", 500)

	heatpumpModbus.ModbusClient.SetUnitId(1)

	heatpumpModbus.CentralHeatingEnable(true)
	heatpumpModbus.SetInsideTargetTemperature(targetTemperatureUint16)
	heatpumpModbus.SetFlowTemperature(flowTemperatureUint16)

	fmt.Printf("Flow Temperature: %.2f\n", *flowTemperature)
	fmt.Printf("Target Temperature: %.2f\n", *targetTemperature)

}

func float64PtrToUint16(ptr *float64) (uint16, error) {
	if ptr == nil {
		return 0, fmt.Errorf("nil pointer")
	}

	// Ensure the float64 value is within the range of uint16
	if *ptr < 0 || *ptr > math.MaxUint16 {
		return 0, fmt.Errorf("value out of range for uint16: %f", *ptr)
	}

	return uint16(*ptr), nil
}
