# go-article
[![Build Status](https://travis-ci.com/mochadwi/go-article.svg?branch=master)](https://travis-ci.com/mochadwi/go-article)[![codecov](https://codecov.io/gh/mochadwi/go-article/branch/master/graph/badge.svg)](https://codecov.io/gh/mochadwi/go-article)

### Description
This is an article management with replica of Clean Architecture in (Golang) projects.

### Requirements
- Go 1.12+
- Dep

### Tools Used:
- All libraries listed in [`Gopkg.toml`](https://github.com/mochadwi/go-article/blob/master/Gopkg.toml)

### How To Run This Project
- Clone this repository `go get github.com/mochadwi/go-article`
- Make sure to modify a `config.json` on this repo is only a placholder.
- Install and run your postgres 11.2 DB `pg_ctl -D /path/to/postgres star`
- `cd $GOPATH/src/mochadwi/go-article`
- Run `dep ensure`
- Run `make`
- Run `gin -p 9091 -a 9090 article_clean` or `./gin-bin`
- Run postman on `localhost:9091`
- Happy coding

### Install using script
- Run `chmod a+x install.sh`
- Just run `./install.sh` to do the above commands

### Postman Docs
- https://documenter.getpostman.com/view/1033160/RzZ3LNT3


### Tech Stack
- Database (PostgreSQL - Gorm)
- Http (Echo web framework)
