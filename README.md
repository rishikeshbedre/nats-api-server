# NATS API SERVER

![Licence](https://img.shields.io/github/license/rishikeshbedre/nats-api-server)
[![Build Status](https://travis-ci.com/rishikeshbedre/nats-api-server.svg?branch=master)](https://travis-ci.com/rishikeshbedre/nats-api-server)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/940a0ccb560941fb9cdbd1b277d6af17)](https://app.codacy.com/manual/rishikeshbedre/nats-api-server?utm_source=github.com&utm_medium=referral&utm_content=rishikeshbedre/nats-api-server&utm_campaign=Badge_Grade_Dashboard)
[![codecov](https://codecov.io/gh/rishikeshbedre/nats-api-server/branch/master/graph/badge.svg)](https://codecov.io/gh/rishikeshbedre/nats-api-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/rishikeshbedre/nats-api-server)](https://goreportcard.com/report/github.com/rishikeshbedre/nats-api-server)

NATS API Server is a REST based configuration server for [NATS-Server](https://github.com/nats-io/nats-server). It features REST end-points to configure user authorization and reload the NATS-Server. It is written using [Gin Web Framework](https://github.com/gin-gonic/gin) and [jsoniter](https://github.com/json-iterator/go) to make server high performant.

## Contents

- [NATS API SERVER](#nats-api-server)
  - [Contents](#contents)
  - [How it works](#how-it-works)
  - [Usage](#usage)
  - [API Documentation](#api-documentation)
    - [Add User](#add-user)
    - [Delete User](#delete-user)
    - [Show User](#show-user)
    - [Add Topic](#add-topic)
    - [Delete Topic](#delete-topic)
    - [Download Configuration](#download-configuration)
  - [Docker](#docker)
  - [Kubernetes](#kubernetes)
  - [Testing](#testing)

## How it Works

![nats-api-server](https://github.com/rishikeshbedre/nats-api-server/blob/master/extras/nats-api-server.jpg)

NATS API Server has rest end points to add|delete user|topic where it writes the authorization configuration to a file. The API Server also has an option to send reload signal to NATS-Server where it reads this configuration file and allows only authenticated users to connect to NATS-Server.

## Usage

To install NATS API Server, you need to install [Go](https://golang.org/)(**version 1.12+ is required**) and set your Go workspace.

1. This project uses go modules and provides a make file. You should be able to simply install and start:

```sh
$ git clone https://github.com/rishikeshbedre/nats-api-server.git
$ cd nats-api-server
$ make
$ ./nats-api-server
```

2. Then you need to install [NATS-Server](https://docs.nats.io/nats-server/installation#installing-from-the-source) and start the server using the configuration file present in the [NATS API Server](https://github.com/rishikeshbedre/nats-api-server/blob/master/configuration/nats-server.conf).

## API Documentation

### Add User

Adds new user to the authorization configuration.

- **URL:**
  `/user`

- **Method:**
  `POST`

- **Request:**
  - **Header:**
    - **Content-Type:** `application/json`
  - **Body:** `{"user":"xyz","password":"123"}`

- **Success Response:**
  - **Code:** `200` 
  - **Content:** `{"message":"User:xyz added"}`
 
- **Error Response:**
  - **Code:** `400 STATUS BAD REQUEST` 
  - **Content:** `{"error":"User:xyz already present"}`

  OR

  - **Code:** `400 STATUS BAD REQUEST` 
  - **Content:** `{"error":"Key: 'AddUserJSON.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`

- **Sample Call:**

  ```ssh
    $curl --header "Content-Type: application/json" --request POST --data '{"user":"xyz","password":"123"}' http://localhost:6060/user
  ```

### Delete User

Deletes the user from authorization configuration.

- **URL:**
  `/user`

- **Method:**
  `DELETE`

- **Request:**
  - **Header:**
    - **Content-Type:** `application/json`
  - **Body:** `{"user":"xyz"}`

- **Success Response:**
  - **Code:** `200` 
  - **Content:** `{"message":"User:xyz deleted"}`
 
- **Error Response:**
  - **Code:** `400 STATUS BAD REQUEST`
  - **Content:** `{"error":"User:xyz cannot be deleted"}`

  OR

  - **Code:** `400 STATUS BAD REQUEST`
  - **Content:** `{"error":"Key: 'DeleteUserJSON.User' Error:Field validation for 'User' failed on the 'required' tag"}`

- **Sample Call:**

  ```ssh
    $curl --header "Content-Type: application/json" --request DELETE --data '{"user":"xyz"}' http://localhost:6060/user
  ```

### Show User

Returns the current authorization configuration.

- **URL:**
  `/user`

- **Method:**
  `GET`

- **Request:** `NONE`

- **Success Response:**
  - **Code:** `200` 
  - **Content:** `{"message":[{"user":"natsdemouser","permissions":{"publish":null,"subscribe":null}}]}`
 
- **Error Response:**
  - **Code:** `400 STATUS BAD REQUEST`
  - **Content:** `{"error":"???jsonbinderror"}`

- **Sample Call:**

  ```ssh
    $curl --request GET http://localhost:6060/user
  ```

### Add Topic

Adds the topics to the particular user in authorization configuration. If any of the topics are present in the request JSON are available in the authorization configuration for that particular user, this end point returns a error message.

- **URL:**
  `/topic`

- **Method:**
  `POST`

- **Request:**
  - **Header:**
    - **Content-Type:** `application/json`
  - **Body:** `{"user":"xyz","permissions":{"publish":["test","quest"],"subscribe":["test","quest"]}}`

- **Success Response:**
  - **Code:** `200` 
  - **Content:** `{"message":"Topics Added for the user:xyz"}`
 
- **Error Response:**
  - **Code:** `400 STATUS BAD REQUEST` 
  - **Content:** `{"error":"test topic is already present for the user:xyz"}`

  OR

  - **Code:** `400 STATUS BAD REQUEST` 
  - **Content:** `{"error":"Key: 'AddDeleteTopicJSON.User' Error:Field validation for 'User' failed on the 'required' tag"}`

- **Sample Call:**

  ```ssh
    curl --header "Content-Type: application/json" --request POST --data '{"user":"xyz","permissions":{"publish":["test","quest"],"subscribe":["test","quest"]}}' http://localhost:6060/topic
  ```

### Delete Topic

Deletes the topics from the particular user in authorization configuration. If any of the topics are present in the request JSON are not available in the authorization configuration for that particular user, this end point returns a error message.

- **URL:**
  `/topic`

- **Method:**
  `DELETE`

- **Request:**
  - **Header:**
    - **Content-Type:** `application/json`
  - **Body:** `{"user":"xyz","permissions":{"publish":["quest"],"subscribe":["quest"]}}`

- **Success Response:**
  - **Code:** `200` 
  - **Content:** `{"message":"Topics deleted for the user:xyz"}`
 
- **Error Response:**
  - **Code:** `400 STATUS BAD REQUEST` 
  - **Content:** `{"error":"Cannot delete topics for the user:xyz"}`

  OR

  - **Code:** `400 STATUS BAD REQUEST` 
  - **Content:** `{"error":"Key: 'AddDeleteTopicJSON.User' Error:Field validation for 'User' failed on the 'required' tag"}`

- **Sample Call:**

  ```ssh
    curl --header "Content-Type: application/json" --request DELETE --data '{"user":"xyz","permissions":{"publish":["quest"],"subscribe":["quest"]}}' http://localhost:6060/topic
  ```

### Download Configuration

Stores the authorization configuration to the file and reload the nats server.<br>
**Note:** Until you send this request to NATS API Server, add|delete user|topic requests doesn't reflect in NATS Server.

- **URL:**
  `reload`

- **Method:**
  `POST`

- **Request:** `NONE`

- **Success Response:**
  - **Code:** `200` 
  - **Content:** `{"message":"Download and reload of Configuration Successful"}`
 
- **Error Response:**
  - **Code:** `400 STATUS BAD REQUEST` 
  - **Content:** `{"error":"??filewriteerror or ??jsonbinderror or ??cmderror"}`

- **Sample Call:**

  ```ssh
    curl --request POST http://localhost:6060/reload
  ```

## Docker

Building the image for nats api server acutually builds both nats api server and nats server in one container, so when you run the container two services will run in the same container.

1. To build the image run following command:

  ```ssh
    $./extras/build.sh
  ```

2. While running the image you can persist the configuration file by mounting the volume to the host and container. To run the container just run the following command:

  ```ssh
    $docker run -it -p 4222:4222 -p 6060:6060 -v /home/rishikesh/Desktop/nats-data:/home/nats/configuration nats-api-server:0.0.1
  ```

## Kubernetes

You can run this setup in kubernetes also by using this [yaml file](https://github.com/rishikeshbedre/nats-api-server/blob/master/extras/nats-api-server.yaml):

  ```ssh
    $kubectl apply -f ./nats-api-server.yaml
  ```

## Testing

To run test just run following command:

  ```ssh
    $go mod download
    $make test
  ```