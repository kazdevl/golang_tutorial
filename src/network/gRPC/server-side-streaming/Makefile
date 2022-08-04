.PHONY: proto-gen
proto-gen:
        protoc \
        --go_out=. \
        --go-grpc_out=. \
        proto/*.proto

MSG = Hello;
push-message:
        grpcurl -d '{"message": "${MSG}"}' --plaintext --proto proto/example.proto localhost:8080 pb.ExampleService/Push

get-count:
        grpcurl --plaintext --proto proto/example.proto localhost:8080 pb.ExampleService/GetCount