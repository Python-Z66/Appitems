package respones

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "message": msg})
}

// 成功
func Success(c *gin.Context, code int, msg string, data gin.H) {
	Response(c, http.StatusOK, code, data, msg)
}

// 失败
func Fail(c *gin.Context, code int, msg string, data gin.H) {
	Response(c, http.StatusBadRequest, code, data, msg)
}
