package databases

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

//DB DBのインターフェース
type DatabaseInterface interface {
	Open() *gorm.DB
}

type RedisInterface interface {
	Open() *redis.Client
}
