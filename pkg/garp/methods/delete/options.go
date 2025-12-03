package delete

import "github.com/10Narratives/dart/pkg/garp/resource"

type DeleteResourceOptions struct {
	soft         bool
	force        bool
	etag         string
	allowMissing bool
	resourceName *resource.ResourceName
}

type DeleteResourceOption func(rdo *DeleteResourceOptions)

func NewDeleteResourceOptions(opts ...DeleteResourceOption) *DeleteResourceOptions {
	options := DefaultDeleteResourceOptions()
	for _, opt := range opts {
		opt(options)
	}
	return options
}

func DefaultDeleteResourceOptions() *DeleteResourceOptions {
	return &DeleteResourceOptions{}
}

func WithSoft(soft bool) DeleteResourceOption {
	return func(rdo *DeleteResourceOptions) {
		rdo.soft = soft
	}
}

func WithForce(force bool) DeleteResourceOption {
	return func(rdo *DeleteResourceOptions) {
		rdo.force = force
	}
}

func WithEtag(etag string) DeleteResourceOption {
	return func(rdo *DeleteResourceOptions) {
		rdo.etag = etag
	}
}

func WithAllowMissing(allowMissing bool) DeleteResourceOption {
	return func(rdo *DeleteResourceOptions) {
		rdo.allowMissing = allowMissing
	}
}

func WithResourceName(name *resource.ResourceName) DeleteResourceOption {
	return func(rdo *DeleteResourceOptions) {
		rdo.resourceName = name
	}
}

func (rdo *DeleteResourceOptions) Soft() bool {
	return rdo.soft
}

func (rdo *DeleteResourceOptions) Force() bool {
	return rdo.force
}

func (rdo *DeleteResourceOptions) Etag() string {
	return rdo.etag
}

func (rdo *DeleteResourceOptions) AllowMissing() bool {
	return rdo.allowMissing
}

func (rdo *DeleteResourceOptions) ResourceName() *resource.ResourceName {
	return rdo.resourceName
}
