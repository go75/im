package initialize

import "im/ws"

func Init() {
	initConfig()
	initDB()
	initRedis()
	initLog()
	ws.InitDispatcher()
}