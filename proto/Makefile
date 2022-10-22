PROTOC_INC_PATH=/usr/local/include
GOOGLEAPIS_PATH=googleapis
PROTOC_FLAGS:=-I. -I$(GOOGLEAPIS_PATH)

.PHONY: generate
generate: helloworld pb

.PHONY: helloworld
helloworld:
	protoc $(PROTOC_FLAGS) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto

.PHONY: pb
pb:
	protoc $(PROTOC_FLAGS) --include_imports --include_source_info --descriptor_set_out=proto.pb helloworld/helloworld.proto