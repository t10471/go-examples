.PHONY: proto-gen
proto-gen: proto-gen-go proto-gen-python

.PHONY: proto-gen-go
proto-gen-go:
	rm -rf go/proto/*
	protoc --go_out=./go/proto/. --go_opt=module=github.com/t10471/go-examples/grpc-examples/go/proto \
	--go-grpc_out=./go/proto/. --go-grpc_opt=module=github.com/t10471/go-examples/grpc-examples/go/proto \
	./protos/*.proto

.PHONY: proto-gen-python
proto-gen-python:
	ls python/proto| grep -v -E "^__init__" | xargs rm -rf
	cd python && poetry run python -m grpc_tools.protoc -I../protos \
	--python_out=./proto/. \
	--grpc_python_out=./proto/. ../protos/*.proto
