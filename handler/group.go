package handler

import (
	"im/common/res"
	"im/dao"
	"im/global"
	"im/log"
	"im/model"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GroupHead(c *gin.Context){
	file, err := os.Open("./head/group/"+c.Param("filename"))
	if err != nil {
		res.Err(c, "file not found")
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		res.Err(c, "read file err")
		return
	}
	c.Writer.Header().Add("Content-Type", "image/png")
	c.Status(http.StatusOK)
	c.Writer.Write(data)
}

func GroupRegist(c  *gin.Context) {
	groupName := c.PostForm("name")
	introduce := c.PostForm("introduce")
	masterId := c.GetUint("id")
	log.Info.Println(groupName, introduce, masterId)
	group := &model.Group{
		Name: groupName,
		MasterId: masterId,
		Introduce: introduce,
	}
	
	transaction := global.DB.Begin()

	if dao.CreateGroup(group).RowsAffected != 1 {
		transaction.Rollback()
		c.JSON(http.StatusInternalServerError, "创建失败")
		return
	}

	if dao.GetGroupByName(group).Error != nil {
		transaction.Rollback()
		c.JSON(http.StatusInternalServerError, "创建失败")
		return
	}

	session := model.GroupSession{
		GroupId: group.ID,
		UserId: masterId,
	}

	if dao.CreateGroupSession(&session).RowsAffected != 1 {
		transaction.Rollback()
		c.JSON(http.StatusInternalServerError, "创建失败")
		return
	}

	var filepath string
	var tmp *os.File
	var content []byte

	// 存储用户头像, 若出现错误,则用户使用默认头像
	file, _, err := c.Request.FormFile("head")
	if err != nil {
		goto END
	}
	defer file.Close()
	content, err = io.ReadAll(file)
	if err!=nil {
		goto END
	}

	filepath = "./head/group/" + groupName + ".png"

	tmp, err = os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0666)
	if err!=nil {
		goto END
	}
	defer tmp.Close()
	_, err = tmp.Write(content)
	if err != nil {
		os.Remove(filepath)
		goto END
	}
	END:
	c.JSON(http.StatusOK,  group.ID)
}