package main 

import (
	"fmt"
	"machine"
	"time"
)


// this function will warn the user if the temperature of the device is too high or too low
func HandleTemperature(pin machine.Pin) {
	if machine.ADC{pin}.Get() > 85 {
		fmt.Println(ERRORS[OVERHEATING])
	}
	if machine.ADC{pin}.Get() < -45 {
		fmt.Println(ERRORS[OVERCOOLING])
	}
} 
