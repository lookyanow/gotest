#!/bin/bash

go mod tidy
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' server.go

