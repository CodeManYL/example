package mod

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

//用户基本信息
type GameUsers struct {
	GameID   int64  `json:"game_id" xorm:"int(11) pk notnull 'GameID'"`
	PrettyID int64  `json:"pretty_id" xorm:"int(11) pk notnull 'PrettyID'"`
	NickName string `json:"nick_name" xorm:"varchar(32) pk notnull 'NickName'"`
	Password string `json:"password" xorm:"varchar(32) pk notnull 'Password'"`
	Mobile   string `json:"mobile" xorm:"varchar(32) pk notnull 'Mobile'"`
	VipLevel int    `json:"vip_level" xorm:"int(11) pk notnull 'VipLevel'"`
}

//银行信息
type WebNewBankInfo struct {
	GameID          int64  `xorm:"int(11) pk notnull 'GameID'"`
	AlipayName      string `xorm:"varchar(32) pk notnull 'AlipayName'"`
	AlipayAccount   string `xorm:"varchar(32) pk notnull 'AlipayAccount'"`
	BankCardName    string `xorm:"varchar(32) pk notnull 'BankCardName'"`
	BankCardAccount string `xorm:"varchar(32) pk notnull 'BankCardAccount'"`
	BankName        string `xorm:"varchar(32) pk notnull 'BankName'"`
	BankBranchName  string `xorm:"varchar(32) pk notnull 'BankBranchName'"`
}

//导航
type WebNewNavigate struct {
	Index string `xorm:"varchar(32) pk notnull 'Index'"`
	Title string `xorm:"varchar(32) pk notnull 'Title'"`
	Icon string `xorm:"varchar(32)  notnull 'Icon'"`
}

var engine *xorm.Engine

func InitModEngine(dbName, dbAddress string) (err error) {
	engine, err = xorm.NewEngine(dbName, dbAddress)
	if err != nil {
		return
	}

	return nil
}
