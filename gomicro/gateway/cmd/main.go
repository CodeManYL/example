package main

import (
	"example/gomicro/conf/gateway_conf"
	"example/gomicro/gateway/logic"
	"example/gomicro/user_rpc/mod"
	user_pro "example/gomicro/user_rpc/protoc"
	"example/gomicro/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

func InitConf() *gateway_conf.GatewayConf{
	//初始化配置文件
	conf,err := gateway_conf.InitConfig()
	if err != nil {
		panic(fmt.Sprintf("init config failed err :%v",err))
	}
	//初始化数据库
	if err := mod.InitModEngine(conf.Db.Name,conf.Db.Address) ; err != nil {
		panic(fmt.Sprintln("init db failed err : %v",err))
	}
	//初始化日志
	if err := utils.InitLogger(fmt.Sprintf("%v.txt",conf.ServerName));err != nil {
		panic(fmt.Sprintf("init log failed err :%v",err))
	}
	return conf
}



func main(){
	conf := InitConf()
	fmt.Println("conf:",conf)
	//etc注册中心

	etcd := etcdv3.NewRegistry(
		registry.Addrs(conf.Etcd.Address),
	)

	gateway := web.NewService(
		web.Name(conf.ServerName),
		web.Registry(etcd),
		web.Address(conf.ServerAddr),
		web.Flags(&cli.StringFlag{
			Name:  "e",
			Value: "./config/config_rpc.json",
			Usage: "please use xxx -f config_rpc.json",
		}),
	)

	if err := gateway.Init();err != nil {
		panic(err)
	}


	//client
	logic.UserClient = user_pro.NewUserService(conf.UserRpcServer.ServerName,client.DefaultClient)
	//gin
	router := gin.Default()
	userRouterGroup := router.Group("/user")
	{
		userRouterGroup.POST("/login", logic.Login)
	}

	gateway.Handle("/", router)

	if err := gateway.Run(); err != nil {
		panic(err)
	}


}