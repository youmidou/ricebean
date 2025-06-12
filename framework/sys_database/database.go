package sys_database

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"phoenix-tudou/framework/sys_mysql"
	"phoenix-tudou/framework/sys_redis"
)

/*
import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"phoenix-tudou/framework/log4"
	"phoenix-tudou/framework/sys_mysql"
	"phoenix-tudou/framework/sys_redis"
	"phoenix-tudou/examples/CentralServer/sys_config"
)

// 初始化数据库
func Init(dbInfo *sys_config.DataBaseInfo) {

	//初始化mysql数据库

	if dbInfo == nil {
		log4.Panic("database.Init() -> dbinfo is nil")
	}
	log4.Info("Initialize mysql database ...")
	sys_mysql.Init(dbInfo.Mysql)

	//初始化redis数据库
	log4.Info("Initialize Redis database ...")
	sys_redis.Init(dbInfo.Redis)
	//log4.Info("Redis connection success ...")

}
*/
// Provide a transactional scope around a series of operations.
func GetDBSession() *gorm.DB {
	return sys_mysql.GetDBSession()
}
func GetRedis(key string) *redis.Client {
	return sys_redis.GetRedis(key)
}
