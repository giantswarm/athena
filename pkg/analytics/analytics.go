package analytics

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/giantswarm/microerror"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type Config struct {
	Log *zap.SugaredLogger

	CredentialsJSON string
	Environment     string
}

type Analytics struct {
	log             *zap.SugaredLogger
	firestoreClient *firestore.Client

	environment string
}

func New(config Config) (*Analytics, error) {
	if config.Log == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Log must not be empty", config)
	}
	if len(config.CredentialsJSON) < 1 {
		return nil, microerror.Maskf(invalidConfigError, "%T.CredentialsJSON must not be empty", config)
	}

	var firestoreClient *firestore.Client
	{
		ctx := context.Background()
		credentials := option.WithCredentialsJSON([]byte(config.CredentialsJSON))

		app, err := firebase.NewApp(ctx, nil, credentials)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		firestoreClient, err = app.Firestore(ctx)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	a := &Analytics{
		log:             config.Log,
		firestoreClient: firestoreClient,
		environment:     config.Environment,
	}

	return a, nil
}

func (a *Analytics) Report(ctx context.Context, e Event) (Event, error) {
	a.log.Debugf("received event for session %s on app %s", e.SessionID, e.AppID)

	a.log.Debugf("applying defaults to event for session %s on app %s", e.SessionID, e.AppID)

	e = a.applyDefaultsToEvent(e)

	a.log.Debugf("validating event for session %s on app %s", e.SessionID, e.AppID)

	err := a.validateEvent(e)
	if err != nil {
		return Event{}, microerror.Mask(err)
	}

	a.log.Debugf("writing event record for session %s on app %s", e.SessionID, e.AppID)

	_, _, err = a.firestoreClient.Collection(e.CollectionName()).Add(ctx, e)
	if err != nil {
		return Event{}, microerror.Mask(err)
	}

	a.log.Debugf("wrote event record for session %s on app %s successfully", e.SessionID, e.AppID)

	return e, nil
}

func (a *Analytics) validateEvent(e Event) error {
	if !e.AppID.IsValid() {
		return microerror.Maskf(validationError, "unknown app id")
	}

	return nil
}

func (a *Analytics) applyDefaultsToEvent(e Event) Event {
	if len(e.EnvironmentClass) < 1 {
		e.EnvironmentClass = a.environment
	}

	return e
}
