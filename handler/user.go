package handler

import (
	"fmt"
	"im/common/res"
	resmodel "im/common/resModel"
	"im/dao"
	"im/global"
	"im/log"
	"im/model"
	"im/utils"
	"im/ws/manager"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UserHead(c *gin.Context){
	file, err := os.Open("./head/user/"+c.Param("filename"))
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

func UserLogin(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	var identity model.UserIdentity

	global.DB.Raw("select * from user_identity where name=?", name).Scan(&identity)
	//global.DB.Where("name=?", identity.Name).Find(identity)
	//dao.QueryUserIdentity(identity)
	if ok := utils.CheckPwd(pwd, identity.Salt, identity.Pwd); ok {
		
		user := &model.User{
			Name: name,
		}
		
		dao.GetUserByName(user)
		if user.ID == 0 {
			// 未查询到用户
			res.Err(c, "该用户不存在")
			return
		}
		if manager.Get(user.ID) != nil {
			// 当前用户已在线, 无法继续登录
			res.Err(c, "用户已登录")
			return
		}
		token, err := utils.GenerateToken(user.ID, name)
		if err != nil {
			log.Error.Println("generate token err: ", err)
			res.Err(c, "登录失败")
			return
		}

		log.Info.Println("generate token: ", token)
		data := resmodel.LoginData {
			ID: user.ID,
			Name: user.Name,
			Token: token,
		}
		res.OkWithData(c, "登录成功", data)
	} else {
		res.Err(c, "密码错误")
	}
}

func UserRegist(c *gin.Context) {
	pwd := c.PostForm("pwd")
	if pwd != c.PostForm("check") {
		res.Err(c, "两次密码不一致")
		log.Warn.Println("两次密码不一致")
		return
	}

	name := c.PostForm("name")
	headname := name + ".png"
	salt := fmt.Sprintf("%10d", rand.Int31())
	identity := &model.UserIdentity{
		Name: name,
		Pwd: utils.MakePwd(pwd, salt),
		Salt: salt,
	}
	user := &model.User {
		Name: identity.Name,
	}

	// 先将用户信息存入数据库
	transaction := global.DB.Begin()

	db := dao.CreateUserIdentity(identity)
	if db.RowsAffected != 1 {
		transaction.Rollback()
		res.Err(c, "用户标记信息新增失败")
		log.Warn.Println(*identity)
		log.Warn.Println("用户标记信息新增失败")
		return
	}

	db = dao.CreateUser(user)
	if db.RowsAffected != 1 {
		transaction.Rollback()
		res.Err(c, "用户新增失败")
		log.Warn.Println("用户新增失败")
		return
	}
	transaction.Commit()

	// 存储用户头像, 若出现错误,则用户使用默认头像
	file, _, err := c.Request.FormFile("head")
	if err != nil {
		log.Info.Println("头像获取失败: " + err.Error())
		res.Ok(c, "注册成功,头像获取失败,请登录后切换")
		return
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err!=nil {
		res.Ok(c, "注册成功,头像读取失败,请登录后切换")
		return
	}

	filepath := "./head/user/" + headname

	tmp, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0666)
	if err!=nil {
		res.Ok(c, "注册成功,头像创建失败,请登录后切换")
		return
	}
	defer tmp.Close()
	_, err = tmp.Write(content)
	if err != nil {
		os.Remove(filepath)
		res.Ok(c, "注册成功,头像内容填充是失败,请登录后切换")
		return
	}
	res.Ok(c, "注册成功")
}

func DeleteUser(c *gin.Context) {
	pwd := c.PostForm("pwd")

	if pwd != c.PostForm("check") {
		res.Err(c, "两次密码不一致")
		return
	}

	name := c.GetString("name")
	identity := &model.UserIdentity{
		Name: name,
	}
	dao.QueryUserIdentity(identity)
	if ok := utils.CheckPwd(pwd, identity.Salt, identity.Pwd); ok {
		user := &model.User{

		}
		db := dao.DeleteUser(user)
		if db.RowsAffected != 1 {
			res.Err(c, "删除失败")
			return
		}
		db = dao.DeleteUserIdentity(identity)
		if db.RowsAffected != 1 {
			res.Err(c, "删除失败")
			return
		}
	}
	res.Ok(c, "删除成功")
}

func UpdateUser(c *gin.Context) {
	user := &model.User {
		Introduce: c.PostForm("introduce"),
	}
	dao.UpdateUser(user)
	res.Ok(c, "修改成功")
}

func QueryUserByName(c *gin.Context) {
	name := c.Query("name")
	log.Info.Println(name)
	if name == "" {
		// name为空
		res.Err(c, "用户名称为空")
		return
	}
	user := &model.User{
		Name: name,
	}
	db := dao.GetUserByName(user)
	if db.RowsAffected != 1 {
		res.Err(c, "未查询到用户")
	} else {
		res.OkWithData(c, "查询成功", user)
	}
}