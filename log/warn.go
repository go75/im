package log

import (
	"log"
	"os"
)

//警告
var Warn *log.Logger = log.New(os.Stdout, "\u001B[1;33m[Warn]:\u001B[0m", log.Ldate|log.Ltime|log.Lshortfile)