package main

import (
	"machine"
)

const (
	FAST = 800
	MID  = 450
	SLOW = 100
	A4 = machine.PF4
	ADC4 = machine.PF4
)

var uart = machine.UART2


func HandleAmbientLights() {
	// configure the LED
	uart.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.A2, RX: machine.A3})

	uart.Write([]byte("AT+JOIN\r\n"))

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// configure the photoresistor
	photoresistor := machine.ADC{ADC4}
	photoresistor.Configure( machine.ADCConfig{})

	// read the photoresistor value
	for {
		// read the value
		value := photoresistor.Get()

		// if the value is greater than 1000, turn the LED on
		if value > FAST {
			// send an AT command to the LoRa module to shutdown the led 
			// uart.Write([]byte("AT+MSG""\r\n"))
		} else if value > MID {
			// uart.Write([]byte("AT+MSG""\r\n"))
		} else {
			// uart.Write([]byte("AT+MSG""\r\n"))
		}
	}
}

