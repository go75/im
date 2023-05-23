package handler

import (
	"im/log"
	"im/utils"
	"im/ws/entity"
	"math/rand"
	"time"
)

func FiveChess(r *entity.Request) {
	utils.Send(r.ProcessId, r.ID, entity.Text, r.SenderId, r.SenderName)
}

func AgreeFiveChess(r *entity.Request) {
	log.Info.Println("agree five chess request")
	rand.Seed(time.Now().Unix())
	switch rand.Int() % 2 {
	case 0:
		utils.SendStr(r.ProcessId, r.ID, entity.Text, r.SenderId, "1")
		utils.SendStr(r.SenderId, r.ID, entity.Text, r.ProcessId, "2")
	case 1:
		utils.SendStr(r.ProcessId, r.ID, entity.Text, r.SenderId, "2")
		utils.SendStr(r.SenderId, r.ID, entity.Text, r.ProcessId, "1")
	}
}

func RefuseFiveChess(r *entity.Request) {
	utils.Send(r.ProcessId, r.ID, entity.Text, r.ID, nil)
}

func DownChessPosition(r *entity.Request) {
	log.Info.Println("落子")
	utils.Send(r.ProcessId, r.ID, entity.Text, r.SenderId, r.Payload)
}