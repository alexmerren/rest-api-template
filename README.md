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

 * [Golang 1.16](https://go.dev/dl/)

## Installation 

Clone the source 

`git clone https://github.com/alexmerren/golang-api-template.git`

Build and run the app

`make build && make run`

And visit

`http://localhost:8080/api/health`

## Endpoints and Operations

This template implements a few basic CRUD operations, each of which has a specific endpoint.

| Endpoint         | Method | Operation                                                                      |
| ---------------- | ------ | ------------------------------------------------------------------------------ |
| /api/health/     | GET    | A health check to ensure that the service has started correctly                |
| /api/create/     | POST   | Create a new contact, with the parameters as the request body.                 |
| /api/read/       | GET    | Read all contacts that are stored.                                             |
| /api/read/{id}   | GET    | Read a contact whose ID is given as a parameter in the URL.                    |
| /api/update/{id} | PUT    | Update a contact whose ID is in the URL, and body containing fields to update. |
| /api/delete/{id} | POST   | Delete a contact whose ID is in the URL.                                       |

## Configuration

| Settings          | Description                   | Default                |
| ----------------- | ----------------------------- | ---------------------- |
| port              | The port number to listen on  | 8080                   |
| logger.loglevel   | The default loggin level      | "debug"                |

## Make Commands

```
help     Print this message
build    Create the binary
run      Run the binary
vendor   Download the vendored dependencies
lint     Lint the project
test     Test the project
mocks    Generate mocks for the project
```
