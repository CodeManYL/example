package gateway_conf

import (
	"example/gomicro/utils"
	"flag"
	"github.com/micro/go-micro/v2/config"
)

type GatewayConf struct {
	ServerName string `json:"server_name"`
	ServerAddr string `json:"server_addr"`
	Etcd struct{
		Address string `json:"address"`
	} `json:"etcd"`
	Db struct {
		Name string `json:"name"`
		Address string `json:"address"`
	} `json:"db"`
	UserRpcServer struct {
		ClientName string `json:"client_name"`
		ServerName string `json:"server_name"`
	} `json:"user_rpc_server"`
}


func InitConfig() (conf *GatewayConf, err error) {
	//初始化配置文件
	configFile := flag.String("e", "../../conf/gateway_conf/gateway_conf.json", "please use xxx -f config_rpc.json")
	flag.Parse()
	conf = &GatewayConf{}
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