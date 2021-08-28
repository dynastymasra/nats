# Golang NATS Jetstream

[![Go](https://img.shields.io/badge/go-1.17.0-00E5E6.svg)](https://golang.org/)

This is golang example for use NATS with publisher and queue subscriber.
Before run this apps, you must installed [NATS Streaming Server](https://nats.io/) first.
  
## How To Run and Deploy
Before run this service. Make sure all requirements dependencies has been installed likes **Golang, and NATS**

### Local
This app has 2 folders, `publish` and `subscribe`.

App in folder `publish` used for send message to NATS server and App in folder `subscribe` used for consuming message from NATS server.

Use command `go run main.go` in `publish` or `subscribe` folders for run this application.