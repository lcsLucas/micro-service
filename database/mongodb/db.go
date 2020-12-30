package mongodb

import (
	"context"
	"fmt"

	"github.com/lcslucas/micro-service/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBURL string

func Connect(ctx context.Context, c config.ConfigDB) (*mongo.Client, error) {
	credential := options.Credential{
		Username: "root",
		Password: "root",
	}

	fmt.Println(fmt.Sprintf("%s://%s:%s@%s:%s/pedidos?authMechanism=SCRAM-SHA-1", c.Driver, c.User, c.Password, c.Host, c.Port))
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s@%s:%s", c.Driver, c.User, c.Password, c.Host, c.Port)).SetAuth(credential))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
