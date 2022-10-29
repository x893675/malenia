PROTOC_INC_PATH=/usr/local/include
GOOGLEAPIS_PATH=googleapis
PROTOC_FLAGS:=-I. -I$(GOOGLEAPIS_PATH)

.PHONY: generate
generate: helloworld iam pb

.PHONY: helloworld
helloworld:
	protoc $(PROTOC_FLAGS) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto

.PHONY: cr
cr:
	protoc $(PROTOC_FLAGS) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative cr/cr.proto

.PHONY: iam
iam:
	protoc $(PROTOC_FLAGS) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative iam/iam.proto

.PHONY: pb
pb:
	protoc $(PROTOC_FLAGS) --include_imports --include_source_info --descriptor_set_out=proto.pb helloworld/helloworld.proto cr/cr.proto iam/iam.proto