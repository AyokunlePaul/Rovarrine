proto:
	protoc accounts/rpc/*.proto --go_out=plugins=grpc:.
	protoc transactions/rpc/*.proto --go_out=plugins=grpc:.