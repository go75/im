package middleware

import (
	"im/common/imerr"
	"im/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 检权中间件
func Jwt(c *gin.Context) {

	// 允许直接通过的路径
	switch c.Request.URL.Path {
	case "/login":
		return
	case "/regist":
		return
	}

	// 通过http cookie中的token解析来认证
	token, err := c.Cookie("im-token")
	if err != nil {
		c.JSON(http.StatusBadRequest, "token解析错误")
		c.Abort()
		return
	}

	// 未携带token直接返回
	if token == "" {
		c.JSON(http.StatusBadRequest, "未携带token")
		c.Abort()
		return
	}

	// 初始化一个JWT对象实例，并根据结构体方法来解析token
	// 解析token中包含的相关信息(有效载荷)
	claims, err := utils.ParseToken(token)

	if err != nil {
		// token过期
		if err == imerr.ExpiredTokenErr {
			c.JSON(http.StatusBadRequest, "token过期")
			c.Abort()
			return
		}

		// 其他错误
		c.JSON(http.StatusBadRequest, "身份认证失败")
		c.Abort()
		return
	}

	// 将解析后的有效载荷claims的数据写入gin.Context引用对象中
	c.Set("id", claims.ID)
	c.Set("name", claims.Name)
	
	c.Next()
}