package main

import (
	autoenv "github.com/Solituderr/autoenv/kitex_gen/chatgpt/model/autoenv/etcdenvservice"
	"log"
)

func main() {
	svr := autoenv.NewServer(new(EtcdEnvServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
