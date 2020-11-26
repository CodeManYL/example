package mod

import (
	"example/gomicro/utils"
)

func Login (mobile,password string)(userid int64,username string,ok bool,err error){
	user := &GameUsers{}
	//password = utils.Md5(password)
	ok,err = engine.Where("Mobile = ? and Password = ?",mobile,password).Get(user)
	if !ok || err != nil{
		return
	}
	userid = user.GameID
	username = user.NickName
	return
}

func GetBankInfoByGameID(gameId int64) (bankInfo *WebNewBankInfo,ok bool,err error){
	bankInfo = &WebNewBankInfo{}
	ok,err = engine.Where("GameID = ?",gameId).Get(bankInfo)
	if !ok {
		return
	}
	if err != nil {
		utils.BgLogger.Errorf("get bank info failed: %v err : %v",gameId,err)
		return
	}
	return
}

func UpdateBankInfoByGameID(bankInfo *WebNewBankInfo) (err error){
	_,err = engine.Where("GameID = ?",bankInfo.GameID).Update(bankInfo)
	if err != nil {
		utils.BgLogger.Errorf("update bank info failed: %v err : %v",bankInfo.GameID)
		return
	}
	return
}

func CreateBankInfoByGameID(bankInfo *WebNewBankInfo) (err error){
	_,err = engine.Insert(bankInfo)
	if err != nil {
		utils.BgLogger.Errorf("create bank info failed: %v err : %v",bankInfo.GameID,err)
		return
	}
	return
}

func GetNavigate()(navigate []*WebNewNavigate,err error){

	err = engine.Find(&navigate)
	if err != nil {
		utils.BgLogger.Errorf("get navigate  failed:  err : %v",err)
		return
	}
	return
}
