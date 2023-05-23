package initialize

import (
	"im/global"
	mylog "im/log"
	"io"
	"log"
	"os"
)

func initLog() {
	file, err := os.OpenFile(global.Config.Log.Location, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("无法打开错误的位于" + global.Config.Log.Location + "log文件: " + err.Error())
	}
	mylog.Error = log.New(io.MultiWriter(file, os.Stderr), "\u001B[1;31m[Error]:\u001B[0m", log.Ldate|log.Ltime|log.Lshortfile)
}