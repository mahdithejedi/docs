## Asynchronous operations
Sometimes a POST, PUT, PATCH, or DELETE operation might require processing that takes a while to complete
<br />
If so, consider making the operation asynchronous. Return HTTP status code 202 (Accepted) to indicate the request was accepted for processing but is not completed.
<br />
You should expose an endpoint that returns the status of an asynchronous request
<br />
```
HTTP/1.1 202 Accepted
Location: /api/status/12345
```
OR
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "status":"In progress",
    "link": { "rel":"cancel", "method":"delete", "href":"/api/status/12345" }
}
```
If the asynchronous operation creates a new resource, the status endpoint should return status code 303 (See Other) after the operation completes. In the 303 response, include a Location header that gives the URI of the new resource:
```
HTTP/1.1 303 See Other
Location: /api/orders/12345
```

## Filter and paginate data
GET requests over collection resources can potentially return a large number of items. You should design a web API to limit the amount of data returned by any single request.
<br />
***Also consider imposing an upper limit on the number of items returned, to help prevent Denial of Service attacks***
<br />
You can extend this approach to limit the fields returned for each item, if each item contains a large amount of data

# API Versioning
* No Versioning
* URI versioning
* Query string versioning
* Header versioning

### Resources
[dissertion Representational State Transfer (REST)](https://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm)
<br />
[Microsoft API Design](https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design)

