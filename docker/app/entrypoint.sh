#!/bin/sh
# Run service
go mod download
go run ./cmd/service/main.go