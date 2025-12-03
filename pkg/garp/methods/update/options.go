package update

import "github.com/10Narratives/dart/pkg/garp/resource"

type UpdateResourceOptions struct {
	res          resource.Resource
	paths        []string
	allowMissing bool
	etag         string
}

type UpdateResourceOption func(ruo *UpdateResourceOptions)

func NewUpdateResourceOptions(opts ...UpdateResourceOption) *UpdateResourceOptions {
	options := DefaultUpdateResourceOptions()
	for _, opt := range opts {
		opt(options)
	}
	return options
}

func DefaultUpdateResourceOptions() *UpdateResourceOptions {
	return &UpdateResourceOptions{}
}

func WithAllowMissing(allowMissing bool) UpdateResourceOption {
	return func(ruo *UpdateResourceOptions) {
		ruo.allowMissing = allowMissing
	}
}

func WithEtag(etag string) UpdateResourceOption {
	return func(ruo *UpdateResourceOptions) {
		ruo.etag = etag
	}
}

func WithResource(res resource.Resource) UpdateResourceOption {
	return func(ruo *UpdateResourceOptions) {
		ruo.res = res
	}
}

func WithUpdateMask(paths []string) UpdateResourceOption {
	return func(ruo *UpdateResourceOptions) {
		ruo.paths = make([]string, len(paths))
		copy(ruo.paths, paths)
	}
}

func (ruo *UpdateResourceOptions) Resource() resource.Resource {
	return ruo.res
}

func (ruo *UpdateResourceOptions) UpdateMask() []string {
	return ruo.paths
}

func (ruo *UpdateResourceOptions) AllowMissing() bool {
	return ruo.allowMissing
}

func (ruo *UpdateResourceOptions) Etag() string {
	return ruo.etag
}
