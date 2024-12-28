#!/bin/bash -x

# Step 1: Get all the dependency Modules
go mod tidy

# Step 2: Set Swag Path
export PATH=$(go env GOPATH)/bin:$PATH

# Step 3: Generate Swagger documentations 
swag init --pdl 3

# Step 4: Build User Service application
go build main.go