package user_rpc_conf

import (
	"example/gomicro/utils"
	"flag"
	"github.com/micro/go-micro/v2/config"
)

type UserRpcConf struct {
	ServerName string `json:"server_name"`
	ServerAddr string `json:"server_addr"`
	Etcd struct{
		Address string `json:"address"`
	} `json:"etcd"`
	Db struct {
		Name string `json:"name"`
		Address string `json:"address"`
	} `json:"db"`

}


func InitConfig() (conf *UserRpcConf, err error) {
	//初始化配置文件
	configFile := flag.String("e", "../../conf/user_rpc_conf/user_rpc_conf.json", "please use xxx -f config_rpc.json")
	flag.Parse()
	conf = &UserRpcConf{}
	if err := config.LoadFile(*configFile); err != nil {
		utils.BgLogger.Error(err)
		return nil, err
	}
	if err := config.Scan(conf); err != nil {
		utils.BgLogger.Error(err)
		return nil, err
	}
	return
}