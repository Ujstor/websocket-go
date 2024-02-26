# Project websockets

Simple ws implementation within the blueprint boilerplate.

## Getting Started

```bash
make run
```

Run clinets

```bash
go run client/ws-data-read/main.go

go run client/ws-chat/main.go
```

### Using docker

When docker is used, clinets connection string must mach compose service name.
```go
websocket.Dial("ws://server:8082/ws")
```

```bash
make docker-run
```

```bash
docker exec -it <container-id-chat1> sh
./clinet-chat
```

```bash
docker exec -it <container-id-chat2> sh
./clinet-chat
```


## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```