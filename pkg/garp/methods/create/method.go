package create

import (
	"context"

	"github.com/10Narratives/dart/pkg/garp/resource"
)

type ResourceCreator[T resource.Resource] interface {
	CreateResource(ctx context.Context, opts *CreateResourceOptions) (T, error)
}
