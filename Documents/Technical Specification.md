# Technical specification

## Appsolu

### Author(s)
Clément Caton

## Team
Théo Diancourt,
Clémentine Curel,
Malo Archimbaud,
Clément Caton

## Created on
01-10-2022

## Last updated
05-10-2022

## Summary

- [1. Introduction](#1-introduction)
- [2. Solutions](#2-solutions)
- [3. Further Considerations](#3-further-considerations)
- [4. Work Breakdown Structure](#4-work-breakdown-structure)
- [5. Conclusion](#5-conclusion)
- [6. References](#6-references)


# 1. Introduction

## 1.a. Overview

- The problem is that luminous signage is consuming a lot of energy and in case of a issue, it is hard to know where the problem is. It also takes time to fix the problem. The solution is to create a system that will allow to monitor luminous signage, to diagnose the problem in order to fix it in a fast way and to reduce the energy consumption.

## 1.b. Glossary

### TinyGo

- TinyGo is a Go compiler for small places. It compiles Go programs to WebAssembly, microcontrollers, and command-line tools.

### LoRa-E5 Developpement board

- LoRa-E5 Dev board is a LoRaWAN development board based on the Semtech SX1262 LoRa transceiver and the STM32WL55JC microcontroller. It will have our code in it and will be used for the communication between the luminous signage and the things network.


### The Things Network

- The Things Network is a global community building an open-source LoRaWAN network. It is a non-profit organization that provides a global, decentralized, open-source LoRaWAN network.

### LoRaWAN

- LoRaWAN is a media access control (MAC) protocol for wide area networks. It is designed to allow low-power communication between battery-powered devices over long range wireless connections.

### 7 Segment display

- A 7-segment display is a form of electronic display device for displaying decimal numerals that is an alternative to the more complex dot-matrix displays. It is a form of multiplexing, where a single 7-segment LED display is used to display all of the individual digits, which creates the illusion that all of the digits are being displayed simultaneously.

### Arduino daughter board

- The Arduino we are using is an arduino 33 BLE. It is controlled by another card, the LoRa-e5 board, making it a daughter board. The use of daughter boards is to delegate the tasks to other boards, making the main board lighter and more efficient.

### I2C

- I2C is a synchronous, multi-master, multi-slave, packet switched, single-ended, serial computer bus invented by Philips Semiconductor (now NXP Semiconductors). It is commonly used for attaching lower-speed peripheral ICs to processors and microcontrollers in short-distance, intra-board communication.

## 1.c. Context

- The problem is worth solving because energy consumtion is a big problem in the world. It is also a problem for the environment. It also make it faster to fix the problem resulting in a better reputation for the company.



## 1.d. Goals

- As a user, I want to be able to monitor the luminous signage in order to know if it is working or not.

    A module monitoring the voltage going through the transformator will be used to know if the luminous signage is working or not.

    The function will take the voltage as an input and will return a boolean. If the voltage is above a certain threshold, the function will return true, if it is below, it will return false.

- As a user, I want to be able to monitor the temperature of the hardware in order to prevent its degradation.

    A module to monitor the temperature will be used to calculate if the temperature is too high. 

    The function should take the temperature as an input and return a boolean, true if the temperature is normal and false if it is too high.

- As a user, I want to be able to switch on and off the luminous signage in order to save energy.

    With a certain LoRaWAN signal, The LoRa-E5 board will say to the daughter board if the signage should be on or off.

    The arduino daughter board will wait and listen to one of it's digital pins and if it receives a signal in this pin, it will switch on or off the signage.

- As a user, I want to be able to change the brightness of the luminous signage in order to save energy.

    With a certain LoRaWAN signal, The LoRa-E5 board will say to the daughter board what the brightness should be.

    The arduino daughter board will wait and listen to one of it's digital pins and if it receives a signal in this pin, it will read an analog pin and change the brightness of the signage in accordance with the value of the analog pin.

- As a user, I want to be able to know if the luminous signage is not behaving as intended in order to prevent damage.

    A sensor monitoring the voltage going through the transformator will be used to know if the luminous signage is working or not.

    The sensor used has an accuracy of 0.2% which allow us to know if the current drop is normal or not in a line having up to 250 Leds. If a sign is having more than 250 Leds, we will have to wait for 2 Leds to be off to know if the current drop is normal or not.

    If the voltage drops by at least 1/250th of it's usual lowest value, the function will return true, if it is not the case, it will return false.

- As a user, I want to be able to create a schedule for the luminous signage in order to do all of the above automatically.

    With a certain LoRaWAN signal, The LoRa-E5 board will re program the arduino daughter board with a new schedule. Inside the schedule, there will be the time when the signage should be on or off and the brightness.

    I don't know how to do it with the pins we have at our disposal but I think it is possible. For example using the I2C protocol.

- As a user I want the signage to automatically adapt it's brightness to the outside light in order to save energy.

    A module to monitor the light will be used to calculate the brightness of the signage.

    The function should take the light as an input and change the brightness of the signage.

- As a user, I want to be able to see the error code on the device with a press of a button in case I forgot what were the errors.

    In the code and the manual, error codes will be constant integer and will be displayed using a 7 segment display. The display won't turn on except if the user presses a button. This will save energy.

    The function controlling the 7 segment display will take the error code as an input and display it on the 7 segment display. 
    ```go
    0:  "No led"
    1:  "Led is broken"
    2:  "No Power"
    3:  "OverHeating"
    4:  "OverCooling"
    5:  "Low Battery"
    6:  "Sensor is broken"
    7:  "Device not connected"
    8:  "LoRa is not responding"
    ```

## 1.e. Out of scope

- A user-friendly interface

# 2. Solutions

## 2.a. Current solution

- The current solution is to send a technician to the luminous signage to check if it is working or not. If it is not working, the technician will have to fix it. It is a long process and it is not very efficient.

|pros               |cons|
|----               |----|
|It is a solution   |It is not very efficient                 |
|                   |It is not very fast                      |
|                   |It is not very cheap                     |
|                   |Requires human interaction               |

## 2.b. Proposed solution

- The proposed solution is to create a system that will allow to monitor luminous signage and to diagnose the problem in order to fix it in a fast way.

- It will also allow the user to setup a schedule for the luminous signage in order to save energy.

## 2.c. Test plan

- Every function is being tested in a separate test file in go. It allows us to test each function individually and to make sure that it works as intended.

- We will simulate user tests in order to be as accurate as posible on to what the user wants.

- The integration test will be done on the LoRa-E5 Dev board once we have what we need to flash the code into it...

# 3. Further Considerations

## 3.a. Third party services

- LoRaWAN

    If the network is down, unnaccessible in the city or even deprecated, the system will not work anymore.

- The things network

    If the support of the things network stops, we will have to find another way of communication between the luminous signage and the client.

## 3.b. Cost analysis

### 3.b.1. Cost per month

- 1 AA batteries cost around 1.5€.
Depending on how long the batterie lasts, the price may vary.

- In the worst case scenario, the batteries only lasts six month, that means that the cost per month is 0.50€. 

### 3.b.2. One time cost

To have a functional project we need :

- 1 LoRaWAN Gateway : 300.00€

- 1 ST-Link : 20.00€

- 1 LoRa-E5 Development Kit : 25.00€

- 1 USB to TTL : 5.00€

- Sensors : 5.00€

- 1 Blue Pill : 2.00€

For the first deployment it will cost 52€.

For all other deployment it will cost 32€. 

If a LoRaWAN Gateway is needed, it will cost an additional 300€.

## 3.c. Security

### 3.c.1. Security threats

- The luminous signage can be hacked and the brightness can be changed.

- The lighting schedule can be changed.

- False positive alerts can be sent to the client.

- The attacker can turn on the luminous signage during illegal hours.

### 3.c.2. Security measures

- Communication between the luminous signage and the client is encrypted.

### 3.c.3. Privacy considerations

- The Solution is not collecting any personal data.

### 3.c.4. Regional consideration

- The LoRaWAN network is not available in every city so he will have to check if said city is covered by the network.

- If the user wants to extend the system to another country, he will have to check if the LoRaWAN network is available in the country.

- He will have to check what are the laws regarding the luminous signage.

### 3.c.5. Risks

- Getting not suited hardware.

- Delivering the project late.

- Having the connection not soldered properly resulting in having the hardware not working.

# 4. Work Breakdown Structure

- We follow a strict timeline using trello as a support tool. We have a lot of tasks to do and we have to be sure that we are on time.

# 5 Conclusion

- The problem is worth solving because luminous signage is consuming a lot of energy and in case of an issue, it is hard to know where the problem is. It also takes time to fix the problem. The solution will allow to monitor luminous signage and to diagnose the problem in order to fix it in a fast way.

- This solution will also help reduce the carbon footprint of luminous signage.

# 6. References

- https://www.electricitymap.org/

- https://www.thethingsnetwork.org/

- https://github.com/github/copilot-docs/
