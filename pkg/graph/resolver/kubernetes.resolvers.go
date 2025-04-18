package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"

	"github.com/giantswarm/athena/pkg/graph/model"
)

// Kubernetes is the resolver for the kubernetes field.
func (r *queryResolver) Kubernetes(ctx context.Context) (*model.Kubernetes, error) {
	k := &model.Kubernetes{
		APIURL:  r.installationK8sApiUrl,
		AuthURL: r.installationK8sAuthUrl,
		CaCert:  r.installationK8sCaCert,
	}

	return k, nil
}
