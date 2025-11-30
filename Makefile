generate:
	protoc \
	  --proto_path=schema/proto \
	  --proto_path=schema/proto/third_party \
	  --go_out=paths=source_relative:pkg \
	  --go-grpc_out=paths=source_relative:pkg \
	  --grpc-gateway_out=paths=source_relative:pkg \
	  --grpc-gateway_opt=logtostderr=true \
	  --grpc-gateway_opt=generate_unbound_methods=true \
	  --validate_out=lang=go,paths=source_relative:pkg \
	  --openapiv2_out=docs/ \
	  --openapiv2_opt=logtostderr=true \
	  schema/proto/dart/gateway/project/v1/project.proto \
	  schema/proto/dart/gateway/project/v1/project_service.proto

.PHONY: generate