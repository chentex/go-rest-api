# go-rest-api

Simple GOlang REST API example with HipChat Integration

Checkout the reposority

`git clone https://github.com/chentex/go-rest-api.git`

Go to the folder

`cd go-rest-api/`

## Run Application

To run the application, build application with

`go build`

The execute with

`./go-rest-api`

From a Browser try this URL

`http://localhost:8080/test`

## Build Image

To build the image with the application run

`docker build -f Dockerfile . -t gorestapi:latest`

To run the image execute

`docker run --rm -p 8080:8080 gorestapi:latest`

To the image try this URLs

`http://localhost:8080/test`


`http://localhost:8080/hola/chentex`
