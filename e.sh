#!/usr/bin/zsh
go build
./euler -p $1
go clean
