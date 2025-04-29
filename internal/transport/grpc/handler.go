package grpc

import (
	"context"
	userpb "github.com/Imnarka/project-protos/proto/user"
	"github.com/Imnarka/user-service/internal/users"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	userpb.UnimplementedUserServiceServer
	svc users.Service
}

func NewHandler(svc users.Service) userpb.UserServiceServer {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u, err := h.svc.CreateUser(req.Email)
	if err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	u, err := h.svc.GetUserByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.User{
		Id:    uint32(u.ID),
		Email: u.Email,
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	u, err := h.svc.UpdateUser(uint(req.Id), req.Email)
	if err != nil {
		return nil, err
	}
	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := h.svc.DeleteUser(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{Success: true}, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *emptypb.Empty) (*userpb.ListUsersResponse, error) {
	userList, err := h.svc.ListUsers()
	if err != nil {
		return nil, err
	}

	var pbUsers []*userpb.User
	for _, u := range userList {
		pbUsers = append(pbUsers, &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		})
	}

	return &userpb.ListUsersResponse{
		Users:      pbUsers,
		TotalCount: uint32(len(pbUsers)),
	}, nil
}
