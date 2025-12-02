package projectapi

import (
	"context"

	projectdomain "github.com/10Narratives/dart/dart-gateway/internal/domain/project"
)

type ProjectService interface {
	CreateProject(ctx context.Context, args CreateProjectArgs) (*projectdomain.Project, error)
}

type CreateProjectArgs struct {
	ProjectID   string
	DisplayName string
	Description string
}
