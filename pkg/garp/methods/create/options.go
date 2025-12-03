package create

import "github.com/10Narratives/dart/pkg/garp/resource"

type CreateResourceOptions struct {
	parent *resource.ResourceName
	id     string
	res    resource.Resource
}

type CreateResourceOption func(rco *CreateResourceOptions)

func NewCreateResourceOptions(opts ...CreateResourceOption) *CreateResourceOptions {
	options := DefaultCreateResourceOptions()
	for _, opt := range opts {
		opt(options)
	}
	return options
}

func DefaultCreateResourceOptions() *CreateResourceOptions {
	return &CreateResourceOptions{}
}

func WithResourceID(id string) CreateResourceOption {
	return func(rco *CreateResourceOptions) {
		rco.id = id
	}
}

func WithParent(parent *resource.ResourceName) CreateResourceOption {
	return func(rco *CreateResourceOptions) {
		rco.parent = parent
	}
}

func WithResource(res resource.Resource) CreateResourceOption {
	return func(rco *CreateResourceOptions) {
		rco.res = res
	}
}

func (co *CreateResourceOptions) ResourceID() string {
	return co.id
}

func (co *CreateResourceOptions) Parent() *resource.ResourceName {
	return co.parent
}

func (co *CreateResourceOptions) Resource() resource.Resource {
	return co.res
}
