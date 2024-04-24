package service

import (
	"database/sql"
	"github.com/MuhammadIbraAlfathar/go-product/dto"
	"github.com/MuhammadIbraAlfathar/go-product/entity"
	"github.com/MuhammadIbraAlfathar/go-product/helper"
	"github.com/MuhammadIbraAlfathar/go-product/repository"
	"log"
)

type UserService interface {
	RegisterUser(req *dto.RegisterRequest) error
	FindUsers() ([]dto.RegisterResponse, error)
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserServiceImpl(userRepository repository.UserRepository, DB *sql.DB) *userServiceImpl {
	return &userServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
	}
}

func (s userServiceImpl) RegisterUser(req *dto.RegisterRequest) error {

	// validate password confirmation
	if req.Password != req.PasswordConfirmation {
		return &helper.BadRequestError{
			Message: "Password not match",
		}
	}

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitOrRollback(tx)

	user := entity.Users{
		Name:     req.Name,
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.Password,
		Gender:   req.Gender,
		Address:  req.Address,
	}

	err = s.UserRepository.RegisterUser(tx, &user)

	if err != nil {
		panic(err)
	}

	return nil

}

func ToUserResponse(user entity.Users) dto.RegisterResponse {
	return dto.RegisterResponse{
		Name:     user.Name,
		UserName: user.UserName,
		Email:    user.Email,
		Gender:   user.Gender,
		Address:  user.Address,
	}
}

func ToUserResponses(users []entity.Users) []dto.RegisterResponse {
	var userResponse []dto.RegisterResponse

	for _, user := range users {
		userResponse = append(userResponse, ToUserResponse(user))
	}

	return userResponse
}

func (s userServiceImpl) FindUsers() ([]dto.RegisterResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		errorHelper := &helper.InternalServerError{
			Message: "Internal Server Error",
		}
		return nil, errorHelper
		log.Println("Erorr DB userService")
	}

	defer helper.CommitOrRollback(tx)

	findUsers, err := s.UserRepository.FindUsers(tx)

	//var registerResponse []dto.RegisterResponse
	//
	//for _, user := range users {
	//	registerResponse = append(registerResponse, ToUserResponse(user))
	//}

	return ToUserResponses(findUsers), err
}
