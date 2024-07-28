package main

import (
	"github.com/Solituderr/autoenv/conf"
	autoenv "github.com/Solituderr/autoenv/kitex_gen/chatgpt/model/autoenv/etcdenvservice"
	"github.com/Solituderr/autoenv/model"
	"github.com/Solituderr/autoenv/service"
	"github.com/Solituderr/autoenv/util"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
)

func main() {
	model.Init()
	util.Init()
	service.Init()
	svr := autoenv.NewServer(
		new(EtcdEnvServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.Config.Service.Name}),
		server.WithSuite(model.EtcdServerSuite),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
