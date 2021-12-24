# Golang REST API Example 
 
 This is an example REST API framework written in Golang, using Mux, and a Zap logger.

## Contents

 * [Requirements](#Requirements) 
 * [Configuration](#Configuration) 
 * [Installation](#Installation) 

## Requirements

This progrom does require a couple of things:

 * [mySQL](https://dev.mysql.com/downloads/)
 * [Golang 1.16](https://go.dev/dl/)

## Configuration

| Settings          | Description                   | Default                |
| ----------------- | ----------------------------- | ---------------------- |
| Host              | The host address to listen on | "localhost"            |
| Port              | The port number to listen on  | 8080                   |
| Logger.Level      | The default loggin level      | "debug"                |
| Logger.Encoding   | Logging format                | "console"              |
| Database.Username | The database username         | "h9zK7nneEB"           |
| Database.Password | The database user's password  | "e4gJhwu2Nk5KS5Kqa5Ue" |
| Database.Name     | The database name             | "contacts"             |
| Database.Port     | The database port             | 3306                   |

## Installation 

Clone the source 

`git clone https://github.com/alexmerren/golang-api-template.git`

Build and run the app

`make build && make run`

And visit

`http://localhost:8080/api/test`

## Make Commands

```
help     Print this message
build    Create the binary
run      Run the binary
vendor   Download the vendored dependencies
lint     Lint the project
test     Test the project
rename   Rename the project
```
