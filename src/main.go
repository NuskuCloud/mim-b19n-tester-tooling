package main

import (
	"flag"
	"fmt"
	"github.com/nuskucloud/samsung_mimb19n"
	"math"
	"os"
	"strconv"
	"time"
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

	heatpumpModbus := samsung_mimb19n.NewClient("COM6", 500*time.Millisecond)

	err := heatpumpModbus.ModbusClient.SetUnitId(1)
	if err != nil {
		fmt.Println("Error setting unit id")
		return
	}

	err = heatpumpModbus.ModbusClient.Open()
	if err != nil {
		// error out if we failed to connect/open the device
		// note: multiple Open() attempts can be made on the same client until
		// the connection succeeds (i.e. err == nil), calling the constructor again
		// is unnecessary.
		// likewise, a client can be opened and closed as many times as needed.
		fmt.Println("Error opening modbus client")
		panic(err)
		return
	}

	heatpumpModbus.ModbusClient.WriteRegisters(6000, []uint16{HexToUint16("8238"), HexToUint16("8204")})
	if err != nil {
		fmt.Println("Error enabling hidden registers")
		panic(err)
	}

	err = heatpumpModbus.CentralHeatingEnable(true)
	if err != nil {
		fmt.Println("Error enabling central heating")
		panic(err)
		return
	}

	// Set it to heating mode
	err = heatpumpModbus.ModbusClient.WriteRegister(53, 4)
	if err != nil {
		fmt.Println("Could not set register 53 to value 4 (heating mode")
		panic(err)
		return
	}

	err = heatpumpModbus.SetInsideTargetTemperature(targetTemperatureUint16)
	if err != nil {
		fmt.Println("Error setting target temperature")
		panic(err)
		return
	}

	err = heatpumpModbus.SetFlowTemperature(flowTemperatureUint16)
	if err != nil {
		fmt.Println("Error setting flow temperature")
		panic(err)
		return
	}

	fmt.Println("Successfully set temperatures I think, with the following values:")
	fmt.Printf("Flow Temperature: %.2f\n", *flowTemperature)
	fmt.Printf("Target Temperature: %.2f\n", *targetTemperature)

	temperature, err := heatpumpModbus.ReadIndoorTemperature()
	if err != nil {
		fmt.Println("Error reading indoor temperature")
		panic(err)
		return
	}
	fmt.Println("modbus indoor temperature:")
	fmt.Println(temperature)

	temperature, err = heatpumpModbus.ReadOutdoorTemperature()
	if err != nil {
		fmt.Println("Error reading outdoor temperature")
		panic(err)
		return
	}
	fmt.Println("modbus outdoor temperature:")
	fmt.Println(temperature)

	//temperature, err = heatpumpModbus.ReadFlowTemperature()
	//if err != nil {
	//	fmt.Println("Error reading flow temperature")
	//	panic(err)
	//	return
	//}
	//fmt.Println(fmt.Sprintf("modbus flow temperature: %g", temperature))
	//
	//temperature, err = heatpumpModbus.ReadReturnTemperature()
	//if err != nil {
	//	fmt.Println("Error reading return temperature")
	//	panic(err)
	//	return
	//}
	//fmt.Println(fmt.Sprintf("modbus return temperature: %g", temperature))

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

func HexToUint16(hexStr string) uint16 {
	// Convert the hex string to a decimal integer (base 16)
	value, err := strconv.ParseUint(hexStr, 16, 16)
	if err != nil {
		panic(err)
	}
	return uint16(value)
}
