gen-grpc-stubs:
	@protoc --go_out=plugins=grpc:pb ./*.proto

gen-certs:
	@openssl genrsa -out cert.key 2048
	@openssl req -new -x509 -sha256 -key cert.key -out cert.crt -days 3650
