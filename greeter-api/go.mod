module github.com/scaleway/greeter-api

go 1.19

replace github.com/scaleway/protobuf/go => ../go

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0
	github.com/scaleway/protobuf/go v0.0.0-00010101000000-000000000000
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.10
	google.golang.org/grpc v1.51.0
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/net v0.3.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221227171554-f9683d7f8bef // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
