module github.com/giantswarm/athena

go 1.16

require (
	cloud.google.com/go/firestore v1.6.1
	cloud.google.com/go/iam v0.2.0 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/99designs/gqlgen v0.17.1
	github.com/giantswarm/microerror v0.4.0
	github.com/google/go-cmp v0.5.7
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.1
	github.com/vektah/gqlparser/v2 v2.4.0
	go.uber.org/zap v1.21.0
	google.golang.org/api v0.70.0
)

replace (
	github.com/coreos/etcd => github.com/etcd-io/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
)
