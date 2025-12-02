package projectapi

import (
	"context"

	projectdomain "github.com/10Narratives/dart/dart-gateway/internal/domain/project"
	grpcsrv "github.com/10Narratives/dart/pkg/components/transport/grpc/server"
	"github.com/10Narratives/dart/pkg/dart/gateway/project/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type api struct {
	project.UnimplementedProjectServiceServer
	projectService ProjectService
}

func Registration(
	projectService ProjectService,
) grpcsrv.ServiceRegistration {
	return func(server *grpc.Server) {
		project.RegisterProjectServiceServer(server, &api{projectService: projectService})
	}
}

func (a *api) CreateProject(ctx context.Context, req *project.CreateProjectRequest) (*project.Project, error) {
	args := CreateProjectArgs{
		ProjectID:   req.GetProjectId(),
		DisplayName: req.GetProject().GetDisplayName(),
		Description: req.GetProject().GetDescription(),
	}

	created, err := a.projectService.CreateProject(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create new project instance: %w", err)
	}

	return convertToGrpc(created), nil
}

func (a *api) DeleteProject(ctx context.Context, req *project.DeleteProjectRequest) (*emptypb.Empty, error) {

	panic("unimplemented")
}

func (a *api) GetProject(ctx context.Context, req *project.GetProjectRequest) (*project.Project, error) {
	panic("unimplemented")
}

func (a *api) ListProjects(ctx context.Context, req *project.ListProjectsRequest) (*project.ListProjectsResponse, error) {
	panic("unimplemented")
}

func (a *api) UpdateProject(ctx context.Context, req *project.UpdateProjectRequest) (*project.Project, error) {
	panic("unimplemented")
}

func convertToGrpc(src *projectdomain.Project) *project.Project {
	if src == nil {
		return nil
	}
	return &project.Project{
		Name:        src.Name,
		DisplayName: src.DisplayName,
		Description: src.Description,
		CreateTime:  timestamppb.New(src.CreateTime),
		UpdateTime:  timestamppb.New(src.UpdateTime),
	}
}

func convertFromGrpc(src *project.Project) *projectdomain.Project {
	if src == nil {
		return nil
	}
	return &projectdomain.Project{
		Name:        src.GetName(),
		DisplayName: src.GetDisplayName(),
		Description: src.GetDescription(),
		CreateTime:  src.GetCreateTime().AsTime(),
		UpdateTime:  src.GetCreateTime().AsTime(),
	}
}
