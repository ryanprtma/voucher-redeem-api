package commons

import "github.com/gin-gonic/gin"

type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func WriteSuccess(ctx *gin.Context, data interface{}, statusCode int) {
	resp := ResponseSuccess{
		Success: true,
		Data:    data,
	}

	ctx.JSON(statusCode, resp)
}

func WriteError(ctx *gin.Context, message string, statusCode int) {
	resp := ResponseError{
		Success: false,
		Message: message,
	}

	ctx.JSON(statusCode, resp)
}
