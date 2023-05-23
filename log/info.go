package log

import (
	"log"
	"os"
)

//重要信息
var Info *log.Logger = log.New(os.Stdout, "\u001B[1;34m[Info]:\u001B[0m", log.Ldate|log.Ltime|log.Lshortfile)