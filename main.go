package main

import (
	"embed"
	"im/initialize"
	"im/router"

	"github.com/gin-gonic/gin"
)

//go:embed front/dist/assets/*
var static embed.FS

//go:embed front/dist/index.html
var index []byte

//go:embed front/view/login.html
var login []byte

//go:embed front/view/regist.html
var regist []byte

//go:embed front/dist/favicon.ico
var favicon []byte

func main() {
	initialize.Init()
	r := gin.New()
	router.Router(r, static, index, login, regist, favicon)
	panic(r.Run(":9999"))
}