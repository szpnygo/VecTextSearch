#!/bin/bash
export $(grep -v '^#' .env | xargs)
go run ./cmd/main.go
