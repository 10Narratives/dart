package delete

import (
	"context"

	"github.com/10Narratives/dart/pkg/garp/resource"
)

type ResourceDeleter[T resource.Resource] interface {
	DeleteResource(ctx context.Context, opts *DeleteResourceOptions) (T, error)
}
