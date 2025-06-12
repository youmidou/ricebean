package sys_redis

/*
redis 数据
注意:根据配置建立 0-15 redis.Client 数据库连接
1.Redis 为了减少数据压力 我们将user 放在db0里面

added by yh @ 2023-04-24
*/

import (
	"fmt"
	"github.com/go-redis/redis"
	"phoenix-tudou/framework/log4"
)

// 客户端连接列表
var redisList = make(map[string]*redis.Client)

//var client *redis.Client

// 0-15
func Init(redisInfoList map[string]*RedisInfo) {

	for key, info := range redisInfoList {
		//log4.Infof("key:%s, value:%d\n", key, value)
		// 创建redis客户端
		client := redis.NewClient(&redis.Options{
			Addr:     info.Addr,     //"localhost:6379",
			DB:       info.DB,       // 0 Redis数据库索引
			Password: info.Password, //"" Redis密码
		})
		redisList[key] = client
	}

}
func GetRedis(key string) *redis.Client {
	if redisList[key] == nil {
		log4.Panic(fmt.Sprintf("GetRedis key =%s is nil 需要配置", key))
	}
	return redisList[key]
}

/*
// 写入
func HSet() {
	// 将数据存储到Redis中
	err := client.HSet("folder_name", "field_name", "data").Err()
	if err != nil {
		panic(err)
	}
}

// 写入
func Set() {
	// 写入数据
	user := &dbtable.DBRole{}
	user.Id = 100
	user.FacebookId = "fbid"
	user.NewUser = 1
	user.PkgId = 0
	str, _ := sys_json.MarshalToString(user)
	err := client.Set("lqyy:user_keywawa3", str, 0).Err()
	if err != nil {
		panic(err)
	}
}

// 读取
func Get() {
	// 读取数据
	val, err := client.Get("keywawa").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("keywawa", val)

}
*/
