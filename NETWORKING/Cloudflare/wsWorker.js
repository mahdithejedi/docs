// Original source: https://gist.github.com/mkornatz/d7daca0203260340ffff7e85399a48db
// We support the GET, POST, HEAD, and OPTIONS methods from any origin,
// and allow any header on requests. These headers must be present
// on all responses to all CORS preflight requests. In practice, this means
// all responses to OPTIONS requests.
const corsHeaders = {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Methods": "GET,HEAD,POST,OPTIONS",
    "Access-Control-Max-Age": "86400",
}

// The URL for the remote third party API you want to fetch from
// but does not implement CORS
const API_URL = "https://SERVICE_NAME.datahub.figment.io"
const WS_API_URL = "wss://slm.everyonehello.site"

// List all of the domains here that you want to be able to access this proxy
const ALLOW_ALL = true
const ALLOW_HTTP = true
const ALLOWED_DOMAINS = [
    'my.example.com'
]

/**
 * Receives a HTTP request, proxies the request, and returns the response. If the request is a websocket requests,
 * it hands the request off to a separate handler for creating a websocket proxy.
 * @param {Request} request
 * @returns {Promise<Response>}
 */
async function handleRequest(request) {
    const { url, headers } = request
    const { host, pathname } = new URL(url)
    const dhURL = API_URL + pathname
    const request_origin = headers.get("Origin")

    const origin = request_origin ? new URL(request_origin) : { host: headers.get('Host') }

    if (ALLOWED_DOMAINS.includes(origin.host) || ALLOW_ALL) {
        let response

        // Websocket requests are identified with an "Upgrade:websocket" HTTP header
        const upgradeHeader = request.headers.get("Upgrade")
        if (upgradeHeader && upgradeHeader === "websocket") {
            const dataHubWebsocketURL = WS_API_URL + pathname
            response = await handleWebsocketRequest(dataHubWebsocketURL)
        } else if (ALLOW_HTTP) {
            dataHubRequest = new Request(dhURL, request)
            dataHubRequest.headers.set("Authorization", API_AUTH_KEY);
            dataHubRequest.headers.set("Origin", new URL(dhURL).origin)
            response = await fetch(dataHubRequest)

            // Recreate the response so we can modify the headers
            response = new Response(response.body, response)
        } else {
            return new Response("HTTP NOT ALLOWED", { status: 405 })
        }

        // Set CORS headers
        response.headers.set("Access-Control-Allow-Origin", headers.get("Origin"))

        // Append to/Add Vary header so browser will cache response correctly
        response.headers.append("Vary", "Origin")

        return response
    }
    else {
        return new Response("Not Found for " + host, { status: 404 })
    }
}

/**
 * Receives a HTTP request and replies with a websocket proxy
 * @param {Request} request
 * @returns {Promise<Response>}
 */
async function handleWebsocketRequest(dataHubWebsocketURL) {
    // Establish the websocket connection to DataHub
    const dataHubResponse = await fetch(dataHubWebsocketURL, { headers: { "Upgrade": "websocket" } })
    if (dataHubResponse.status !== 101) {
        return new Response(null, {
            status: dataHubResponse.status,
            statusText: dataHubResponse.statusText
        })
    }
    const dataHubSocket = dataHubResponse.webSocket
    dataHubSocket.accept()

    // Create a client/server to act as the proxy layer
    const proxyWebSocketPair = new WebSocketPair()
    const [client, server] = Object.values(proxyWebSocketPair)

    // tell the Workers runtime that it should listen for WebSocket data and keep the connection open with client
    server.accept()

    // Any messages from the client are forwarded to the DataHub socket
    server.addEventListener("message", event => {
        dataHubSocket.send(event.data)
    })

    // Any messages coming from DataHub are forwarded back to the client
    dataHubSocket.addEventListener("message", event => {
        server.send(event.data)
    })

    const response = new Response(null, {
        status: 101,
        webSocket: client
    })

    return response
}

/**
 * Responds with an uncaught error.
 * @param {Error} error
 * @returns {Response}
 */
function handleError(error) {
    console.error('Uncaught error:', error)

    const { stack } = error
    return new Response(stack || error, {
        status: 500,
        headers: {
            'Content-Type': 'text/plain;charset=UTF-8'
        }
    })
}

function handleOptions(request) {
    // Make sure the necessary headers are present
    // for this to be a valid pre-flight request
    let headers = request.headers;
    if (
        headers.get("Origin") !== null &&
        headers.get("Access-Control-Request-Method") !== null &&
        headers.get("Access-Control-Request-Headers") !== null
    ) {
        // Handle CORS pre-flight request.
        // If you want to check or reject the requested method + headers
        // you can do that here.
        let respHeaders = {
            ...corsHeaders,
            // Allow all future content Request headers to go back to browser
            // such as Authorization (Bearer) or X-Client-Name-Version
            "Access-Control-Allow-Headers": request.headers.get("Access-Control-Request-Headers"),
        }

        return new Response(null, {
            headers: respHeaders,
        })
    }
    else {
        // Handle standard OPTIONS request.
        // If you want to allow other HTTP Methods, you can do that here.
        return new Response(null, {
            headers: {
                Allow: "GET, HEAD, POST, OPTIONS",
            },
        })
    }
}

addEventListener("fetch", event => {
    const request = event.request
    if (request.method === "POST" ||
        request.method === "GET"
    ) {
        event.responseWith(handleRequest(request))
    }
}


//     if (request.method === "OPTIONS") {
//       // Handle CORS preflight requests
//       event.respondWith(handleOptions(request))
//     }
//     else if(
//       request.method === "GET" ||
//       request.method === "HEAD" ||
//       request.method === "POST"
//     ){
//       // Handle requests to the API server
//       event.respondWith(handleRequest(request))
//     }
//     else {
//       event.respondWith(
//         new Response(null, {
//           status: 405,
//           statusText: "Method Not Allowed",
//         }),
//       )
//     }
//   })