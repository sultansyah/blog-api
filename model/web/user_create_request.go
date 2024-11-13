package web

type UserRegisterRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required,max=200,min=1" json:"username"`
	Password string `validate:"required,max=200,min=1" json:"password"`
}
