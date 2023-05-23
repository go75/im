package router

import (
	"im/handler"

	"github.com/gin-gonic/gin"
)

func group(r *gin.RouterGroup) {
	r.GET("/head/:filename", handler.GroupHead)
	r.POST("/regist", handler.GroupRegist)
}