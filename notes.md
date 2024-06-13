# Building an API:

We need to think about:
 - Payload size
 - Latency
 - Scalability
 - Load Balancing
 - Languages Interop
 - Auth, Monitoring, Log

## 4 types of API in gRPC
 * Unary: closest to REST API
 * Server streaming: One request by Client multiple reponses by the server(client to get updating in Real time)
 * Client streaming: Multiple requests by the Client and One response by the Server (Uploading, Updating Information)
 * Bi directional streaming: Client sent multiple requests and Server response multiple responses in any order.

## Scalability
 * Server: Async
 * Client: Async or Blocking
Example: Google: 10 B requests/sec

## Security(SSL)
 * Schema based serialization
 * Easy SSL certificates initializaion
 * Interceptors for Auth

 ## gRPC vs REST
 * gRPC:
    - Protocol Buffers
    - HTTP 2
    - Streaming
    - Bi directional
    - Free design

 * REST:
    - JSON
    - HTTP 1
    - Unary
    - Client -> Server
    - GET/POST/UPDATE/DELETE
    
