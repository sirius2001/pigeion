package data

import (
	"errors"
	"pigeon2/internal/conf"
	"pigeon2/pkg/persistence"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
	NewTemporalRepo,
	NewWorkFlowRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	var dialer gorm.Dialector
	switch c.Database.Driver {
	case "mysql":
		mysql.Open(c.Database.Source)
	default:
		return &Data{}, cleanup, errors.New("unknown database driver")
	}

	db, err := gorm.Open(dialer, &gorm.Config{})
	if err != nil {
		return &Data{}, cleanup, err
	}

	persistence.SetDefault(db)

	return &Data{}, cleanup, nil
}
