PACKAGE_ROOT=api
OUT_DIR=internal/pb
GATEWAY_INCLUDE=$(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2)

PROTOC_GEN_GO=$(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC=$(shell which protoc-gen-go-grpc)
PROTOC_GEN_GRPC_GATEWAY=$(shell which protoc-gen-grpc-gateway)
PROTOC_GEN_OPENAPIV2=$(shell which protoc-gen-openapiv2)

.PHONY: all proto check-tools install-tools

all: proto

check-tools:
ifndef PROTOC_GEN_GO
	$(error protoc-gen-go is not installed)
endif
ifndef PROTOC_GEN_GO_GRPC
	$(error protoc-gen-go-grpc is not installed)
endif
ifndef PROTOC_GEN_GRPC_GATEWAY
	$(error protoc-gen-grpc-gateway is not installed)
endif
ifndef PROTOC_GEN_OPENAPIV2
	$(error protoc-gen-openapiv2 is not installed)
endif

proto: check-tools
	find $(PACKAGE_ROOT) -name "*.proto" | while read -r file; do \
		echo "Generating for $$file"; \
		REL_PATH=$$(dirname $$file); \
		OUT_PATH=$(OUT_DIR)/$$REL_PATH; \
		mkdir -p "$$OUT_PATH"; \
		protoc \
			-I . \
			-I ./vendor.protogen \
			-I $(GATEWAY_INCLUDE) \
			--go_out=$(OUT_DIR) \
			--go_opt=paths=source_relative \
			--go_opt=Mapi/v1/model/drones.proto=github.com/VeneLooool/drones-api/internal/pb/api/v1/model \
			--go-grpc_out=$(OUT_DIR) \
			--go-grpc_opt=paths=source_relative \
			--grpc-gateway_out=$(OUT_DIR) \
			--grpc-gateway_opt=paths=source_relative \
			--openapiv2_out="$$OUT_PATH" \
			--openapiv2_opt="logtostderr=true" \
			$$file; \
	done

install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
