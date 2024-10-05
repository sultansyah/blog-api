package web

import "time"

type UserLoginResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}
