# gRPC Example

To run this example, you need to generate the Go code from the protobuf file.

1. Install protoc compiler.
2. Install Go plugins:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
3. Generate code:
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/service.proto
   ```
4. Uncomment the code in `server/main.go` and `client/main.go` that uses the generated `pb` package.

5. Run server:
   ```bash
   go run server/main.go
   ```
6. Run client:
   ```bash
   go run client/main.go
   ```
