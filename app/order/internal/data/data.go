package data

import (
	"context"
	"database/sql"
	"shopping/app/order/internal/conf"

	userv1 "shopping/api/user/v1"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrderRepo, NewUserServiceClient, NewDiscovery)

// Data .
type Data struct {
	db *sql.DB
	uc userv1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc userv1.UserClient) (*Data, func(), error) {
	db, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		db.Close()
	}
	d := &Data{
		db: db,
		uc: uc,
	}
	return d, cleanup, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: conf.Etcd.Endpoints,
	})
	if err != nil {
		panic(err)
	}
	return etcd.New(client)
}

func NewUserServiceClient(r registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///shopping.user.service"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}
