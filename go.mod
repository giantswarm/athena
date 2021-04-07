module github.com/giantswarm/athena

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/giantswarm/microerror v0.3.0
	github.com/google/go-cmp v0.5.5
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/vektah/gqlparser/v2 v2.1.0
	go.uber.org/zap v1.16.0
)

replace (
	github.com/coreos/etcd => github.com/etcd-io/etcd v3.3.25+incompatible
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
)
