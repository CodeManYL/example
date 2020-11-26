package utils

import (
	"github.com/micro/go-micro/v2/errors"
	"net/http"
)

const (
	ErrCodeSuccess           = 200
	ErrCodeParameter         = 1001
	ErrCodeQueryObjNotExist  = 1002
	ErrCodeDb = 1013
	ErrCodeGetTokenFailed = 1014
	ErrCodeRpc = 1015
	ErrCodeLogin = 1016
	ErrCodeValidToken = 1017
	ErrCodeParameterChange        = 1018
	ErrCodeGoodsNums              = 1019
	//ErrCodeUserNotExist      = 1004
	//ErrCodeUserPasswordWrong = 1005
	//ErrCodeCaptionHit        = 1006
	//ErrCodeContentHit        = 1007
	//ErrCodeNotLogin          = 1008
	//ErrInvalidToken       = 1009
	//ErrCodeGetTokenFailed       = 1011
	//ErrRPC = 1013
	//ErrCodeParamNotExist = 1014
	//ErrCodeResultNotExist = 1016

)

func GetMessage(code int) (message string) {
	switch code {
	case ErrCodeSuccess:
		message = "success"
	case ErrCodeParameter:
		message = "参数缺少"
	case ErrCodeParameterChange:
		message = "参数转化错误"
	case ErrCodeQueryObjNotExist:
		message = "查询对象不存在"
	case ErrCodeDb:
		message = "数据库相关错误"
	case ErrCodeGetTokenFailed:
		message = "获取token失败"
	case ErrCodeRpc:
		message = "rpc相关错误"
	case ErrCodeLogin:
		message = "账号或密码错误"
	case ErrCodeValidToken:
		message = "token失效"
	case ErrCodeGoodsNums:
		message = "商品数量不足"
	default:
		message = "未知错误"
	}
	return
}

func RpcError(code int32)(err  error){
	err = &errors.Error{
		Code: code,//code = 1001 表示查询不存在
		Status: http.StatusText(500),
	}
	return
}