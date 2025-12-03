package pgcl

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Component struct {
	dsn  string
	conn *pgx.Conn
}

func NewComponent(dsn string, opts ...ComponentOption) *Component {
	co := defaultComponentOptions()
	for _, opt := range opts {
		opt(co)
	}

	return &Component{}
}

func (c *Component) Startup(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, c.dsn)
	if err != nil {
		return fmt.Errorf("cannot connect to postgres %s: %w", c.dsn, err)
	}
	c.conn = conn
	return nil
}

func (c *Component) Shutdown(ctx context.Context) error {
	if c.conn != nil {
		return c.conn.Close(ctx)
	}
	return nil
}

func (c *Component) Connection() *pgx.Conn {
	return c.conn
}
