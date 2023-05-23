package api

import (
	"im/log"
	"im/ws/entity"
	"strconv"
	"sync"
)

type routerMap struct {
	// key:消息id; value:该id对应的回调函数
	api map[uint]func(r *entity.Request)
	lock *sync.RWMutex
}

var router = &routerMap{
	api: make(map[uint]func(r *entity.Request)),
	lock: new(sync.RWMutex),
}

func Regist(id uint, fn func(r *entity.Request)) {
	router.lock.Lock()
	defer router.lock.Unlock()
	router.api[id] = fn
}

func get(id uint) func(r *entity.Request) {
	router.lock.RLock()
	defer router.lock.RUnlock()
	return router.api[id]
}

func Do(r *entity.Request) {
	//根据消息id获取绑定的handler方法
	fn :=  get(r.ID)

	if fn == nil {
		log.Warn.Println("api type:" + strconv.Itoa(int(r.ID)) + " is not regist")
		return
	}

	fn(r)
}