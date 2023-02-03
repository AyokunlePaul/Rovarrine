proto:
	protoc accounts/rpc/*.proto --go_out=plugins=grpc:.
	protoc transactions/rpc/*.proto --go_out=plugins=grpc:.
regenerate:
	protoc --go_out=. --go_opt=paths=source_relative accounts/rpc/*.proto
	protoc --go_out=. --go_opt=paths=source_relative transactions/rpc/*.proto
server:
	go run account_server.go
client:
	go run account_client.go
transaction_server:
	go run transaction_server.go