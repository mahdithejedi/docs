# Definition And Usage

Create helper services that send network requests on behalf of a consumer service or application. An ambassador service
can be thought of as an out-of-process proxy that is co-located with the client.

## Disadvantages

* The proxy adds some latency overhead
* Consider the possible impact of including generalized features in the proxy. For example, the ambassador could handle
  retries, but that might not be safe unless all operations are idempotent.
* Consider whether to use a single shared instance for all clients or an instance for each client.

## When NOT to use it?

* When network request latency is critical. A proxy introduces some overhead, although minimal, and in some cases this
  may affect the application.
* When connectivity features can't be generalized and require deeper integration with the client application.
