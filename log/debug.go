package log

import (
	"log"
	"os"
)

//调试信息
var Debug *log.Logger = log.New(os.Stdout, "\u001B[1;36m[Debug]:\u001B[0m", log.Ltime|log.Llongfile)
