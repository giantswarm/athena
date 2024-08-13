package server

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func New(es graphql.ExecutableSchema, enableIntrospection bool) *handler.Server {
	srv := handler.New(es)
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	if enableIntrospection {
		srv.Use(extension.Introspection{})
	}
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return srv
}
