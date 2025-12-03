package projectrepo

import (
	"context"

	projectdomain "github.com/10Narratives/dart/dart-gateway/internal/domain/project"
	projectsrv "github.com/10Narratives/dart/dart-gateway/internal/services/project"
	"github.com/10Narratives/dart/pkg/garp/methods/create"
	"github.com/10Narratives/dart/pkg/garp/methods/delete"
	"github.com/10Narratives/dart/pkg/garp/methods/get"
	"github.com/10Narratives/dart/pkg/garp/methods/list"
	"github.com/10Narratives/dart/pkg/garp/methods/update"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

var _ projectsrv.ProjectRepository = &Repository{}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		db: conn,
	}
}

// CreateResource implements [projectsrv.ProjectRepository].
func (r *Repository) CreateResource(ctx context.Context, opts *create.CreateResourceOptions) (*projectdomain.Project, error) {
	const query = `select * from code.create_project();`
	r.db.Query(ctx, query)

	panic("unimplemented")
}

// DeleteResource implements [projectsrv.ProjectRepository].
func (r *Repository) DeleteResource(ctx context.Context, opts *delete.DeleteResourceOptions) (*projectdomain.Project, error) {
	panic("unimplemented")
}

// GetResource implements [projectsrv.ProjectRepository].
func (r *Repository) GetResource(ctx context.Context, opts *get.GetResourceOptions) (*projectdomain.Project, error) {
	panic("unimplemented")
}

// ListResources implements [projectsrv.ProjectRepository].
func (r *Repository) ListResources(ctx context.Context, opts *list.ListResourceOptions) ([]*projectdomain.Project, string, error) {
	panic("unimplemented")
}

// UpdateResource implements [projectsrv.ProjectRepository].
func (r *Repository) UpdateResource(ctx context.Context, opts *update.UpdateResourceOptions) (*projectdomain.Project, error) {
	panic("unimplemented")
}
