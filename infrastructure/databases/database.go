package databases

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

//DB DBのインターフェース
type DatabaseInterface interface {
	Open() *gorm.DB
}

type RedisInterface interface {
	Open() *redis.Client
}
