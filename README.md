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
  - [API documentation](#api-documentation)
  - [Docker](#docker)
  - [Kubernetes](#kubernetes)
  - [Testing](#testing)

## How it Works

<img align="center" width="500px" src="https://raw.githubusercontent.com/rishikeshbedre/nats-api-server/extras/nats-api-server.jpg">
