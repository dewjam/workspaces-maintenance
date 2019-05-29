#!/bin/bash

echo "Installing Dependencies..."
glide install

echo "Testing Code..."
go test

echo "Building Binary..."
export GOARCH="amd64"
export GOOS="linux"
go build -o bin/wm

echo "Cleaning up..."
unset GOARCH
unset GOOS
