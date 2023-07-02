package handlers

import (
	"context"
	"errors"
	pb "github.com/antoineaudrain/go-task/workspace/api"
	"github.com/antoineaudrain/go-task/workspace/internal/app/services"
	"github.com/antoineaudrain/go-task/workspace/internal/domain/member"
	"github.com/antoineaudrain/go-task/workspace/internal/domain/workspace"
	customErrors "github.com/antoineaudrain/go-task/workspace/internal/shared/errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkspaceHandler struct {
	*pb.UnimplementedWorkspaceServiceServer
	workspaceService *services.WorkspaceService
	memberService    *services.MemberService
	logger           *zap.Logger
}

func NewWorkspaceHandler(logger *zap.Logger, workspaceRepo workspace.Repository, memberRepo member.Repository) *WorkspaceHandler {
	workspaceService := services.NewWorkspaceService(workspaceRepo, memberRepo)
	memberService := services.NewMemberService(workspaceRepo, memberRepo)

	return &WorkspaceHandler{
		workspaceService: workspaceService,
		memberService:    memberService,
		logger:           logger,
	}
}

func (h *WorkspaceHandler) CreateWorkspace(ctx context.Context, req *pb.CreateWorkspaceRequest) (*pb.CreateWorkspaceResponse, error) {
	authenticatedUserID := ctx.Value("authenticatedUserID").(*uuid.UUID)
	if authenticatedUserID == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	name := req.GetName()

	w, err := h.workspaceService.CreateWorkspace(ctx, name, *authenticatedUserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create w")
	}

	return &pb.CreateWorkspaceResponse{
		Workspace: &pb.Workspace{
			Id:        w.ID.String(),
			Name:      w.Name,
			OwnerId:   w.OwnerID.String(),
			CreatedAt: w.CreatedAt.String(),
			UpdatedAt: w.UpdatedAt.String(),
		},
	}, nil
}

func (h *WorkspaceHandler) GetWorkspace(ctx context.Context, req *pb.GetWorkspaceRequest) (*pb.GetWorkspaceResponse, error) {
	authenticatedUserID := ctx.Value("authenticatedUserID").(*uuid.UUID)
	if authenticatedUserID == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	workspaceID, err := uuid.Parse(req.GetWorkspaceId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid w ID")
	}

	w, err := h.workspaceService.GetWorkspace(ctx, workspaceID)
	if err != nil {
		if errors.Is(err, customErrors.ErrWorkspaceNotFound) {
			return nil, status.Errorf(codes.NotFound, "Workspace not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get w")
	}

	return &pb.GetWorkspaceResponse{
		Workspace: &pb.Workspace{
			Id:        w.ID.String(),
			Name:      w.Name,
			OwnerId:   w.OwnerID.String(),
			CreatedAt: w.CreatedAt.String(),
			UpdatedAt: w.UpdatedAt.String(),
		},
	}, nil
}

func (h *WorkspaceHandler) UpdateWorkspace(ctx context.Context, req *pb.UpdateWorkspaceRequest) (*pb.UpdateWorkspaceResponse, error) {
	authenticatedUserID := ctx.Value("authenticatedUserID").(*uuid.UUID)
	if authenticatedUserID == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	workspaceID, err := uuid.Parse(req.GetWorkspaceId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid w ID")
	}
	name := req.GetName()

	w, err := h.workspaceService.UpdateWorkspace(ctx, workspaceID, name, *authenticatedUserID)
	if err != nil {
		if errors.Is(err, customErrors.ErrWorkspaceNotFound) {
			return nil, status.Errorf(codes.NotFound, "Workspace not found")
		}
		if errors.Is(err, customErrors.ErrUnauthorized) {
			return nil, status.Errorf(codes.PermissionDenied, "Unauthorized to update w")
		}
		return nil, status.Errorf(codes.Internal, "Failed to update w")
	}

	return &pb.UpdateWorkspaceResponse{
		Workspace: &pb.Workspace{
			Id:        w.ID.String(),
			Name:      w.Name,
			OwnerId:   w.OwnerID.String(),
			CreatedAt: w.CreatedAt.String(),
			UpdatedAt: w.UpdatedAt.String(),
		},
	}, nil
}

func (h *WorkspaceHandler) DeleteWorkspace(ctx context.Context, req *pb.DeleteWorkspaceRequest) (*pb.DeleteWorkspaceResponse, error) {
	authenticatedUserID := ctx.Value("authenticatedUserID").(*uuid.UUID)
	if authenticatedUserID == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	workspaceID, err := uuid.Parse(req.GetWorkspaceId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid workspace ID")
	}

	err = h.workspaceService.DeleteWorkspace(ctx, workspaceID, *authenticatedUserID)
	if err != nil {
		if errors.Is(err, customErrors.ErrWorkspaceNotFound) {
			return nil, status.Errorf(codes.NotFound, "Workspace not found")
		}
		if errors.Is(err, customErrors.ErrUnauthorized) {
			return nil, status.Errorf(codes.PermissionDenied, "Unauthorized to delete workspace")
		}
		return nil, status.Errorf(codes.Internal, "Failed to delete workspace")
	}

	return &pb.DeleteWorkspaceResponse{
		Success: true,
	}, nil
}

func (h *WorkspaceHandler) ListWorkspaceMembers(ctx context.Context, req *pb.ListWorkspaceMembersRequest) (*pb.ListWorkspaceMembersResponse, error) {
	authenticatedUserID := ctx.Value("authenticatedUserID").(*uuid.UUID)
	if authenticatedUserID == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	workspaceID, err := uuid.Parse(req.GetWorkspaceId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid workspace ID")
	}

	members, err := h.memberService.ListByWorkspaceID(ctx, workspaceID)
	if err != nil {
		if errors.Is(err, customErrors.ErrWorkspaceNotFound) {
			return nil, status.Errorf(codes.NotFound, "Workspace not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to list workspace members")
	}

	var pbMembers []*pb.WorkspaceMember
	for _, m := range members {
		pbMember := &pb.WorkspaceMember{
			Id:           m.ID.String(),
			WorkspaceId:  m.WorkspaceID.String(),
			UserId:       m.UserID.String(),
			Role:         string(m.Role),
			InvitationId: m.InvitationId.String(),
			CreatedAt:    m.CreatedAt.String(),
			UpdatedAt:    m.UpdatedAt.String(),
		}
		pbMembers = append(pbMembers, pbMember)
	}

	return &pb.ListWorkspaceMembersResponse{
		Members: pbMembers,
	}, nil
}

func (h *WorkspaceHandler) InviteWorkspaceMember(ctx context.Context, req *pb.InviteWorkspaceMemberRequest) (*pb.InviteWorkspaceMemberResponse, error) {
	authenticatedUserID := ctx.Value("authenticatedUserID").(*uuid.UUID)
	if authenticatedUserID == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	workspaceID, err := uuid.Parse(req.GetWorkspaceId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid workspace ID")
	}
	email := req.GetEmail()

	m, err := h.memberService.InviteWorkspaceMember(ctx, workspaceID, email, *authenticatedUserID)
	if err != nil {
		if errors.Is(err, customErrors.ErrWorkspaceNotFound) {
			return nil, status.Errorf(codes.NotFound, "Workspace not found")
		}
		if errors.Is(err, customErrors.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "User not found")
		}
		if errors.Is(err, customErrors.ErrUnauthorized) {
			return nil, status.Errorf(codes.PermissionDenied, "Unauthorized to invite members")
		}
		return nil, status.Errorf(codes.Internal, "Failed to invite workspace m")
	}

	pbMember := &pb.WorkspaceMember{
		Id:           m.ID.String(),
		WorkspaceId:  m.WorkspaceID.String(),
		UserId:       m.UserID.String(),
		Role:         string(m.Role),
		InvitationId: m.InvitationId.String(),
		CreatedAt:    m.CreatedAt.String(),
		UpdatedAt:    m.UpdatedAt.String(),
	}

	return &pb.InviteWorkspaceMemberResponse{
		Member: pbMember,
	}, nil
}

func (h *WorkspaceHandler) RemoveWorkspaceMember(ctx context.Context, req *pb.RemoveWorkspaceMemberRequest) (*pb.RemoveWorkspaceMemberResponse, error) {
	authenticatedUserID := ctx.Value("authenticatedUserID").(*uuid.UUID)
	if authenticatedUserID == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	workspaceID, err := uuid.Parse(req.GetWorkspaceId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid workspace ID")
	}
	memberID, err := uuid.Parse(req.GetMemberId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid member ID")
	}

	err = h.memberService.RemoveWorkspaceMember(ctx, workspaceID, memberID, *authenticatedUserID)
	if err != nil {
		if errors.Is(err, customErrors.ErrWorkspaceNotFound) {
			return nil, status.Errorf(codes.NotFound, "Workspace not found")
		}
		if errors.Is(err, customErrors.ErrMemberNotFound) {
			return nil, status.Errorf(codes.NotFound, "Member not found")
		}
		if errors.Is(err, customErrors.ErrUnauthorized) {
			return nil, status.Errorf(codes.PermissionDenied, "Unauthorized to remove members")
		}
		return nil, status.Errorf(codes.Internal, "Failed to remove workspace member")
	}

	return &pb.RemoveWorkspaceMemberResponse{
		Success: true,
	}, nil
}
