About the project

This project is a reverse proxy that sits between multiple clients and one or serveral instances of an upstream service.
The reverse proxy supports multiple upstream services with multiple instances (these can be updated in config.yaml). It listens on HTTP requests and forwards them to one of the instances of an upstream service that will process the requests.
Requests are load-balanced between the instances of an upstream service and it supports RANDOM or ROUND ROBIN load-balancing strategies (lbPolicy in config.yaml). After processing the request, the upstream service will respond with the HTTP response back to the reverse proxy. The reverse proxy forwards the response back to the client (downstream) making the request.

Quick start
Build
