package main

import (
	"flag"
	"fmt"
	"github.com/nuskucloud/samsung_mimb19n"
	"os"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	flowTemperature := flag.Float64("flow_temperature", -1000.0, "Flow temperature")
	targetTemperature := flag.Float64("target_temperature", -1000.0, "Target temperature")

	flag.Parse()

	if *flowTemperature == -1000.0 || *targetTemperature == -1000 {
		// dump
		os.Exit(1)
	}

	heatpumpModbus := samsung_mimb19n.NewClient("/dev/ttyUSB0", 500)

	heatpumpModbus.DhwEnable(true)
	heatpumpModbus.SetInsideTargetTemperature(uint16(*targetTemperature))
	heatpumpModbus.SetFlowTemperature(uint16(*flowTemperature))

	fmt.Printf("Flow Temperature: %.2f\n", *flowTemperature)
	fmt.Printf("Target Temperature: %.2f\n", *targetTemperature)

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
