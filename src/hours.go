package main

import (
	"machine"
	"time"
)


// This function will automatically shutdown the board between 1.00am and 5.00am 
func LowConsumption(a int, b int) {
	for {
		if time.Now().Hour() >= a && time.Now().Hour() <= b {
			uart.Configure(machine.UARTConfig{TX: machine.UART_TX_PIN, RX: machine.UART_RX_PIN})
			uart.Write([]byte("AT+SLEEP\r\n"))
		}
		if time.Now().Hour() >= b && time.Now().Hour() <= a {
			uart.Configure(machine.UARTConfig{TX: machine.UART_TX_PIN, RX: machine.UART_RX_PIN})
			uart.Write([]byte("AT+ALARM\r\n"))
		}
	}
}


