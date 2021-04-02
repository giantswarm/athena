package resolvers

import (
	"github.com/giantswarm/microerror"
	"go.uber.org/zap"
)

//go:generate rm -rf generated
//go:generate go run github.com/99designs/gqlgen

type ResolverConfig struct {
	Log *zap.SugaredLogger
}

type Resolver struct {
	log *zap.SugaredLogger
}

func NewResolver(config ResolverConfig) (*Resolver, error) {
	if config.Log == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Log must not be empty", config)
	}

	r := &Resolver{
		log: config.Log,
	}

	return r, nil
}
