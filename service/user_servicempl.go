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
	LoginUser(req *dto.LoginRequest) (*dto.LoginResponse, error)
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

	hashPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return &helper.InternalServerError{
			Message: "Internal Server Error",
		}
	}

	defer helper.CommitOrRollback(tx)

	user := entity.Users{
		Name:     req.Name,
		UserName: req.UserName,
		Email:    req.Email,
		Password: hashPassword,
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
		log.Println("Erorr DB userService")
		errorHelper := &helper.InternalServerError{
			Message: "Internal Server Error",
		}
		return nil, errorHelper
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

func (s userServiceImpl) LoginUser(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	tx, err := s.DB.Begin()
	if err != nil {
		log.Println("Erorr DB userService")
		errorHelper := &helper.InternalServerError{
			Message: "Internal Server Error",
		}
		return nil, errorHelper
	}

	defer helper.CommitOrRollback(tx)

	//cek email user
	user, err := s.UserRepository.FindUserByEmail(tx, req.Email)
	if err != nil {
		log.Println("Salah email")
		return nil, &helper.NotfoundError{
			Message: "Wrong email or password",
		}
	}

	//verifikasi password
	err = helper.ValidatePassword(user.Password, req.Password)
	if err != nil {
		log.Println("Salah password")
		return nil, &helper.InternalServerError{
			Message: "Wrong email or password",
		}
	}

	data = dto.LoginResponse{
		ID:       user.Id,
		UserName: user.UserName,
		Name:     user.Name,
	}

	return &data, nil

}
