#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o release/gotain main.go
