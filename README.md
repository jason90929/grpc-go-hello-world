# Create gRPC Go format via protobuff

`$ cd ./pbfiles`
`$ protoc --go_out=../protos --go_opt=paths=source_relative --go-grpc_out=../protos --go-grpc_opt=paths=source_relative echo.proto`

If you received error `protoc --go_out=../protos --go_opt=paths=source_relative --go-grpc_out=../protos --go-grpc_opt=paths=source_relative echo.proto                ⬡ system [±main ●●]
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.`

This solves mine personal env: https://stackoverflow.com/a/65746184

```
$ export GO111MODULE=on
export PATH="$PATH:$(go env GOPATH)/bin"
```

# Start local server and client:

Open two Terminals to start both client and server

```
cd ./unary/server
go run server.go
```

```
cd ./unary/client
go run client.go
```

type anything in client Terminal, receives some logs in both Terminals 

