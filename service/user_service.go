package service

import (
	"context"

	"github.com/vaelvex/Go-API/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userID int)
	FindById(ctx context.Context, userID int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
