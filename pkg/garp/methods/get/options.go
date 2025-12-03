package get

import "github.com/10Narratives/dart/pkg/garp/resource"

type GetResourceOptions struct {
	name        *resource.ResourceName
	showDeleted bool
}

type GetResourceOption func(rgo *GetResourceOptions)

func NewGetResourceOptions(opts ...GetResourceOption) *GetResourceOptions {
	options := DefaultGetResourceOptions()
	for _, opt := range opts {
		opt(options)
	}
	return options
}

func DefaultGetResourceOptions() *GetResourceOptions {
	return &GetResourceOptions{}
}

func WithShowDeleted(showDeleted bool) GetResourceOption {
	return func(rgo *GetResourceOptions) {
		rgo.showDeleted = showDeleted
	}
}

func WithResourceName(name *resource.ResourceName) GetResourceOption {
	return func(rgo *GetResourceOptions) {
		rgo.name = name
	}
}

func (rgo *GetResourceOptions) Name() *resource.ResourceName {
	return rgo.name
}

func (rgo *GetResourceOptions) ShowDeleted() bool {
	return rgo.showDeleted
}
