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
)

type Repository struct {
}

var _ projectsrv.ProjectRepository = &Repository{}

func NewRepository() *Repository {
	return &Repository{}
}

// CreateResource implements [projectsrv.ProjectRepository].
func (r *Repository) CreateResource(ctx context.Context, opts *create.CreateResourceOptions) (*projectdomain.Project, error) {
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
