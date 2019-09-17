# iofog-sdk-go

This SDK contains a set of Golang packages developers can use for the purposes of:
* Developing your own edge microservices that run on ioFog Edge Compute Networks
* Interacting with ioFog Controller REST API

## Packages

The following is a high-level overview of the functionality provided by each package.

Each package contains its own README.md so please refer to those for further details.

#### Microservices

The `microservices` package contains functionality required to implement edge microservices that run on ioFog Edge Compute Networks. This includes functionality to access microservice configuration, connecting to the control signal websocket, connecting to the messages websocket, and being able to receive and post data messages via REST.

#### Client

The `client` package contains an HTTP client to use with ioFog Controller's REST API. You can view see the full REST API specification at [iofog.org](https://iofog.org/docs/1.3.0/controllers/rest-api.html).