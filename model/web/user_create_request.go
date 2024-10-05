package web

type UserCreateRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required,max=200,min=1" json:"username"`
	Password string `validate:"required,max=200,min=1" json:"password"`
}
