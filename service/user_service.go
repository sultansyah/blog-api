package service

import (
	"context"

	"github.com/sultansyah/blog-api/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
	FindById(ctx context.Context, request web.UserFindByIdRequest) web.UserResponse
	FindByUsername(ctx context.Context, request web.UserFindByUsernameRequest) web.UserResponse
}
