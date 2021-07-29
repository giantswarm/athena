package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/giantswarm/athena/pkg/graph/model"
)

func (r *queryResolver) Identity(ctx context.Context) (*model.Identity, error) {
	i := &model.Identity{
		Provider: r.Resolver.installationProvider,
		Codename: r.Resolver.installationCodename,
	}

	return i, nil
}
