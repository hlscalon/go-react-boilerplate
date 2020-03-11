#!/bin/bash

# first format all packages
go fmt ./...

# then check if something is wrong
go vet ./...

# unit tests
go test ./... -v
