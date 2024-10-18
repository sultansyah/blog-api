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

func ToUserLoginResponse(user domain.User, accessToken string, refreshToken string) web.UserLoginResponse {
	return web.UserLoginResponse{
		Id:           user.Id,
		Username:     user.Username,
		CreatedAt:    user.CreatedAt,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
