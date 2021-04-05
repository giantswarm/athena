package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/giantswarm/microerror"
	"go.uber.org/zap"

	"github.com/giantswarm/athena/pkg/graph/generated"
	"github.com/giantswarm/athena/pkg/graph/resolvers"
	"github.com/giantswarm/athena/pkg/server/middleware"
)

type Config struct {
	Log *zap.SugaredLogger

	AllowedOrigins         []string
	ListenAddress          string
	InstallationProvider   string
	InstallationCodename   string
	InstallationK8sApiUrl  string
	InstallationK8sAuthUrl string
	InstallationK8sCaCert  string
}

type Server struct {
	log *zap.SugaredLogger

	allowedOrigins         []string
	listenAddress          string
	installationProvider   string
	installationCodename   string
	installationK8sApiUrl  string
	installationK8sAuthUrl string
	installationK8sCaCert  string
}

func New(config Config) (*Server, error) {
	if config.Log == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Log must not be empty", config)
	}
	if len(config.AllowedOrigins) < 1 {
		return nil, microerror.Maskf(invalidConfigError, "%T.AllowedOrigins must not be empty", config)
	}
	if len(config.ListenAddress) < 1 {
		return nil, microerror.Maskf(invalidConfigError, "%T.ListenAddress must not be empty", config)
	}

	s := &Server{
		log:                    config.Log,
		allowedOrigins:         config.AllowedOrigins,
		listenAddress:          config.ListenAddress,
		installationProvider:   config.InstallationProvider,
		installationCodename:   config.InstallationCodename,
		installationK8sApiUrl:  config.InstallationK8sApiUrl,
		installationK8sAuthUrl: config.InstallationK8sAuthUrl,
		installationK8sCaCert:  config.InstallationK8sCaCert,
	}

	return s, nil
}

func (s *Server) Boot() error {
	var err error

	var middlewareMap *middleware.Middleware
	{
		config := middleware.Config{
			Log:            s.log,
			AllowedOrigins: s.allowedOrigins,
		}
		middlewareMap, err = middleware.New(config)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var rootResolver *resolvers.Resolver
	{
		config := resolvers.ResolverConfig{
			Log:                    s.log,
			InstallationProvider:   s.installationProvider,
			InstallationCodename:   s.installationCodename,
			InstallationK8sApiUrl:  s.installationK8sApiUrl,
			InstallationK8sAuthUrl: s.installationK8sAuthUrl,
			InstallationK8sCaCert:  s.installationK8sCaCert,
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
		mux.Handle("/graphql", middlewareMap.Cors.Middleware(graphQLServer))
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
