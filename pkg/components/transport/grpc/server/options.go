package grpcsrv

import "google.golang.org/grpc"

type componentOptions struct {
	regs          []ServiceRegistration
	serverOptions []grpc.ServerOption
}

type ComponentOption func(co *componentOptions)

func defaultComponentOptions() *componentOptions {
	return &componentOptions{
		regs:          make([]ServiceRegistration, 0),
		serverOptions: make([]grpc.ServerOption, 0),
	}
}

func WithServiceRegistrations(regs []ServiceRegistration) ComponentOption {
	return func(co *componentOptions) {
		co.regs = append(co.regs, regs...)
	}
}

func WithServerOptions(serverOptions []grpc.ServerOption) ComponentOption {
	return func(co *componentOptions) {
		co.serverOptions = append(co.serverOptions, serverOptions...)
	}
}
