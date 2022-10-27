compile:
	tinygo build -target bluepill -o flash.bin *.go 

flash: 
	stm32flash -w ./flash.bin /dev/tty.usbserial-0001

clean: 
	rm -f flash.bin

