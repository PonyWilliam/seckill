package data

import (
	"seckill/internal/conf"
	"seckill/sqld"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	// TODO wrapped database client
	Db *sqld.MySQL_D
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db := sqld.SQL_init()
	err := db.Force_Connect(c.Database.Driver,c.Database.Source)
	if err != nil{
		log.NewHelper(logger).Error(err)
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		db.Close()
	}
	return &Data{Db:db}, cleanup, nil
}
