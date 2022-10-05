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

- The problem is that luminous signage is consuming a lot of energy and in case of a issue, it is hard to know where the problem is. It also takes time to fix the problem. The solution is to create a system that will allow to monitor luminous signage and to diagnose the problem in order to fix it in a fast way.

## 1.b. Glossary

### TinyGo

    TinyGo is a Go compiler for small places. It compiles Go programs to WebAssembly, microcontrollers, and command-line tools.

### LoRa-E5 Developpement board

    LoRa-E5 Dev board is a LoRaWAN development board based on the Semtech SX1262 LoRa transceiver and the STM32WL55JC microcontroller. It will have our code in it and will be used for the communication between the luminous signage and the things network.


### The Things Network

    The Things Network is a global community building an open-source LoRaWAN network. It is a non-profit organization that provides a global, decentralized, open-source LoRaWAN network.

### LoRaWAN

    LoRaWAN is a media access control (MAC) protocol for wide area networks. It is designed to allow low-power communication between battery-powered devices over long range wireless connections.

## 1.c. Context

- The problem is worth solving because energy consumtion is a big problem in the world. It is also a problem for the environment. It also make it faster to fix the problem resulting in a better reputation for the company.

## 1.d. Goals

- As a user, I want to be able to monitor the luminous signage in order to know if it is working or not.

- As a user, I want to be able to switch on and off the luminous signage in order to save energy.

- As a user, I want to be able to change the brightness of the luminous signage in order to save energy.

- As a user, I want to be able to know if the luminous signage is not behaving as intended in order to prevent damage.

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

- This solution is not very efficient because it is not very fast and it is not very cheap. It also requires human interaction.

## 2.b. Proposed solution

- The proposed solution is to create a system that will allow to monitor luminous signage and to diagnose the problem in order to fix it in a fast way.

## 2.c. Test plan

- Every function is being tested in a separate test file in go. It allows us to test each function individually and to make sure that it works as intended.

- We will sumulate user tests in order to be as accurate as posible on to what the user wants.

- The integration test will be done on the LoRa-E5 Dev board once we have what we need to flash the code into it...

# 3. Further Considerations

## 3.a. Third party services

- APIs

    - Electricity disponibility found here : 
        - https://www.electricitymap.org/
        - https://portal.realto.io/ 
    
    the problem is that if one of the API is down, the electricity saving system will be down.

- LoRaWAN

    If the network is down, unnaccessible in the city or even deprecated, the system will not work anymore.

- The things network

    If the support of the things network stops, we will have to find another way of communication between the luminous signage and the client.

## 3.b. Cost analysis

### 3.b.1. Cost per month

- 1 AA batteries cost around 1.5€.
Depending on how long the batterie lasts, the price may vary.

- In the worst case scenario, the batteries only lasts six month, that means that the cost per month is 0.50€. 

### 3.b.2. Cost for deployment

The cost for deployment is :

- 1 LoRa-E5 Development Kit : 25.00€

- 2 AA batterie : 3.00€

Which gives a total of 28.00€.

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

- He will also have to check if the electricity disponibility API is available in the country. 

- He will have to check what are the laws regarding the luminous signage.

### 3.c.5. Risks

- The communication with the API is a risk.
    
    The support for the API can stop or require a payment and in this case, one part of the system will not be able to function properly.

    However, taking this risk can help the client to mitigate its carbon footprint.

# 4. Work Breakdown Structure

- We follow a strict timeline using trello as a support tool. We have a lot of tasks to do and we have to be sure that we are on time.

# 5 Conclusion

- The problem is worth solving because luminous signage is consuming a lot of energy and in case of an issue, it is hard to know where the problem is. It also takes time to fix the problem. The solution will allow to monitor luminous signage and to diagnose the problem in order to fix it in a fast way.

- This solution will also help reduce the carbon footprint of luminous signage.

# 6. References

- https://www.electricitymap.org/

- https://portal.realto.io/

- https://www.thethingsnetwork.org/

- https://github.com/github/copilot-docs/
