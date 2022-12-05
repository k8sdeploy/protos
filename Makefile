.PHONY: setup
setup:
	mkdir -p generated
	mkdir -p generated/key
	mkdir -p generated/permissions
	mkdir -p generated/orchestrator
	mkdir -p generated/company

.PHONY: generate
generate: setup key-service permission-service orchestrator company-service

.PHONY: key-service
key-service:
	protoc -I . \
		--go_out ./generated/key --go_opt paths=source_relative \
		--go-grpc_out ./generated/key --go-grpc_opt paths=source_relative \
		v1/key.proto

.PHONY: permission-service
permission-service:
	protoc -I . \
		--go_out ./generated/permissions --go_opt paths=source_relative \
		--go-grpc_out ./generated/permissions --go-grpc_opt paths=source_relative \
		v1/permissions.proto

.PHONY: orchestrator
orchestrator:
	protoc -I . \
		--go_out ./generated/orchestrator --go_opt paths=source_relative \
		--go-grpc_out ./generated/orchestrator --go-grpc_opt paths=source_relative \
		v1/orchestrator.proto

.PHONY: company-service
company-service:
	protoc -I . \
		--go_out ./generated/company --go_opt paths=source_relative \
		--go-grpc_out ./generated/company --go-grpc_opt paths=source_relative \
		v1/company.proto

