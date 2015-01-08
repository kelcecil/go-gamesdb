#go-gamesdb
## Library to interact with TheGamesDB API via Golang

This is a golang library to easily interact with TheGamesDB V1 API. 

### Dependencies

Go 1.4 - go-gamesdb (will) extensively use go:generate to generate most of the required functions.

### Setup

Make sure your PATH includes ${GOPATH}/bin in order for go generate to find the executable to generate the HTTP call code, and ensure you have make installed.  

```
go get github.com/kelcecil/go-gamesdb
cd ${GOPATH}/src/kelcecil/go-gamesdb
make
```

###License
MIT
