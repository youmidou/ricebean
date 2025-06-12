package sys_redis

import (
	"github.com/go-redis/redis"
	"phoenix-tudou/framework/log4"
	"phoenix-tudou/framework/sys_json"
)

/*
Redis 客户端调
added by yh @ 2023-04-25
*/
type RedisClient struct {
	client *redis.Client
	Info   *RedisInfo
	IsBool bool
}

func NewRedisClient() *RedisClient {
	t := &RedisClient{}
	return t
}
func (t *RedisClient) Connect(info *RedisInfo) bool {
	t.Info = info
	// 创建redis客户端
	t.client = redis.NewClient(&redis.Options{
		Addr:     info.Addr,     //"localhost:6379",
		DB:       info.DB,       // 0 Redis数据库索引
		Password: info.Password, //"" Redis密码
	})

	return t.Ping()
}

func (t *RedisClient) Ping() bool {
	str, _ := sys_json.MarshalToString(t.Info)
	_, err := t.client.Ping().Result()

	if err != nil {
		log4.Panic("Failed to connect to Redis " + str)
		return false
	} else {
		//log4.Info("Connected to Redis Success " + str)
		return true
	}
}

// 写入
func (t *RedisClient) Set(key string, value string) {
	// 写入数据
	err := t.client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

// 读取
func (t *RedisClient) Get(key string) string {
	// 读取数据
	val, err := t.client.Get(key).Result()
	if err != nil {
		log4.Panic("RedisClient->Get " + err.Error())
		panic(err)
	}
	return val
}
func (t *RedisClient) HGet(key string, field string) string {
	// 读取数据
	val, err := t.client.HGet(key, field).Result()
	if err != nil {
		log4.Panic("RedisClient->Get " + err.Error())
		panic(err)
	}
	return val
}

func (t *RedisClient) GetClient() *redis.Client {
	return t.client
}
