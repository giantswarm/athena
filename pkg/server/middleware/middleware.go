package middleware

import (
	"github.com/giantswarm/microerror"
	"go.uber.org/zap"

	"github.com/giantswarm/athena/pkg/server/middleware/cors"
)

type Config struct {
	Log *zap.SugaredLogger

	AllowedOrigins []string
}

type Middleware struct {
	Cors *cors.Middleware

	log            *zap.SugaredLogger
	allowedOrigins []string
}

func New(config Config) (*Middleware, error) {
	if config.Log == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Log must not be empty", config)
	}
	if len(config.AllowedOrigins) < 1 {
		return nil, microerror.Maskf(invalidConfigError, "%T.AllowedOrigins must not be empty", config)
	}

	var err error

	var corsMiddleware *cors.Middleware
	{
		c := cors.Config{
			Log:            config.Log,
			AllowedOrigins: config.AllowedOrigins,
		}

		corsMiddleware, err = cors.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	m := &Middleware{
		Cors: corsMiddleware,

		log:            config.Log,
		allowedOrigins: config.AllowedOrigins,
	}

	return m, nil
}
