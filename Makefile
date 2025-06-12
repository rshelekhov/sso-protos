include vendor.proto.mk

# Use bin in the current directory to install protoc plugins
LOCAL_BIN := $(CURDIR)/bin

# Add bin in the current directory to the PATH when running protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Path to protobuf files
PROTO_PATH := $(CURDIR)/api

# Path to generated .pb.go files
PKG_PROTO_PATH := $(CURDIR)/gen/go

# Path to vendored protobuf files
VENDOR_PROTO_PATH := $(CURDIR)/vendor.protobuf

# Install necessary plugins
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/yoheimuta/protolint/cmd/protolint@latest

generate: export GOBIN := $(LOCAL_BIN)
generate:
	PATH=$(LOCAL_BIN):PATH $(LOCAL_BIN)/buf generate

# Linter
lint: .proto-lint

# Linter proto files
.proto-lint:
	$(LOCAL_BIN)/protolint -config_path ./.protolint.yaml ./proto/

# Format protobuf files
proto-format:
	$(info run buf format...)
	$(LOCAL_BIN)/buf format -w
	
# Declare that the current commands are not files and
# instrument Makefile to not search for changes in the file system
.PHONY: \
	.bin-deps \
	.protoc-generate \
	.tidy \
	.proto-lint \
	proto-format \
	vendor \
	generate \
	build \
	lint