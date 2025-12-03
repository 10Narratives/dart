package projectsrv

import (
	"context"

	projectdomain "github.com/10Narratives/dart/dart-gateway/internal/domain/project"
	"github.com/10Narratives/dart/pkg/garp/methods/create"
	"github.com/10Narratives/dart/pkg/garp/methods/delete"
	"github.com/10Narratives/dart/pkg/garp/methods/get"
	"github.com/10Narratives/dart/pkg/garp/methods/list"
	"github.com/10Narratives/dart/pkg/garp/methods/update"
)

type ProjectRepository interface {
	create.ResourceCreator[*projectdomain.Project]
	delete.ResourceDeleter[*projectdomain.Project]
	get.ResourceGetter[*projectdomain.Project]
	list.ResourceLister[*projectdomain.Project]
	update.ResourceUpdater[*projectdomain.Project]
}

type Service struct {
	projectRepository ProjectRepository
}

func NewService(projectRepository ProjectRepository) *Service {
	return &Service{projectRepository: projectRepository}
}

func (s *Service) CreateResource(ctx context.Context, opts *create.CreateResourceOptions) (*projectdomain.Project, error) {
	return s.projectRepository.CreateResource(ctx, opts)
}

func (s *Service) DeleteResource(ctx context.Context, opts *delete.DeleteResourceOptions) (*projectdomain.Project, error) {
	return s.projectRepository.DeleteResource(ctx, opts)
}

func (s *Service) GetResource(ctx context.Context, opts *get.GetResourceOptions) (*projectdomain.Project, error) {
	return s.projectRepository.GetResource(ctx, opts)
}

func (s *Service) ListResources(ctx context.Context, opts *list.ListResourceOptions) ([]*projectdomain.Project, string, error) {
	return s.projectRepository.ListResources(ctx, opts)
}

func (s *Service) UpdateResource(ctx context.Context, opts *update.UpdateResourceOptions) (*projectdomain.Project, error) {
	return s.projectRepository.UpdateResource(ctx, opts)
}
