package main

import (
	"example/gomicro/conf/user_rpc_conf"
	"example/gomicro/user_rpc/concrol"
	"example/gomicro/user_rpc/mod"
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	user_pro "example/gomicro/user_rpc/protoc"

)

func main() {
	//初始化配置文件
	conf,err := user_rpc_conf.InitConfig()
	if err != nil {
		panic(fmt.Sprintf("init config failed err :%v",err))
	}
	fmt.Println(conf)
	//初始化数据库
	if err := mod.InitModEngine(conf.Db.Name,conf.Db.Address) ; err != nil {
		panic(fmt.Sprintln("init db failed err : %v",err))
	}

	etcd := etcdv3.NewRegistry(
		registry.Addrs(conf.Etcd.Address),
	)

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name(conf.ServerName),
		micro.Transport(grpc.NewTransport()),
		micro.Registry(etcd),
		micro.Address(conf.ServerAddr),
		micro.Flags(&cli.StringFlag{
			Name:  "e",
			Value: "./config/config_rpc.json",
			Usage: "please use xxx -f config_rpc.json",
		}),
	)

	service.Init()

	if err := user_pro.RegisterUserHandler(service.Server(), new(concrol.User));err != nil {
		panic(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}