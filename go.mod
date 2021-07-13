module github.com/giantswarm/athena

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/giantswarm/microerror v0.3.0
	github.com/google/go-cmp v0.5.6
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.8.0
	github.com/vektah/gqlparser/v2 v2.2.0
	go.uber.org/zap v1.17.0
)

replace (
	github.com/coreos/etcd => github.com/etcd-io/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
)
