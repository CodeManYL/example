package logic

import (
	"example/gomicro/user_rpc/mod"
	"example/gomicro/utils"
)

func Login(mobile,password string)(userid int64,token string,ok bool,err error){
	userid,username,ok,err := mod.Login(mobile,password)
	if !ok {
		//err = utils.RpcError(utils.ErrCodeQueryObjNotExist)
		return
	}
	if err != nil {
		err = utils.RpcError(utils.ErrCodeDb)
		utils.BgLogger.Errorf("login:%v,db failed err: %v",userid,err)
		return
	}

	//获取TOKEN
	token,err = utils.GetToken(username,userid)
	if err != nil {
		err = utils.RpcError(utils.ErrCodeGetTokenFailed)
		utils.BgLogger.Errorf("get token failed:%v, err: %v",userid,err)
		return
	}

	return
}

////跟新或新增银行信息
//func UpBankInfoByGameID(bankInfo *mod.WebNewBankInfo)(err error){
//	_,ok,err := mod.GetBankInfoByGameID(bankInfo.GameID)
//	if err != nil {
//		err = utils.RpcError(utils.ErrCodeDb)
//		utils.BgLogger.Errorf("UpBankInfoByGameID:%v,db failed err: %v",bankInfo.GameID,err)
//		return
//	}
//
//	if ok {
//		//更新
//		err = mod.UpdateBankInfoByGameID(bankInfo)
//		if err != nil {
//			err = utils.RpcError(utils.ErrCodeDb)
//			utils.BgLogger.Errorf("UpBankInfoByGameID:%v,db failed err: %v",bankInfo.GameID,err)
//		}
//		return
//	}
//	//创建
//	err = mod.CreateBankInfoByGameID(bankInfo)
//	if err != nil {
//		err = utils.RpcError(utils.ErrCodeDb)
//		utils.BgLogger.Errorf("UpBankInfoByGameID:%v,db failed err: %v",bankInfo.GameID,err)
//	}
//
//	return
//}
//
////获得银行信息
//func GetBankInfoByGameID(gameid int64)(bankInfo *mod.WebNewBankInfo,err error){
//	bankInfo,_,err = mod.GetBankInfoByGameID(gameid)
//	if err != nil {
//		err = utils.RpcError(utils.ErrCodeDb)
//		utils.BgLogger.Errorf("UpBankInfoByGameID:%v,db failed err: %v",gameid,err)
//		return
//	}
//	return
//}
//
////获取导航
//func GetNavigate()(navigate []*mod.WebNewNavigate,err error){
//	navigate,err = mod.GetNavigate()
//	if err != nil {
//		err = utils.RpcError(utils.ErrCodeDb)
//		utils.BgLogger.Errorf("logic.GetNavigate:,db failed err: %v",err)
//		return
//	}
//
//	return
//}
