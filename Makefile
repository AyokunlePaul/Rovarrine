proto:
	protoc accounts/rpc/*.proto --go_out=plugins=grpc:.
	protoc transactions/rpc/*.proto --go_out=plugins=grpc:.
server:
	go run main_server.go
client:
	go run main_client.go