// package analytics no longer provides any functionality. It is kept for backwards compatibility.
package analytics

import (
	"context"

	"github.com/giantswarm/microerror"
	"go.uber.org/zap"
)

type Config struct {
	Log *zap.SugaredLogger

	CredentialsJSON string
	Environment     string
}

type Analytics struct {
	log *zap.SugaredLogger

	environment string
}

func New(config Config) (*Analytics, error) {
	if config.Log == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Log must not be empty", config)
	}

	a := &Analytics{
		log:         config.Log,
		environment: config.Environment,
	}

	return a, nil
}

func (a *Analytics) Report(ctx context.Context, e Event) (Event, error) {
	a.log.Debugf("received event for session %s on app %s, doing nothing with it", e.SessionID, e.AppID)
	return e, nil
}
