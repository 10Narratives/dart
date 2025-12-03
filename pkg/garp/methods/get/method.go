package get

import (
	"context"

	"github.com/10Narratives/dart/pkg/garp/resource"
)

type ResourceGetter[T resource.Resource] interface {
	GetResource(ctx context.Context, opts *GetResourceOptions) (T, error)
}
