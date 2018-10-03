build:
	# protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/theapemachine/gopilot-api/user-service proto/user/user.proto
	protoc --plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go --plugin=protoc-gen-micro=$(GOPATH)/bin/protoc-gen-micro --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-service .

run:
	docker run -p 50051:50051 \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e MICRO_REGISTRY=mdns user-service
