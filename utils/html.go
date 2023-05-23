package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HTML(c *gin.Context, data []byte) {
	c.Writer.Header().Add("Content-Type", "text/html")
	c.Status(http.StatusOK)
	c.Writer.Write(data)
}