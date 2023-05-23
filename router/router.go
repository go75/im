package router

import (
	"embed"
	"im/middleware"
	"im/ws"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, static embed.FS, index, login, regist, favicon []byte) {
	r.Use(middleware.Cors)
	staticFiles, _ := fs.Sub(static, "front/dist/assets")
	r.StaticFS("/assets", http.FS(staticFiles))
	r.Use(middleware.RateLimiter, middleware.Jwt)
	
	r.GET("/login", func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "text/html")
		c.Status(http.StatusOK)
		c.Writer.Write(login)
	})
	r.GET("/regist", func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "text/html")
		c.Status(http.StatusOK)
		c.Writer.Write(regist)
	})
	r.GET("/", func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "text/html")
		c.Status(http.StatusOK)
		c.Writer.Write(index)
	})
	r.GET("/favicon.ico", func(c *gin.Context){
		c.Writer.Header().Add("Content-Type", "image/x-icon")
		c.Status(http.StatusOK)
		c.Writer.Write(favicon)
	})

	ws.Router("/upgrade", r)

	user(r)
	group(r.Group("group"))
}
