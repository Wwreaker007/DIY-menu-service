run-orders:
	@go run orders/*.go

run-kitchen:
	@go run kitchen/*.go

gen-orders:
	@protoc --proto_path=common/types "common/types/orders.proto" --go_out=common/codegen/orders --go_opt=paths=source_relative --go-grpc_out=common/codegen/orders --go-grpc_opt=paths=source_relative