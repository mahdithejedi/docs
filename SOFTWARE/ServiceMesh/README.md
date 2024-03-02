# What is a service mesh?

A service mesh is a networking framework  for adding observability, security, and reliability to distributed applications. It achieves this by providing these functions at the platform layer, instead of having them embedded in the application layer.

Technically, a service mesh is a set of lightweight network proxies, typically deployed alongside application code in a “sidecar container”. These proxies form the data plane of the service mesh, which is controlled and configured by a control plane. Proxies perform several important functions including:

* Handling communication between microservices
* Service discovery
* Load balancing
* Authentication and authorization
* Observability