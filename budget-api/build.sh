#!/bin/bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build kunpeng/budget-api
docker build -t budget-api .