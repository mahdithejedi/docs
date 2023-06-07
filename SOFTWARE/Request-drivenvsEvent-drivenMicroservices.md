## Request Driven Microservices
Benefits:
* Flow is clear

Trade-offs:
* in the Orchestrator service we have Single Point Of Failure (Orchestrator itself)
* There is no easy way to recover the actions
* Rest API of the dependent services cannot be easily modified


## Event-Driven Microservices
Benefits:
* The producer service of the events does not know about its consumer services. On the other hand, the consumers also do not necessarily know about the producer
* Much easier to add, remove or modify services
* If a service goes offline while producer process events, it can replay (rewind) those events once it came back online

Trade-offs:
* There is no clear central place (orchestrator) defining the whole flow.
* Managing distributed transaction could be complex
