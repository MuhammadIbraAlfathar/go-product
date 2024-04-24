package route

import (
	"github.com/MuhammadIbraAlfathar/go-product/config"
	"github.com/MuhammadIbraAlfathar/go-product/handler"
	"github.com/MuhammadIbraAlfathar/go-product/repository"
	"github.com/MuhammadIbraAlfathar/go-product/service"
	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup) {
	db := config.LoadDB()
	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(userRepository, db)
	userHandler := handler.NewAuthHandler(userService)

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.GET("/register", userHandler.FindUsers)

}
