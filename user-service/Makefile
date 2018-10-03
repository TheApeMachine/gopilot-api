build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/theapemachine/gopilot-api/user-service proto/user/user.proto
	# protoc --plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go --plugin=protoc-gen-micro=$(GOPATH)/bin/protoc-gen-micro --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. proto/user/user.proto
