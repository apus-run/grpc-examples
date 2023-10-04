## 运行 server
```bash
cd server
go build -o server
```

## 运行 client
```bash
cd client
go build -o client
./client localhost:50051

go run main.go localhost:50051
```

## 运行 client-json
```bash
cd client-json
go build -o client-json
./client-json localhost:50051 '{"email": "jane@doe.com", "id": "1"}'

go run main.go localhost:50051 '{"email": "jane@doe.com", "id": "1"}'
```