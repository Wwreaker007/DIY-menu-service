run-orders:
	@go run services/orders/*.go

run-kitchen:
	@go run services/kitchen/*.go

gen-orders:
	@protoc --proto_path=common/types "common/types/orders.proto" --go_out=common/codegen/orders --go_opt=paths=source_relative --go-grpc_out=common/codegen/orders --go-grpc_opt=paths=source_relative