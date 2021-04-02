package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/giantswarm/athena/pkg/graph/model"
)

func (r *queryResolver) Kubernetes(ctx context.Context) (*model.Kubernetes, error) {
	k := &model.Kubernetes{
		APIURL:  "",
		AuthURL: "",
		CaCert:  "",
	}

	return k, nil
}
