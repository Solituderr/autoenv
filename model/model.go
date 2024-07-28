package model

import (
	"github.com/Solituderr/autoenv/model/etcdm"
	etcdServer "github.com/kitex-contrib/config-etcd/server"
)

var EtcdServerSuite *etcdServer.EtcdServerSuite

func Init() {
	EtcdServerSuite = etcdm.NewEtcdClient()
}
