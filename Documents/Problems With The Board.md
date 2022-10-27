# LoRa-E5 dev board 

Made: 2022-10-07

by: Group 5 

## Major problems with the board

### 1. We can not flash the board

It means that we can not push code in it, so by extension we can not interact with the board. 

### 2.We can not connect it to the LoRa network

We can not connect the board to the LoRa network, 

Since the board has only one antenna capable of connecting to the network from 10km away, 

And the nearest one is at bourges (about 40km away),

## Solutions 

We need both of these solutions in order to complete this project.

### 1. St-link + Gateway

It seems that the st-link is the only thing that can flash the board. 

It worked for the group 8, because they bought a st-link.

We can install a Gateway near the school to be able to connect to the network. 

If we do not we will not be able to try the features.


| pros      | cons |
| ----------- | ----------- |
| We can continue with the same hardware,       | expensive       |
| An st-link is not as expensive as an hypotetical new board   | take times to install         |

[Link of the ST-link](https://www.amazon.fr/gp/product/B07RKVM7H8/ref=ppx_yo_dt_b_asin_title_o00_s00?ie=UTF8&psc=1)

[Link of the gateway](https://www.atim.com/produits/gateway-lorawan-outdoor/)


[Link of another gateway](https://login.resiot.io/store/resiot-indoor-lorawan-gateways?)

### 2. Change of board 

We can change of board, but we will have to start from scratch, it will probably be the most expensive solution.

And we will be able to pass by another network than LoRa.


| pros      | cons |
| ----------- | ----------- |
| Clock built-in to have a more accurate device      | expensive       |
| Paragraph   | possibly useless        |
| Paragraph   |         |

### 3. Change of network


We can change the whole network, we could use Sigfox or the wifi network, 

area covered by sigfox : 
![SigFox]("./images/../../France%20Coverage%20sigfox.png")

Vierzon area covered by sigfox :

![Vierzon]("./images/../../Vierzon%20Coverage%20sigfox.png)

| pros      | cons |
| ----------- | ----------- |
| We can continue with the same hardware,      | need to restart from scratch       |
| French network   | not as documented as LoRa        |
| Area covered throughout the world  |         |

### Solution bis

All the project class can take place in Bourges.

## Minor problems with the board

### 1. We do not have an internal clock 

We are limited by this problem, because we do not have an internal clock.

## Problems with the CP2102 module (USB to TTL)

The USB to TTL is built-in the board, so we can not use it.

And it was not the right tool for the job.

