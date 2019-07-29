GO_PROC = $(shell protoc)


GO_PROTO_BASE = ${GO_PROC} --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. account/account.proto 

# 编译account
#${GO_PROTO_BASE} account\/account.proto