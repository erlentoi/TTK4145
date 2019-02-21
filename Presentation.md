Presentation
==========

Based off of https://github.com/aalexjo/TTK4145-RTP design, following the Design evaluation sheet on Blackboard: https://ntnu.blackboard.com/bbcswebdav/pid-586132-dt-content-rid-19462142_1/courses/194_TTK4145_1_2019_V_1/designevaluation.pdf


## 1. What must be handled:
-Button presses
-Evaluating stopping on floor when a floor is reached
-Distributing all the state to all other elevators(when and how often does this happen?) (peer-to-peer)
-Calculate which action should be taken (cost module) every event(button pressed,floor reached, etc)
-Handling the connected elevators when one or more elevators loses connection to the network (Permitted assumption: at least 1 remain at all times, see design spesification), including having a elevator connect to the network
-Handling the elevator that loses connection to the network, including reconnecting the elevators to the network
-Packet loss (We will be using UDP which has no built in reliable data transf., must make our own (heartbeat?))
-software crash/loss of power
-..

### 2.











