package web

type UserFindByUsernameRequest struct {
	Username string `validate:"required" json:"username"`
}
