package web

type UserLoginRequest struct {
	Username string `validate:"required,max=200,min=1" json:"username"`
	Password string `validate:"required,max=200,min=1" json:"password"`
}
