package utils

import (
	"encoding/json"
	"im/ws/entity"
	"im/ws/manager"

	"github.com/gorilla/websocket"
)

func Send(connID, id, option, processId uint, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	obj := entity.Obj {
		ID: id,
		Type: option,
		ProcessId: processId,
		Payload: string(data),
	}
	data, err = json.Marshal(obj)
	if err != nil {
		return err
	}
	err  = manager.Send(connID, data)
	return err
}

func SendStr(connID, id, option, processId uint, payload string) error {
	conn := manager.Get(connID)
	obj := entity.Obj {
		ID: id,
		Type: option,
		ProcessId: processId,
		Payload: payload,
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}
	return nil
}

// func Publish(channel string, id, option, processId uint, payload any) error {
// 	data, err := json.Marshal(payload)
// 	if err != nil {
// 		return err
// 	}
// 	obj := entity.Obj{
// 		ID: id,
// 		Type: option,
// 		ProcessId: processId,
// 		Payload: string(data),
// 	}
// 	data, err = json.Marshal(obj)
// 	if err != nil {
// 		return err
// 	}
// 	global.Rd.Publish(channel, string(data))
// 	return nil
// }