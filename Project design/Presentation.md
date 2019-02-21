Presentation
==========

Based off of https://github.com/aalexjo/TTK4145-RTP design, following the Design evaluation sheet on Blackboard: https://ntnu.blackboard.com/bbcswebdav/pid-586132-dt-content-rid-19462142_1/courses/194_TTK4145_1_2019_V_1/designevaluation.pdf


## 1. What must be handled:
- Button presses
- Evaluating stopping on floor when a floor is reached
- Distributing all the state to all other elevators(when and how often does this happen?) (peer-to-peer)
- Calculate which action should be taken (cost module) every event(button pressed,floor reached, etc)
- Handling the connected elevators when one or more elevators loses connection to the network (Permitted assumption: at least 1 remain at all times, see design spesification), including having a elevator connect to the network
- Handling the elevator that loses connection to the network, including reconnecting the elevators to the network
- Packet loss (We will be using UDP which has no built in reliable data transf., must make our own (heartbeat?))
- software crash/loss of power
- ..

### 2. Modules:
- main
- fsm
- network(provided)
  - Distribution ("Alternative 2: Reassigning all requests" ? If yes, it is provided TTK4145/Project-resources/blob/master/cost_fns )
    - Reassigning all requests: "In order for this approach to work, it is necessary that either a)(...) or b) all elevators that calculate the redistribution eventually come to the same conclusion"
- cost(provided)
- some kind of module that handles each elevators interface with the network? (like in aalexjo)

### 3. Example scenarios:
- 1. System is idle.A button gets pushed, one elevator serves the order. While the elevator is traveling, another button gets pushed.
In the elevator that responds: Driver adds the order to the queue. FSM is updated with the order and passes it along with the elevator state to the Cost module and Interface module. The Cos module, already having received the order from the running elevator, calcualtes the optimal cost. The Network module then distributes this information to all other elevators.
The elevator that is already running receives the order at the Network module, and it gets passed to its Cost module. The Cost module minimizes the cost and passes the calculated orders back to the interface which accepts the order if it agrees with the other cost calcualtions.











