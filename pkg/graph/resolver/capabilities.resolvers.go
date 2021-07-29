package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/giantswarm/athena/pkg/graph/exec"
	"github.com/giantswarm/athena/pkg/graph/model"
)

func (r *queryResolver) Capabilities(ctx context.Context) (*model.Capabilities, error) {
	k := &model.Capabilities{
		AvailabilityZones: r.Resolver.availabilityZones,
	}

	return k, nil
}

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() exec.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
