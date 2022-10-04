
# Appsolu

# Functionnal specification

Malo Archimbaud
Last update : 04/10/22

# Overview

SignAll is a company which manufacture luminous signage for companies. They want us to create a new device named Appsolu. This device will make their signage smart by being connected.

This functional specification is not the final version.

# Scenarios

**Scenario 1 : Alban**

Alban is in charge of knowing if luminous signage of every restaurant of his brand in France are working. He used to send teams of technicians when the owner of one of his restaurant signaled a problem. The problem is that it could take weeks to fix a signage and it would damage the image of the brand. 
Thanks to Appsolu, he will be able to check in real time if a signage need to be fixed and send a team on it more quickly.

**Scenario 2 : Mike**

Mike is the owner of a restaurant from a famous fast-food brand. Concerned by the environnement concerns, he would like to be able to switch off the lights when his restaurant is not open and be able to check his energy consumption. It will allow him to make financial savings and have a better reputation. 
Thanks to Appsolu, he will be able to switch off and on the lights in a much easier way than before. He will also be able to check his energy consumption and reduce the power of the led to save energy.

**Scenario 3 : Mike**

The same Mike would also like that, for the sake of energy saving, the power given to the LEDs reduce when the outside luminosity is lower than usual, during night, bad weather or during an eclipse.
Thanks to Appsolu, some captors will be able to recognize a drop of luminosity and adapt the power gave to the LEDs.

# Out of scope

This version will **not** support these features:

 - A user-friendly interface

# Description of the functionalities

## LEDs monitoring

It will allow the user to check if the LEDs are on, off or if they are broken. The user will also be able to check if LEDs are overheating. 

## Signage monitoring

It will allow the user to switch the signage on and off remotely. The user will also be able to change the brightness of the signage.

## Compliance with laws

Thanks to this functionality, the user will be sure that the usage of the signage will be in phase with environmental and ecological laws.

## Auto brightness

This feature will ensure that the brightness of the signage will automatically adapt in function of the ambiant luminosity.

## Money savings

When the price of energy will increase, the signage will be able to reduce its energy consumption to ensure that the user is not overpaying.

## Battery warnings

The device will be able to send a message when the battery reach a low level. It will allow the user to change the battery before losing signal with the device.

## Schedule programmation

The user will be able to give his schedule to be sure that the signage are switched on only when his store his open.

## Data encryption

The data will be sent in an encrypted way to ensure that only the user who is abilited to receive the data will receive it and that no one else will be able to use the data.

# Bug management

Every detected bug will be saved in a database with who found it and when. The person finding the bug will also have to provide a detailled description on how to reproduce it to ensure that the software engineer will be able to fix it. When fixed, the software engineer will have to precise when and how he fixed it.
