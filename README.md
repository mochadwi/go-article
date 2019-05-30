# go-article
[![Build Status](https://travis-ci.com/mochadwi/go-article.svg?branch=master)](https://travis-ci.com/mochadwi/go-article)

### Description
This is an article management with replica of Clean Architecture in (Golang) projects.

### Requirements
- Go 1.9+
- Dep

### Tools Used:
- All libraries listed in [`Gopkg.toml`](https://github.com/mochadwi/go-article/blob/master/Gopkg.toml)

### How To Run This Project
- Clone this repository `go get github.com/mochadwi/go-article`
- `cd $GOPATH/src/mochadwi/go-article`
- Run `dep ensure`
- Run `make`
- Run `gin -a 9090 article_clean` or `./gin-bin`
- Run postman on `localhost:9090 # default gin auto-reload port`
- Happy coding

### Install using script
- Run `chmod a+x install.sh`
- Just run `./install.sh` to do the above commands

### Postman Docs
- https://documenter.getpostman.com/view/1033160/RzZ3LNT3


### Tech Stack
- Database (PostgreSQL - Gorm)
- Http (Echo web framework)
