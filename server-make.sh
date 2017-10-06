#!/bin/bash


CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' server.go
#go build server.go
