# Golang NATS

[![Go](https://img.shields.io/badge/go-1.11.1-00E5E6.svg)](https://golang.org/)
[![Dep](https://img.shields.io/badge/dep-0.4.1-A78552.svg)](https://golang.github.io/dep/)

This is golang example for use NATS with publisher and queue subscriber.
Before run this apps, you must installed [NATS Streaming Server](https://nats.io/) first.

## Libraries
Use [Golang Dep](https://golang.github.io/dep/) command for install all dependencies required this application.
  - Use command `dep ensure -v` for install all dependency.
  - Use command `dep ensure -v -update` for update all dependency.
  
## How To Run and Deploy
Before run this service. Make sure all requirements dependencies has been installed likes **Golang, NATS, and Dep**

### Local
This app has 2 folder, `publish` and `subscribe`.

App in folder `publish` used for send message to NATS server and App in folder `subscribe` used for consuming message from NATS server.

Use command `go run main.go` in `publish` or `subscribe` folder for run this application.