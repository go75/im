package ws

import (
	"encoding/json"
	"im/global"
	"im/log"
	"im/ws/api"
	"im/ws/entity"
)

var channel = make(chan *entity.Request, 1024)

func InitDispatcher() {
	for i := 0; i < global.Config.Dispatcher.Size; i++ {
		go func() {
			var req *entity.Request
			//订阅消息, 不断阻塞等待请求队列的请求
			for req = range channel {
				bytes, err := json.Marshal(*req)
				if err != nil {
					log.Warn.Println("请求出错")
				}
				log.Info.Println("请求: ", string(bytes))
				api.Do(req)
			}
		}()
	}
}