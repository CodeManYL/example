package data

import (
	user "example/engineering-layout/app/service/users/api"
	"example/engineering-layout/app/service/users/internal/biz"
	"github.com/pkg/errors"
	"github.com/xormplus/xorm"
	_ "github.com/go-sql-driver/mysql"

)

var _ biz.UserMod = (*userMod)(nil)

const  (
	_getUserByMobileAndPassword = "select * from web_base_user where user_name = ? and  password = ?"
)

type userMod struct {
	engine *xorm.Engine
}

func NewUserData(dbName, dbAddress string) (biz.UserMod,error) {
	engine, err := xorm.NewEngine(dbName, dbAddress)
	if err != nil {
		return nil,err
	}
	return &userMod{engine,},nil
}

func (u *userMod) GetUserByUsernameAndPassword(username,password string) (*biz.UserInfo,error){
	userInfo := &biz.UserInfo{}
	ok,err := u.engine.SQL(_getUserByMobileAndPassword,username,password).Get(userInfo)
	if err != nil {
		return nil, errors.Wrap(err,_getUserByMobileAndPassword)
	}

	if !ok {
		return nil,user.ErrQueryNotExist
	}

	return userInfo,nil
}