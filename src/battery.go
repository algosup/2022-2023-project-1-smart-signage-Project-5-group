package main

import (
	"machine"
)

func HandleBattery() {
	// Configure the ADC
	machine.ADC{machine.ADC0}.Configure(machine.ADCConfig{})

	// Read the value
	value := machine.ADC{machine.ADC0}.Get()

	// Print the value to the serial monitor
	if value <= 10 {
		print(LOW_BATTERY)
	}
}

