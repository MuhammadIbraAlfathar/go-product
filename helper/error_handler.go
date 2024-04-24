package helper

import (
	"github.com/MuhammadIbraAlfathar/go-product/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotfoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	}

	response := Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	c.JSON(statusCode, response)
}
