package data

import (
	"ad-marketing/app/adgroup/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAdgroupRepo, NewDb, NewRedisCmd)

// Data .
type Data struct {
	db       *gorm.DB
	redisCli redis.Cmdable
}

// NewData .
func NewData(logger log.Logger, db *gorm.DB, redisCli redis.Cmdable) (*Data, func(), error) {
	d := &Data{
		redisCli: redisCli,
		db:       db,
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return d, cleanup, nil
}

func NewDb(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "adgroup-service/data/client"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connect error: %v", err)
	}
	return db
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "adgroup-service/data/client"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}
