package logic

import (
	user "example/gomicro/user_rpc/protoc"
	"example/gomicro/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	UserClient  user.UserService
)

func Login(c *gin.Context){
	fmt.Println("进来了")
	username := c.DefaultPostForm("username","")
	password := c.DefaultPostForm("password","")
	if password == "" || username == "" {
		utils.ResponseError(c,utils.ErrCodeParameter,nil)
		return
	}

	res,err := UserClient.Login(c,&user.Request{
		Username: username,
		Password: password,
	})

	if err != nil {
		utils.BgLogger.Errorf("Login:%v err:%v",username,err)
		utils.ResponseError(c,utils.ErrCodeRpc,nil)
		return
	}

	if !res.Ok {
		utils.ResponseError(c,utils.ErrCodeLogin,nil)
		return
	}

	userInfo := make(map[string]interface{})
	userInfo["token"] = res.Token
	userInfo["userId"] = res.UserId
	utils.ResponseSuccess(c,userInfo)
}

