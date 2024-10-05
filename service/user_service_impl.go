package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/sultansyah/blog-api/helper"
	"github.com/sultansyah/blog-api/model/domain"
	"github.com/sultansyah/blog-api/model/web"
	"github.com/sultansyah/blog-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository repository.UserRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func newUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		repository: userRepository,
		DB:         DB,
		Validate:   validate,
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (service UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitOrRollback(tx)

	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		panic(err)
	}

	user := domain.User{
		Email:        request.Email,
		Username:     request.Username,
		PasswordHash: hashedPassword,
	}

	user = service.repository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitOrRollback(tx)

	// user := service.repository.FindById(ctx, tx, request.Username)

	// isValidPassword := CheckPasswordHash(request.Password, )
}
func (service UserServiceImpl) FindById(ctx context.Context, request web.UserFindByIdRequest) web.UserResponse {

}
