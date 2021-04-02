package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/giantswarm/athena/pkg/graph/generated"
	"github.com/giantswarm/athena/pkg/graph/model"
)

func (r *queryResolver) Identity(ctx context.Context) (*model.Identity, error) {
	i := &model.Identity{
		Provider: "",
		Codename: "",
	}

	return i, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
