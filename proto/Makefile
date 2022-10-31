PROTOC_INC_PATH=/usr/local/include
GOOGLEAPIS_PATH=googleapis
VALIDATE_PATH=protoc-gen-validate
PROTOC_FLAGS:=-I. -I$(GOOGLEAPIS_PATH) -I$(VALIDATE_PATH)

.PHONY: generate
generate: helloworld iam pb

.PHONY: helloworld
helloworld:
	protoc $(PROTOC_FLAGS) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto

.PHONY: cr
cr:
	protoc $(PROTOC_FLAGS) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --validate_opt=paths=source_relative --validate_out="lang=go:." cr/cr.proto

.PHONY: iam
iam:
	protoc $(PROTOC_FLAGS) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative iam/iam.proto

.PHONY: pb
pb:
	protoc $(PROTOC_FLAGS) --include_imports --include_source_info --descriptor_set_out=proto.pb helloworld/helloworld.proto cr/cr.proto iam/iam.proto

.PHONY: openapi
openapi:
	protoc $(PROTOC_FLAGS) --openapiv2_out ./openapiv2 --openapiv2_opt logtostderr=true cr/cr.proto iam/iam.proto helloworld/helloworld.proto