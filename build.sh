#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o playerBackground main.go
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o playerBackground main.go
mv broker api/
cp -r conf api/