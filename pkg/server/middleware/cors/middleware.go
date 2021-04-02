package cors

import (
	"net/http"

	"github.com/giantswarm/microerror"
	"go.uber.org/zap"
)

type Config struct {
	Log *zap.SugaredLogger

	AllowedOrigins []string
}

type Middleware struct {
	log *zap.SugaredLogger

	allowedOrigins []string
}

func New(config Config) (*Middleware, error) {
	if config.Log == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Log must not be empty", config)
	}
	if len(config.AllowedOrigins) < 1 {
		return nil, microerror.Maskf(invalidConfigError, "%T.AllowedOrigins must not be empty", config)
	}

	m := &Middleware{
		log:            config.Log,
		allowedOrigins: config.AllowedOrigins,
	}

	return m, nil
}

func (m *Middleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer next.ServeHTTP(w, r)

		origin := string(r.Header.Get("Origin"))
		if !m.isOriginAllowed(m.allowedOrigins, origin) {
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", origin)

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Cache-Control, Content-Type, DNT, If-Modified-Since, Keep-Alive, User-Agent, X-Request-ID, X-Requested-With")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	})
}

func (m *Middleware) isOriginAllowed(allowed []string, origin string) bool {
	for _, v := range allowed {
		if v == origin || v == "*" {
			return true
		}
	}

	return false
}
