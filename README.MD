# Manual instalation

## Requirement
please make sure you already have:

* Golang installed

*  golang wire

    ```go get github.com/google/wire/cmd/wire```

    and ensuring that $GOPATH/bin is added to your $PATH.

*  docker
see [docker documentation](https://docs.docker.com/) 

* postgres database

## Run Containerized Postgres if have no postgres installed
docker run --name postgres -p 5432:5432 -d postgres

## Setup environtment variable 
you can see the example from .env.example


## Run Application 
```go run cmd/api/main.go cmd/api/wire_gen.go```


# Run using docker

## Requirement

*  docker
see [docker documentation](https://docs.docker.com/) 

## Run application