#!/bin/sh
GO111MODULE=off
rm -rf ./logs/*.log
go run main.go
