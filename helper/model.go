package helper

import (
	"github.com/sultansyah/blog-api/model/domain"
	"github.com/sultansyah/blog-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}
