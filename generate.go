package hashService

//go:generate protoc --go_out=pkg/pb --go_opt=Mproto/hash.proto=github.com/ElDanto/HashService --go-grpc_out=pkg/pb --go-grpc_opt=Mproto/hash.proto=github.com/ElDanto/HashService --proto_path=proto proto/*.proto
