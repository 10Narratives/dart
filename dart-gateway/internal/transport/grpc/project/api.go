package projectapi

import (
	"context"

	projectdomain "github.com/10Narratives/dart/dart-gateway/internal/domain/project"
	grpcsrv "github.com/10Narratives/dart/pkg/components/transport/grpc/server"
	"github.com/10Narratives/dart/pkg/dart/gateway/project/v1"
	"github.com/10Narratives/dart/pkg/garp/methods/create"
	"github.com/10Narratives/dart/pkg/garp/methods/delete"
	"github.com/10Narratives/dart/pkg/garp/methods/get"
	"github.com/10Narratives/dart/pkg/garp/methods/list"
	"github.com/10Narratives/dart/pkg/garp/methods/update"
	"github.com/10Narratives/dart/pkg/garp/resource"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProjectService interface {
	create.ResourceCreator[*projectdomain.Project]
	delete.ResourceDeleter[*projectdomain.Project]
	get.ResourceGetter[*projectdomain.Project]
	list.ResourceLister[*projectdomain.Project]
	update.ResourceUpdater[*projectdomain.Project]
}

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
	opts := create.NewCreateResourceOptions(
		create.WithResourceID(req.GetProjectId()),
		create.WithResource(projectFromProto(req.GetProject())),
	)

	created, err := a.projectService.CreateResource(ctx, opts)
	if err != nil {
		return nil, errorToProto(err)
	}

	return projectToProto(created), nil
}

func (a *api) DeleteProject(ctx context.Context, req *project.DeleteProjectRequest) (*emptypb.Empty, error) {
	name, err := resource.NewResourceName(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected name format: %v", err)
	}

	opts := delete.NewDeleteResourceOptions(
		delete.WithResourceName(name),
	)

	if _, err := a.projectService.DeleteResource(ctx, opts); err != nil {
		return nil, errorToProto(err)
	}

	return &emptypb.Empty{}, nil
}

func (a *api) GetProject(ctx context.Context, req *project.GetProjectRequest) (*project.Project, error) {
	name, err := resource.NewResourceName(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected name format: %v", err)
	}

	opts := get.NewGetResourceOptions(
		get.WithResourceName(name),
	)

	retrieved, err := a.projectService.GetResource(ctx, opts)
	if err != nil {
		return nil, errorToProto(err)
	}

	return projectToProto(retrieved), nil
}

func (a *api) ListProjects(ctx context.Context, req *project.ListProjectsRequest) (*project.ListProjectsResponse, error) {
	opts := list.NewListResourceOptions(
		list.WithPageSize(int(req.GetPageSize())),
		list.WithPageToken(req.GetPageToken()),
	)

	listed, token, err := a.projectService.ListResources(ctx, opts)
	if err != nil {
		return nil, errorToProto(err)
	}

	converted := make([]*project.Project, 0, len(listed))
	for _, project := range listed {
		converted = append(converted, projectToProto(project))
	}

	return &project.ListProjectsResponse{
		Projects:      converted,
		NextPageToken: token,
	}, nil
}

func (a *api) UpdateProject(ctx context.Context, req *project.UpdateProjectRequest) (*project.Project, error) {
	opts := update.NewUpdateResourceOptions(
		update.WithResource(projectFromProto(req.GetProject())),
		update.WithUpdateMask(req.GetUpdateMask().GetPaths()),
	)

	updated, err := a.projectService.UpdateResource(ctx, opts)
	if err != nil {
		return nil, errorToProto(err)
	}

	return projectToProto(updated), nil
}

func projectFromProto(src *project.Project) *projectdomain.Project {
	if src == nil {
		return nil
	}

	name, err := resource.NewResourceName(src.GetName())
	if err != nil {
		return nil
	}

	return projectdomain.NewProject(name,
		src.GetDisplayName(),
		src.GetDescription(),
		src.GetCreateTime().AsTime(),
		src.GetUpdateTime().AsTime(),
	)
}

func projectToProto(src *projectdomain.Project) *project.Project {
	if src == nil {
		return nil
	}

	return &project.Project{
		Name:        src.ResourceName().Value(),
		DisplayName: src.DisplayName(),
		Description: src.Description(),
		CreateTime:  timestamppb.New(src.CreateTime()),
		UpdateTime:  timestamppb.New(src.UpdateTime()),
	}
}

func errorToProto(err error) error {
	switch {
	default:
		return status.Error(codes.Internal, "cannot create project")
	}
}
