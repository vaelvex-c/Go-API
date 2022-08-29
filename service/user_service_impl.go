package service

import (
	"context"
	"database/sql"
	"github.com/chwlr/golang-api/exception"
	"github.com/chwlr/golang-api/helper"
	"github.com/chwlr/golang-api/model/domain"
	"github.com/chwlr/golang-api/model/web"
	"github.com/chwlr/golang-api/repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	//validation
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		IdRole:   request.IdRole,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	user = service.UserRepository.Save(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	//validation
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(
		ctx,
		tx,
		request.Id,
	)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.IdRole = request.IdRole
	user.Name = request.Name
	user.Email = request.Email
	user.Password = request.Password

	user = service.UserRepository.Update(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userID int) {
	//start transaction and check if error
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(
		ctx,
		tx,
		userID,
	)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userID int) web.UserResponse {
	//start transaction and check if error
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	//start transaction and check if error
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	return helper.ToUserResponses(users)
}

