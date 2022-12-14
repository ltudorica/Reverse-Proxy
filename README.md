# Reverse-Proxy
Reverse Proxy - Golang

**About the project**

This project is a reverse proxy that sits between multiple clients and one or serveral instances of an upstream service.

The reverse proxy supports multiple upstream services with multiple instances (these can be updated in config.yaml). It listens on HTTP requests and forwards them to one of the instances of an upstream service that will process the requests.

Requests are load-balanced between the instances of an upstream service and it supports RANDOM or ROUND ROBIN load-balancing strategies (lbPolicy in config.yaml). After processing the request, the upstream service will respond with the HTTP response back to the reverse proxy. The reverse proxy forwards the response back to the client (downstream) making the request.

## Getting started

### Layout
```tree
├───config
└───pkg
    ├───lbPolicy
    ├───parsing
    └───server
        ├───routes
        └───utils
```

### Prerequisites
* go
  ```sh
  https://go.dev/doc/install
  ```
### Build
## Using an IDE
1. Clone the repo
   ```sh
   git clone https://github.com/ltudorica/Reverse-Proxy.git 
   ```
2. Update config.yaml for the desired load balancing strategy - RANDOM or ROUND_ROBIN
3. Run the app
   ```sh
   go run main.go
   ```
   
## Using docker
1. Clone the repo
   ```sh
   git clone https://github.com/ltudorica/Reverse-Proxy.git 
   ```
2. Update config.yaml for the desired load balancing strategy- RANDOM or ROUND_ROBIN
3. Build the image
   ```sh
   docker build --tag reverse-proxy-app .
   ```
4. Run the image
```sh
docker run reverse-proxy-app
```
-> Application will start on port 8080

### Test the app
We can start 3 local services on ports 8081, 8082, 8083 according the config.yaml file. 
Python must be installed (https://www.python.org/downloads/) - use the following command for starting the upstream services for test:
```sh
py -m http.server --bind localhost <port>  
```

The file can be also modified with other services in use

Send more requests to the reverse-proxy and these will be redirected depending on availability & load balancing strategy
