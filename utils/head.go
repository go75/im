package utils

// import (
// 	"im/config"
// 	"im/global"
// 	"io"
// 	"os"
// 	"time"
// )

// func GetUserHead(name string) ([]byte, error) {
// 	data, err := global.Rd.Get(name).Bytes()
// 	if err == nil {
// 		return data, nil
// 	}

// 	file, err := os.Open("./head/name" + ".png")
// 	if err != nil {
// 		return data, err
// 	}
// 	data, err = io.ReadAll(file)
// 	if err != nil {
// 		return data, err
// 	}
// 	global.Rd.Set(name, data, time.Duration(config.Redis.ExpiredTime) * time.Second)
// 	return data, nil
// }