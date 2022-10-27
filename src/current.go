package main

import (
	"machine"
)

func CurrentSensorValue(pin machine.Pin) {
	sensorValue := machine.ADC{pin}.Get()
	if sensorValue < 500 {
		
		print(LED_BROKEN) 
	}
}
