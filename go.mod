module github.com/unibrightio/baseledger

go 1.16

require (
	github.com/containerd/containerd v1.5.2 // indirect
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/docker/docker v20.10.6+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/go-co-op/gocron v1.6.2
	github.com/gogo/protobuf v1.3.3
	github.com/golang-migrate/migrate v3.5.4+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/kthomas/go.uuid v1.2.1-0.20190324131420-28d1fa77e9a4
	github.com/nats-io/nats.go v1.11.0
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/spm v0.0.0-20210524110815-6d7452d2dc4a
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v0.0.0-20200527211525-6c9e30c09db2 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
