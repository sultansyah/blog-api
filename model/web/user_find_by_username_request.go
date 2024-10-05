package web

type UserFindByUsernameRequest struct {
	Username int `validate:"required" json:"username"`
}
