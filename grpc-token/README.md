## compile .proto
```sh
cd ~/Desktop/grpc-token

protoc \
--go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
auth/auth.proto
```
