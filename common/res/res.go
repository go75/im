package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, msg string) {
	Res(c, http.StatusOK, 0, msg, nil)
}

func OkWithData(c *gin.Context, msg string, data any) {
	Res(c, http.StatusOK, 0, msg, data)
}

func Err(c *gin.Context, msg string) {
	Res(c, http.StatusOK, -1, msg, nil)
}

func ErrWithData(c *gin.Context, msg string, data any) {
	Res(c, http.StatusOK, -1, msg, data)
}

func Res(c *gin.Context, status, code int, msg string, data any) {
	c.JSON(status, gin.H {
		"code": code,
		"msg":  msg,
        "data": data,
	})
}

