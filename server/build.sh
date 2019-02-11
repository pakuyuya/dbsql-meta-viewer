#!/bin/bash
dep ensure
go build -o dist/server main.go
