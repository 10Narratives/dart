package grpcsrv

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type ServiceRegistration func(s *grpc.Server)

type Component struct {
	address string
	server  *grpc.Server
}

func NewComponent(address string, opts ...ComponentOption) *Component {
	options := defaultComponentOptions()
	for _, opt := range opts {
		opt(options)
	}

	server := grpc.NewServer(options.serverOptions...)
	for _, reg := range options.regs {
		reg(server)
	}

	return &Component{
		address: address,
		server:  server,
	}
}

func (c *Component) Startup(ctx context.Context) error {
	lis, err := net.Listen("tcp", c.address)
	if err != nil {
		return fmt.Errorf("cannot startup grpc server component: %w", err)
	}

	channel := make(chan error)

	go func() {
		channel <- c.server.Serve(lis)
		close(channel)
	}()

	select {
	case err := <-channel:
		return fmt.Errorf("error during tcp listening on address %s: %w", c.address, err)
	case <-ctx.Done():
		return nil
	}
}

func (c *Component) Shutdown(ctx context.Context) error {
	channel := make(chan struct{})

	go func() {
		c.server.GracefulStop()
		channel <- struct{}{}
		close(channel)
	}()

	select {
	case <-channel:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
