package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/giantswarm/athena/pkg/graph/generated"
	"github.com/giantswarm/athena/pkg/graph/resolvers"
	"github.com/giantswarm/microerror"
	"go.uber.org/zap"
)

type Config struct {
	Log *zap.SugaredLogger

	ListenAddress string
}

type Server struct {
	log *zap.SugaredLogger

	listenAddress string
}

func New(config Config) (*Server, error) {
	if config.Log == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Log must not be empty", config)
	}
	if len(config.ListenAddress) < 1 {
		return nil, microerror.Maskf(invalidConfigError, "%T.ListenAddress must not be empty", config)
	}

	s := &Server{
		log:           config.Log,
		listenAddress: config.ListenAddress,
	}

	return s, nil
}

func (s *Server) Boot() error {
	var err error

	var rootResolver *resolvers.Resolver
	{
		config := resolvers.ResolverConfig{
			Log: s.log,
		}
		rootResolver, err = resolvers.NewResolver(config)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var graphQLServer *handler.Server
	{
		schema := generated.NewExecutableSchema(generated.Config{
			Resolvers: rootResolver,
		})
		graphQLServer = handler.NewDefaultServer(schema)
	}

	mux := http.NewServeMux()
	{
		mux.Handle("/graphql", graphQLServer)
		mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	}

	server := &http.Server{
		Addr:    s.listenAddress,
		Handler: mux,
	}

	s.log.Infof("server listening on address %s", s.listenAddress)
	s.log.Infof("visit %s for the GraphQL playground", s.listenAddress)

	err = server.ListenAndServe()
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
