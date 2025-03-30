# gRPC vs REST in Go 🚀

This project compares **gRPC** and **REST** API performance using Go. It includes simple backend and payment services implemented in both styles and tests them under load using appropriate benchmarking tools.

## ⚙️ gRPC Setup (If You're Starting from Scratch)

1. Initialize the module (inside the `grpc` folder):

  ```bash
  go mod init grpc
  ```

2. Generate the gRPC code from your .proto file. Make sure the `generated` folder exists first.

  ```bash
  protoc --go_out=generated --go-grpc_out=generated proto/user.proto
  ```

3. Install the gRPC package:

  ```bash
  go get google.golang.org/grpc
  ```
 Make sure protoc, protoc-gen-go, and protoc-gen-go-grpc are installed and available in your PATH.

## How to Run

### ✅ gRPC

1. **Start gRPC payment service:**

```bash
cd grpc/payment_service
go run main.go
```

2. **Run gRPC backend client:**

```bash
cd grpc/backend_service
go run main.go
```

3. **Benchmark gRPC using [`ghz`](https://github.com/bojand/ghz):**

```bash
ghz --insecure --proto ./proto/user.proto --call userpb.PaymentService.ProcessUser -d "{\"name\":\"Sir Laughsalot McGiggles\",\"age\":420,\"email\":\"funny.bone@laughterverse.io\",\"phone\":\"+1-800-GIGGLEZ\"}" -c 100 --duration 20s localhost:50051
```

### ✅ REST

1. **Start REST payment service:**

```bash
cd rest/payment_service
go run main.go
```

2. **Run REST backend client:**

```bash
cd rest/backend_service
go run main.go
```

3. **Benchmark REST using [`autocannon`](https://github.com/mcollina/autocannon):**

```bash
autocannon -c 100 -d 20 -m POST -H "Content-Type: application/json" -b "{\"name\":\"Sir Laughsalot McGiggles\",\"age\":420,\"email\":\"funny.bone@laughterverse.io\",\"phone\":\"+1-800-GIGGLEZ\"}" http://localhost:8080/process
```

## 🧪 Load Testing Tools Used

- 🧬 [`ghz`](https://github.com/bojand/ghz) for benchmarking **gRPC**
- 🧪 [`autocannon`](https://github.com/mcollina/autocannon) for benchmarking **REST**

## 📦 Requirements

- Go 1.18+
- `protoc`, `protoc-gen-go`, and `protoc-gen-go-grpc` (for gRPC)
- [`ghz`](https://github.com/bojand/ghz) for load testing gRPC
- [`autocannon`](https://github.com/mcollina/autocannon) for load testing REST (Node.js required)

## 📌 Notes

- gRPC uses Protocol Buffers for efficient binary encoding over HTTP/2
- REST uses JSON over HTTP/1.1, which is more human-readable but generally slower
- This repo demonstrates how each performs under heavy load
