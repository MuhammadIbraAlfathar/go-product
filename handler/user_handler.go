package handler

import (
	"github.com/MuhammadIbraAlfathar/go-product/dto"
	"github.com/MuhammadIbraAlfathar/go-product/helper"
	"github.com/MuhammadIbraAlfathar/go-product/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service service.UserService
}

func NewAuthHandler(s service.UserService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var register dto.RegisterRequest

	err := c.ShouldBindJSON(&register)
	if err != nil {
		helper.HandlerError(c, &helper.BadRequestError{
			Message: err.Error(),
		})
	}

	err = h.service.RegisterUser(&register)
	if err != nil {
		helper.HandlerError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register Successfully",
	})

	c.JSON(http.StatusCreated, res)
}

func (h *AuthHandler) FindUsers(c *gin.Context) {
	users, err := h.service.FindUsers()
	if err != nil {
		helper.HandlerError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Success get data",
		Data:       users,
	})

	c.JSON(http.StatusOK, res)
}
