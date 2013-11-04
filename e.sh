#!/usr/bin/zsh
go build
time ./euler -p $1
go clean
go fmt 
