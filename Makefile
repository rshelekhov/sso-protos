.PHONY: default generate gen

default: generate

generate: ## Generate code from proto files
	protoc -I proto \
           proto/sso/*.proto \
           --go_out=./gen/go/ --go_opt=paths=source_relative \
           --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

gen: generate