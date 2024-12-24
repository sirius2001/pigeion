package data

import (
	"pigeon2/internal/biz"
	"pigeon2/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"go.temporal.io/sdk/client"
)

type temporalRepo struct {
	client client.Client
}

func NewTemporalRepo(c *conf.Data, logger log.Logger) biz.TemproalRepo {
	client, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "default",
		Logger:    nil,
	})

	if err != nil {
		panic(err)
	}

	return &temporalRepo{
		client: client,
	}
}
