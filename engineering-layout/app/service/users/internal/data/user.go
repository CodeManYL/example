package data

import (
	user "example/engineering-layout/app/service/users/api"
	"example/engineering-layout/app/service/users/configs"
	"example/engineering-layout/app/service/users/internal/biz"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/xormplus/xorm"
)

var _ biz.UserMod = (*userMod)(nil)

var MockUserRepoSet = wire.NewSet(NewUserData,wire.Bind(new(biz.UserMod), new(*userMod)))

const  (
	_getUserByMobileAndPassword = "select * from web_base_user where user_name = ? and  password = ?"
)

type userMod struct {
	engine *xorm.Engine
}

func NewUserData2()*userMod{
	return &userMod{}
}

func NewUserData(cof *configs.UserRpcConf) (*userMod,error) {
	engine, err := xorm.NewEngine(cof.Db.Name, cof.Db.Address)
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