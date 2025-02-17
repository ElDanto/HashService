package hashService

//go:generate protoc --go_out=pkg/proto --go_opt=Mproto/hash.proto=github.com/ElDanto/ShortGen --go-grpc_out=pkg/proto --go-grpc_opt=Mproto/hash.proto=github.com/ElDanto/ShortGen --proto_path=proto proto/*.proto
