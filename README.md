This example will get two numbers from client and returns greatest common divisor of that numbers.

#How to run

1. Generate go code from proto file

```
cd pb
protoc -I . --go_out=plugins=grpc:. ./*.proto
```

2. Run server

```
go run server.go
```

3. Run client on other terminal

```
go run client.go 12 10
```