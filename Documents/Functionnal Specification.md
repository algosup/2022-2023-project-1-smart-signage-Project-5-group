
# Appsolu

# Functionnal specification

Malo Archimbaud
Last update : 24/10/22

# Summary

- [Overview](#overview)
- [Personan](#personas)
- [Scenarios](#scenarios)
- [Out of scope](#out-of-scope)
- [Description of the functionalities](#description-of-the-functionalities)
- [Bug management](#bug-management)
- [Risks](#risks)
- [Glossary](#glossary)


# Overview

SignAll is a company which manufacture luminous signage for companies. They want us to create a new device named Appsolu. This device will make their signage smart by being connected.

This functional specification is not the final version.

# Personas

**Pablo**

Pablo is the manager of a small shop of clothes in Vierzon. He is 48 years old. He lives in Vierzon with his wife and his 4 children. He also have a dog. He likes to go fishing and try passing on his passion to his children. 
He is running the shop that his father passed on to him. He is working here since he is 18 and his wife joined him 10 years later. He is a client of signall for 10 years because he had to changed his old signage and wanted to work with a company close to him. 
He wants to lower his energy consumption and he also wants to be able to do preventive maintenance rather than waiting for his signage to be broken because he has no way to know if something is wrong before it is too late. He also wants to be sure that he is not breaking any law, since a there will be a law that states that the signage has to be switched off between 1 and. 

**Pedro** 

Pedro is the CEO of Signall. He is 43 years old. He lives in Vierzon with his wife and his 3 children. He likes playing golf when he haas some free time.
He is the CEO since 2016 and since then, the company is running very well.
Currently, he has to spend money on having external people working with Signall. He wants to spend less money to run his company and he thinks that he could sell his own maintenance service. That will allow him more control on his product he could also earn more money because he will not have to pay another company and creating a margin to reinvest in the company

**Greta**

Greta is a citizen. She is 22 years old. She lives in Vierzon. She is single and she has a cat at home. She is an ecological activist. She participated in a lot of demonstration et pacific walk in favor of the planet and the ecology in general. She showers once a week and buy her food in local store that are in lines with her principles.
She wants to help the planet. She already switch off the signage that she found when she hangs out in town at 3 AM but she would like that all those useless signs are automatically switched off.

# Scenarios

**Scenario 1**

Pablo wants to have some preventive maintenance on his signage. Thanks to Appsolu, he will be able know the condition of his signage and ash for an intervention before it's too late. Also Concerned by the rise of the electricty price and the possible shortage incoming, he wants to be sure that in case of a shortage during this winter, his signage will lower its consumption or shut itself down because she could have troubles with the autorities. Appsolu will include a functionality to ensure that the signage will respect the laws.

**Scenario 2**

Pedro wants to be sure that his own technician will take action properly on the signage of his client. Appsolu will allow him to know what is wrongn and send the right technician at the right place.

**Scenario 3**

Greta is pissed of that signages can stay on at night. The signage that will have Appsolu will be able to be programate to be switche off at nighttime and then save energy to help planet.

# Out of scope

This version will **not** support this feature:

 - A user-friendly interface

# Description of the functionalities

## LEDs monitoring

It will allow the user to check if the LEDs are on, off or if they are broken. The user will also be able to check if LEDs are overheating. 

## Signage monitoring

It will allow the user to switch the signage on and off remotely. The user will also be able to change the brightness of the signage. He will be able to choose via a scrolling bar.

## Compliance with laws

Thanks to this functionality, the user will be sure that the usage of the signage will be in phase with environmental and ecological laws. The user will have to tell his schedule to be sure that the signage is not shutting himself down when the store open.

## Auto brightness

This feature will ensure that the brightness of the signage will automatically adapt in function of the ambiant luminosity. For example, when it will be very sunny outside, the brightness will be at the maximum level. If it is cloudy outside, the brightness will be at 60% and if it is nighttime the brightness will be at 20%.

## Money savings

When the price of energy will increase, the signage will be able to reduce its energy consumption to ensure that the user is not overpaying.

## Schedule programmation

The user will be able to give his schedule to be sure that the signage are switched on only when his store his open. For example, he will be able to program that he is closed the Monday and open from Tuesday to Saturday from 10 AM to 7 PM.

## Data encryption

The data will be sent in an encrypted way to ensure that only the user who is abilited to receive the data will receive it and that no one else will be able to use the data.

# Bug management

Every detected bug will be saved in a database with who found it and when. The person finding the bug will also have to provide a detailled description on how to reproduce it to ensure that the software engineer will be able to fix it. When fixed, the software engineer will have to precise when and how he fixed it.

Example : 

| N°Issue | Issue | Username | Date of report | Date of fix | How to fix |
| --------| ----- | -------- | -------------- | ----------- | ---------- |
|         |       |          |                |             |            |


# Risks

## Hardware

If the board stop working or breaks, the device will not be able to send data to the user. It will be a problem for the user because he will not be able to check if the signage is working or not.

## Software

If the software stop working due to a crash for example, the device will not be able to send data to the user. It will be a problem for the user because he will not be able to check if the signage is working or not.

## Unsatisfaction

The client could be unsatisfied with our product because it could not fullfill client requirement.

## Delay

The project could be delayed because of an hardware problem which could prevent us to keep a good pace.

# Glossary

LED : an LED is an electric component which create light when a current goes through it. 

Signage : a signage is all LEDs together which shapes the logo that the client needs.