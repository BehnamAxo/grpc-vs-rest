# gRPC vs REST in Go

This project compares **gRPC** and **REST** API performance using Go. It includes services implemented in both styles and benchmarks them using realistic payloads and high-concurrency test scripts.


## Requirements

- **Go** 1.18+
- **Node.js** 18+
- [`protoc`](https://grpc.io/docs/protoc-installation/)
- [`protoc-gen-go`](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go)
- [`protoc-gen-go-grpc`](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc)
- [`ghz`](https://github.com/bojand/ghz) (installed globally)
- [`autocannon`](https://github.com/mcollina/autocannon) (installed locally in this project)


## 🔧 Setup Instructions

### gRPC Setup

1. Initialize Go modules inside the `grpc` folder:

   ```bash
   cd grpc
   go mod init grpc
   ```

2. Generate the gRPC Go code from the .proto file (from the project root):

   ```bash
    protoc --go_out=. --go-grpc_out=. proto/user.proto
   ```

3. Install the gRPC dependency:

   ```bash
    go get google.golang.org/grpc
   ```


## Running the Services

### ▶️ gRPC Server

Start the gRPC payment service:

   ```bash
    cd grpc/payment_service
    go run main.go
   ```

### ▶️ REST Server

Start the gRPC payment service:

   ```bash
    cd rest/payment_service
    go run main.go
   ```

## 🧪 Running Benchmarks

### gRPC Load Test

```bash
node grpcTest.js
```


### REST Load Test

```bash
node restTest.js
```


## 📁 Payloads

This repo includes:

- **`payload-small.json`** light payload

- **`payload-large.json`** realistic medium payload

- **`payload-huge.json`** stress test payload

You can modify which payload is used by editing the script files:

- `grpcTest.js`
- `restTest.js`


## Notes
- gRPC uses Protocol Buffers over HTTP/2, providing smaller binary payloads and lower latency.

- REST uses JSON over HTTP/1.1, which is more human-readable but typically less efficient.

- This project helps you compare both under load with different payload sizes.
