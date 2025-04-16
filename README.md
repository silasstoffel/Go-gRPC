# Go-gRPC

## Setup

```shell
docker-compose up -d

# accessing the container
./scripts/golang-app.sh


# run GRPC server
go run cmd/grpc-server/main.go
```

**How use Evans gRPC client**

```shell
evans -r repl

package pb

service CategoryService

```