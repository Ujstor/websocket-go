# Project websockets

Simple ws implementation within the blueprint boilerplate.

## Getting Started

```bash
make run
```

Run clinets

```bash
go run clinet/ws-data-read/main.go

go run clinet/ws-chat/main.go
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