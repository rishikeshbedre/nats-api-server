# NATS API SERVER

![Licence](https://img.shields.io/github/license/rishikeshbedre/nats-api-server)
[![Build Status](https://travis-ci.com/rishikeshbedre/nats-api-server.svg?branch=master)](https://travis-ci.com/rishikeshbedre/nats-api-server)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/940a0ccb560941fb9cdbd1b277d6af17)](https://app.codacy.com/manual/rishikeshbedre/nats-api-server?utm_source=github.com&utm_medium=referral&utm_content=rishikeshbedre/nats-api-server&utm_campaign=Badge_Grade_Dashboard)
[![codecov](https://codecov.io/gh/rishikeshbedre/nats-api-server/branch/master/graph/badge.svg)](https://codecov.io/gh/rishikeshbedre/nats-api-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/rishikeshbedre/nats-api-server)](https://goreportcard.com/report/github.com/rishikeshbedre/nats-api-server)

NATS API Server is a REST based configuration server for [NATS-Server](https://github.com/nats-io/nats-server). It features REST end-points to configure user authorization and reload the NATS-Server. It is written using [Gin Web Framework](https://github.com/gin-gonic/gin) and [jsoniter](https://github.com/json-iterator/go) for high performant server.

## Contents

- [NATS API SERVER](#nats-api-server)
  - [Contents](#contents)
  - [How it works](#how-it-works)
  - [Usage](#usage)
  - [API Documentation](#api-documentation)
  - [Docker](#docker)
  - [Kubernetes](#kubernetes)
  - [Testing](#testing)

## How it Works

![nats-api-server](https://github.com/rishikeshbedre/nats-api-server/blob/master/extras/nats-api-server.jpg)

NATS API Server has rest end points to add|delete user|topic where it writes the authorization configuration to a file. The API Server also has an option to send reload signal to NATS-Server where it reads this configuration file and allows only authenticated users to connect to NATS-Server.

## Usage

To install NATS API Server, you need to install Go and set your Go workspace first.

1. The first need is [Go](https://golang.org/) installed (**version 1.12+ is required**), this project uses go modules and provides a make file. You should be able to simply:

```sh
$ git clone https://github.com/rishikeshbedre/nats-api-server.git
$ cd nats-api-server
$ make
$ ./nats-api-server
```

2. Then you need to install [NATS-Server](https://docs.nats.io/nats-server/installation#installing-from-the-source) and start the server using the configuration file present in the [NATS API Server](https://github.com/rishikeshbedre/nats-api-server/blob/master/configuration/nats-server.conf).

## API Documentation

### Add User
