package web

type UserFindByIdRequest struct {
	Id int `validate:"required" json:"id"`
}
