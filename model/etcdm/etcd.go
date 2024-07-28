package etcdm

import (
	"github.com/Solituderr/autoenv/conf"
	"github.com/kitex-contrib/config-etcd/etcd"
	etcdServer "github.com/kitex-contrib/config-etcd/server"
)

var EtcdClient etcd.Client

func NewEtcdClient() *etcdServer.EtcdServerSuite {
	var err error
	EtcdClient, err = etcd.NewClient(etcd.Options{Node: []string{conf.Config.Etcd.Host}})
	suite := etcdServer.NewSuite(conf.Config.Service.Name, EtcdClient)

	if err != nil {
		panic(err)
	}
	return suite
}
