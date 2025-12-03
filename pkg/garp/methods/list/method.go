package list

import (
	"context"

	"github.com/10Narratives/dart/pkg/garp/resource"
)

type ResourceLister[T resource.Resource] interface {
	ListResources(ctx context.Context, opts *ListResourceOptions) ([]T, string, error)
}
