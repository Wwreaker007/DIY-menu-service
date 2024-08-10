run-orders:
	@go run orders/*.go

run-kitchen:
	@go run kitchen/*.go

gen-orders:
	@protoc --proto_path=common/protos "common/protos/orders.proto" \
	--go_out=common/codegen/orders \
	--go_opt=paths=source_relative \
	--go-grpc_out=common/codegen/orders \
	--go-grpc_opt=paths=source_relative \


gen-cookhouse:
	@protoc --proto_path=common/protos "common/protos/cookhouse.proto" \
	--go_out=common/codegen/cookhouse \
	--go_opt=paths=source_relative \
	--go-grpc_out=common/codegen/cookhouse \
	--go-grpc_opt=paths=source_relative

gen-common:
	@protoc --proto_path=common/protos/common "common/protos/common/common.proto" \
	--go_out=common/codegen/common \
	--go_opt=paths=source_relative \
	--go-grpc_out=common/codegen/common \
	--go-grpc_opt=paths=source_relative

gen-all:
	@make gen-common
	@make gen-orders
	@make gen-cookhouse