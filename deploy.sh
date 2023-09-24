#!/bin/bash

# Install dependencies
go get 
go mod download
go install github.com/swaggo/swag/cmd/swag@latest
go mod vendor -v

# Generate Swagger
swag init

# Run main.go
go run main.go