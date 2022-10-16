.PHONY: client server

all:
	echo "Usage:"
	echo "Run client: make client"
	echo "Run server: make server"
client:
	go run client/client.go

server:
	go run server/server.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protocol/protocol.proto
