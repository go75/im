package initialize

import (
	"encoding/json"
	"im/global"
	"im/log"
	"io"
	"os"
)

func initConfig() {
	file, err := os.Open("./config.json")
	if err !=nil {
		panic(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, global.Config)
	if err != nil {
		panic(err)
	}
	log.Info.Println("config: ", *global.Config)
}