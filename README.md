# REST API Tempalate 

This is a REST API template, using Clean Architecture, and written in Golang. 

## Contents

 * [Requirements](#Requirements) 
 * [Installation](#Installation) 
 * [Endpoints and Operations](#Endpoints-and-Operations)
 * [Configuration](#Configuration) 
 * [Make Commands](#Make-Commands) 

## Requirements

This progrom does require a couple of things:

 * [Golang 1.16+](https://go.dev/dl/)

## Installation 

Clone the source 

`git clone https://github.com/alexmerren/rest-api-template.git`

Build and run the app

`make build && make run`

And visit

`http://localhost:8080/api/v1/health`

## Endpoints and Operations

This template implements a few basic CRUD operations, each of which has a specific endpoint.

| Endpoint                    | Method | Operation                                                                      |
| --------------------------- | ------ | ------------------------------------------------------------------------------ |
| /api/v1/health/             | GET    | A health check to ensure that the service has started correctly                |
| /api/v1/contacts/create/     | POST   | Create new contacts, with the parameters as the request body.                 |
| /api/v1/contacts/read/       | GET    | Read all contacts that are stored.                                             |
| /api/v1/contacts/read/{id}   | GET    | Read a contact whose ID is given as a parameter in the URL.                    |
| /api/v1/contacts/update/{id} | PUT    | Update a contact whose ID is in the URL, and body containing fields to update. |
| /api/v1/contacts/delete/{id} | POST   | Delete a contact whose ID is in the URL.                                       |

## Configuration

| Settings          | Environment Variable   | Description                   | Default                |
| ----------------- | ---------------------- | ----------------------------- | ---------------------- |
| server.port       | REST\_SERVER\_PORT     | The port number to listen on  | 8080                   |
| logger.loglevel   | REST\_LOGGER\_LOGLEVEL | The default loggin level      | "debug"                |

## Make Commands

```
help           Print this message
build          Create the binary
run            Run the binary
vendor         Download the vendored dependencies
lint           Lint the project
test           Run the unit tests for the project
mocks          Generate mocks for the project
docker-build   Build the docker container
docker-run     Run the docker container with some environment variables
```
