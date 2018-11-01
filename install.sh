#!/usr/bin/env bash

# install gin auto-reload web server for golang
go get github.com/codegangsta/gin

# verify installation of gin
gin -h

# install deps of this project
dep ensure

# build binary
make

# auto-reload, change your port
gin -a 3333 article_clean
