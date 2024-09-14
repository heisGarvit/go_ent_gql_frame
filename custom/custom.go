package custom

import (
	"context"
	"project/ent"
)

type GqlResolver struct {
	client *ent.Client
}

type Ping struct {
	Success bool
}

func (r *GqlResolver) Ping(ctx context.Context) (*Ping, error) {
	return &Ping{Success: true}, nil
}
