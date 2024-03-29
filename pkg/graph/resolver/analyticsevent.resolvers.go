package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"encoding/json"
	"time"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/athena/pkg/analytics"
	"github.com/giantswarm/athena/pkg/graph/exec"
	"github.com/giantswarm/athena/pkg/graph/model"
)

// CreateAnalyticsEvent is the resolver for the createAnalyticsEvent field.
func (r *mutationResolver) CreateAnalyticsEvent(ctx context.Context, event model.AnalyticsEventInput) (*model.AnalyticsEvent, error) {
	ctx, cancel := context.WithTimeout(ctx, eventTimeout)
	defer cancel()

	var err error

	newEventPayload := make(map[string]interface{})
	{
		err = json.Unmarshal([]byte(event.Payload), &newEventPayload)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newEvent := analytics.Event{
		AppID:                analytics.AppID(event.AppID),
		SessionID:            event.SessionID,
		PayloadType:          event.PayloadType,
		PayloadSchemaVersion: event.PayloadSchemaVersion,
		Payload:              newEventPayload,
		URI:                  event.URI,
		Timestamp:            time.Now().UTC(),
		InstallationID:       r.installationCodename,
	}

	createdEvent, err := r.analytics.Report(ctx, newEvent)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	createdEventPayload, err := json.Marshal(createdEvent.Payload)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	eventReport := &model.AnalyticsEvent{
		AppID:                string(createdEvent.AppID),
		SessionID:            createdEvent.SessionID,
		PayloadType:          createdEvent.PayloadType,
		PayloadSchemaVersion: createdEvent.PayloadSchemaVersion,
		Payload:              string(createdEventPayload),
		URI:                  createdEvent.URI,
		Timestamp:            createdEvent.Timestamp.Format(dateFormat),
		InstallationID:       createdEvent.InstallationID,
		EnvironmentClass:     createdEvent.EnvironmentClass,
	}

	return eventReport, nil
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
const (
	dateFormat   = time.RFC3339
	eventTimeout = 2 * time.Second
)
