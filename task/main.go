package main

import (
	"task/conf"
	"task/core"
	"task/services"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	conf.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	microService := micro.NewService(
		micro.Name("rpcTaskService"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg),
	)
	microService.Init()
	_ = services.RegisterTaskServiceHandler(microService.Server(), new(core.TaskService))
	_ = microService.Run()
}
