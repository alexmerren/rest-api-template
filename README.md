# Golang REST API Example 
 
 This is an example REST API framework written in Golang, using Mux for
 routing, Zap for logging, and Koanf for config. 

## Contents

 * [Requirements](#Requirements) 
 * [Installation](#Installation) 
 * [Endpoints and Operations](#Endpoints-and-Operations)
 * [Configuration](#Configuration) 
 * [Make Commands](#Make-Commands) 

## Requirements

This progrom does require a couple of things:

 * [mySQL](https://dev.mysql.com/downloads/)
 * [Golang 1.16](https://go.dev/dl/)

## Installation 

Clone the source 

`git clone https://github.com/alexmerren/golang-api-template.git`

Build and run the app

`make build && make run`

And visit

`http://localhost:8080/api/test`

## Endpoints and Operations

This template implements a few basic CRUD operations, each of which has a specific endpoint.

| Endpoint         | Method | Operation                                                                      |
| ---------------- | ------ | ------------------------------------------------------------------------------ |
| /api/test/       | GET    | A test endpoint to check if the system is operational.                         |
| /api/create/     | POST   | Create a new contact, with the parameters as the request body.                 |
| /api/read/       | GET    | Read all contacts that are stored.                                             |
| /api/read/{id}   | GET    | Read a contact whose ID is given as a parameter in the URL.                    |
| /api/update/{id} | POST   | Update a contact whose ID is in the URL, and body containing fields to update. |
| /api/delete/{id} | POST   | Delete a contact whose ID is in the URL.                                       |

## Configuration

| Settings          | Description                   | Default                |
| ----------------- | ----------------------------- | ---------------------- |
| host              | The host address to listen on | "localhost"            |
| port              | The port number to listen on  | 8080                   |
| logger.level      | The default loggin level      | "debug"                |
| logger.encoding   | Logging format                | "console"              |
| database.username | The database username         | "h9zK7nneEB"           |
| database.password | The database user's password  | "e4gJhwu2Nk5KS5Kqa5Ue" |
| database.name     | The database name             | "contacts"             |
| database.port     | The database port             | 3306                   |

## Make Commands

```
help     Print this message
build    Create the binary
run      Run the binary
vendor   Download the vendored dependencies
lint     Lint the project
test     Test the project
```
