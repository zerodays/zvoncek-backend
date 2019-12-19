#!/bin/bash

export GOOS=linux
export GOARCH=amd64

# Uncomment if you want to build for arm
# export GOARCH=arm

echo "Pruning dependencies"
go mod tidy

echo "Compiling..."
go build -o build/zvon

echo "Copmpiling done!"
echo "Copying config files"
cp config.json build/

echo "Building done. Files are located in 'build/'"
