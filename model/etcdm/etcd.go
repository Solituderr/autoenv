package etcdm

import (
	"github.com/Solituderr/autoenv/conf"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/kitex-contrib/config-etcd/etcd"
	etcdregistry "github.com/kitex-contrib/registry-etcd"
)

func NewEtcdRegistry() registry.Registry {
	r, err := etcdregistry.NewEtcdRegistry([]string{conf.Config.Service.Name})
	if err != nil {
		panic(err)
	}
	return r
}

func NewEtcdClient() etcd.Client {
	c, err := etcd.NewClient(etcd.Options{Node: []string{conf.Config.Etcd.Host}})
	if err != nil {
		panic(err)
	}
	return c
}
