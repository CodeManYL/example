package concrol

import (
	"context"
	"example/gomicro/user_rpc/logic"
	user_pro "example/gomicro/user_rpc/protoc"
	"example/gomicro/utils"
)

type User struct{}

func (u *User) Login(ctx context.Context, req *user_pro.Request, rsp *user_pro.Response) error {
	//参数校验
	if req.Password == "" || req.Username == "" {
		err := utils.RpcError(utils.ErrCodeParameter)
		return err
	}
	//登陆逻辑
	userid, token, ok, err := logic.Login(req.Username, req.Password)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	rsp.Ok = true
	rsp.Token = token
	rsp.UserId = userid
	return nil
}

//func (u *User) UpBankInfo(ctx context.Context, req *user_pro.UpBankInfoRequest, rsp *user_pro.UpBankInfoResponse) error {
//	bankInfo := &mod.WebNewBankInfo{
//		GameID:          req.BankInfo.GameId,
//		AlipayName:      req.BankInfo.AlipayName,
//		AlipayAccount:   req.BankInfo.AlipayAccount,
//		BankCardName:    req.BankInfo.BankCardName,
//		BankCardAccount: req.BankInfo.BankCardAccount,
//		BankName:        req.BankInfo.BankName,
//		BankBranchName:  req.BankInfo.BankBranchName,
//	}
//
//	err := logic.UpBankInfoByGameID(bankInfo)
//	if err != nil {
//		utils.BgLogger.Errorf("up bank info failed:%v err :%v", req.BankInfo.GameId, err)
//		return err
//	}
//	return nil
//}
//
//func (u *User) GetBankInfo(ctx context.Context, req *user_pro.GetBankInfoRequest, rsp *user_pro.GetBankInfoResponse) error {
//	bankinfo, err := logic.GetBankInfoByGameID(req.GameId)
//	if err != nil {
//		utils.BgLogger.Errorf("GetBankInfo:%v err :%v", req.GameId, err)
//		return err
//	}
//
//	user := &user_pro.BankInfo{
//		GameId:          bankinfo.GameID,
//		AlipayName:      bankinfo.AlipayName,
//		AlipayAccount:   bankinfo.AlipayAccount,
//		BankCardName:    bankinfo.BankCardName,
//		BankCardAccount: bankinfo.BankCardAccount,
//		BankName:        bankinfo.BankName,
//		BankBranchName:  bankinfo.BankBranchName,
//	}
//	rsp.BankInfo = user
//
//	return nil
//}
//
//func (u *User) GetNavigate(ctx context.Context, req *user_pro.GetNavigateRequest, rsp *user_pro.GetNavigateResponse) error {
//	navigate, err := logic.GetNavigate()
//	if err != nil {
//		utils.BgLogger.Errorf("col.GetNavigate: err :%v", err)
//		return err
//	}
//	for _,val := range navigate {
//		meta := &user_pro.Navigate{
//			Icon:val.Icon,
//			Index:val.Index,
//			Title:val.Title,
//		}
//		rsp.Navigate = append(rsp.Navigate,meta)
//	}
//
//	return nil
//}
