package list

import "github.com/10Narratives/dart/pkg/garp/resource"

type ListResourceOptions struct {
	parent       *resource.ResourceName
	pageSize     int
	pageToken    string
	orderBy      string
	filter       string
	allowMissing bool
}

type ListResourceOption func(rlo *ListResourceOptions)

func NewListResourceOptions(opts ...ListResourceOption) *ListResourceOptions {
	options := DefaultResourceListerOptions()
	for _, opt := range opts {
		opt(options)
	}
	return options
}

func DefaultResourceListerOptions() *ListResourceOptions {
	return &ListResourceOptions{}
}

func WithParent(parent *resource.ResourceName) ListResourceOption {
	return func(rlo *ListResourceOptions) {
		rlo.parent = parent
	}
}

func WithPageSize(pageSize int) ListResourceOption {
	const maxPageSize = 1000
	return func(rlo *ListResourceOptions) {
		if pageSize > maxPageSize {
			pageSize = maxPageSize
		}
		rlo.pageSize = pageSize
	}
}

func WithPageToken(pageToken string) ListResourceOption {
	return func(rlo *ListResourceOptions) {
		rlo.pageToken = pageToken
	}
}

func WithOrderBy(orderBy string) ListResourceOption {
	return func(rlo *ListResourceOptions) {
		rlo.orderBy = orderBy
	}
}

func WithFilter(filter string) ListResourceOption {
	return func(rlo *ListResourceOptions) {
		rlo.filter = filter
	}
}

func WithAllowMissing(allowMissing bool) ListResourceOption {
	return func(rlo *ListResourceOptions) {
		rlo.allowMissing = allowMissing
	}
}

func (rlo *ListResourceOptions) Parent() *resource.ResourceName {
	return rlo.parent
}

func (rlo *ListResourceOptions) PageSize() int {
	return rlo.pageSize
}

func (rlo *ListResourceOptions) PageToken() string {
	return rlo.pageToken
}

func (rlo *ListResourceOptions) OrderBy() string {
	return rlo.orderBy
}

func (rlo *ListResourceOptions) Filter() string {
	return rlo.filter
}

func (rlo *ListResourceOptions) AllowMissing() bool {
	return rlo.allowMissing
}
