# GoLang Playground

## Basic API Project
Creates a basic API with two routes, one GET and one POST that will return a message.
Also creates one test file to test those routes.

### To run the API
```
go mod init api
go get -d "github.com/gin-gonic/gin"
go build api.go # to create the executable in current path
```
or
``` 
go install api.go # to create the executable in GOBIN path
go run api.go
```

### To run the test

```
go test
```
