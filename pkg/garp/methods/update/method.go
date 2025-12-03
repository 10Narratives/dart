package update

import (
	"context"

	"github.com/10Narratives/dart/pkg/garp/resource"
)

type ResourceUpdater[T resource.Resource] interface {
	UpdateResource(ctx context.Context, opts *UpdateResourceOptions) (T, error)
}
