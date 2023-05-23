package router

import (
	"im/handler"

	"github.com/gin-gonic/gin"
)

func user(r *gin.Engine) {
	r.POST("/login", handler.UserLogin)
	r.POST("/regist", handler.UserRegist)
	r.GET("/head/:filename", handler.UserHead)
}