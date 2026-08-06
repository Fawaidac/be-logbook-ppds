package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
}

// Success mengirim response JSON standar: { "code": ..., "message": ..., "data": ... }
func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// SuccessWithToken khusus untuk login: { "code": ..., "message": ..., "data": ..., "token": ... }
func SuccessWithToken(c *gin.Context, code int, message string, data interface{}, token string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
		Token:   token,
	})
}

// Error mengirim response JSON error: { "code": ..., "message": ... }
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

