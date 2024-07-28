package model

import (
	"github.com/Solituderr/autoenv/model/etcdm"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/kitex-contrib/config-etcd/etcd"
)

var EtcdRegistry registry.Registry
var EtcdClient etcd.Client

func Init() {
	EtcdClient = etcdm.NewEtcdClient()
	EtcdRegistry = etcdm.NewEtcdRegistry()
}
