package initialize

import (
	"im/global"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDB() {
	// sql日志
	logger := logger.New (
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config {
			SlowThreshold: time.Duration(global.Config.Database.SlowThreshold) * time.Millisecond, // 慢SQL阈值
			//LogLevel: logger.Info, // 级别
			LogLevel: logger.LogLevel(global.Config.Database.LogLevel),
			Colorful: global.Config.Database.Colorful, // 彩色
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(global.Config.Database.DSN), &gorm.Config {
		Logger: logger,
	})
	if err!=nil {
		panic(err)
	}
}