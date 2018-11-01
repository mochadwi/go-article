# go-article

### Description
This is an article management with replica of Clean Architecture in (Golang) projects.

### Requirements
- Go 1.9+
- Dep

### Tools Used:
- All libraries listed in [`Gopkg.toml`](https://github.com/mochadwi/go-article/blob/master/Gopkg.toml)

### How To Run This Project
- Run `go get github.com/codegangsta/gin`
- Verify it `gin -h`
- Then, clone this repository `go get github.com/mochadwi/go-article`
- `cd $GOPATH/src/mochadwi/go-article`
- Run `dep ensure`
- Run `make`
- Run `gin -a 3333 article_clean` or `./gin-bin`
- Run postman on `localhost:3000 # default gin auto-reload port`
- Happy coding

### Install using script
- Run `chmod a+x install.sh`
- Just run `./install.sh` to do the above commands

### Postman Docs
- https://documenter.getpostman.com/view/1033160/RzZ3LNT3


### Todo
[ * ] Docker

[ * ] Implement Iris
