package main

func main() {

}

func sendMessage(message string) {
	// Send the AT command AT+MSG + message
	uart.Write([]byte("AT+MSG" + message + "\r\n"))
}