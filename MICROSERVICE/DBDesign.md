#HeadLines

[DB management](https://relevant.software/blog/microservices-database-management/)

* [Saga Pattern](https://microservices.io/patterns/data/saga.html): [Microsoft Documentation](https://docs.microsoft.com/en-us/azure/architecture/reference-architectures/saga/saga), [Orchestration vs Choreography](https://www.accionlabs.com/microservices-orchestration-vs-choreography-what-to-prefer)
* [DB Per service](#DB_Per_service)

## [DB Per service](https://microservices.io/patterns/data/database-per-service.html)
### Problem
* Services must be loosely coupled so that they can be developed, deployed and scaled independently
* Some business transactions need to query data that is owned by multiple service
### Solution
**Keep each microservice’s persistent data private to that service and accessible only via its API. A service’s transactions only involve its database.**

***you aren’t required to have a separate database server for each service. Instead, multiple services can share the same database server with a logical separation of their data.***

