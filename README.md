# Simple Service for generate short hash v.0.0.1
## Protoc
1. install protoc
2.  ```go mod tidy ```
3.  ``` go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 ```

## GO:GENERATE
    `go generate ./...`

## Build (Test)
1. `docker build -t hash_service:v0.0.1 .` 
2. `docker run -rm -ti hash_service:v0.0.1`