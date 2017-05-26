gen-grpc-stubs:
	@protoc --go_out=plugins=grpc:pb ./*.proto

run-server:
	@go run server/server.go

run-client:
	@go run client/client.go
