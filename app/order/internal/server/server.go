package server

import (
	"shopping/app/order/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: conf.Etcd.Endpoints,
	})
	if err != nil {
		panic(err)
	}

	// new dis with etcd client
	r := etcd.New(client)
	return r
}
