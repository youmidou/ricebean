package sys_redis

type RedisInfo struct {
	//127.0.0.1:6379
	Addr string
	// Redis数据库索引 1
	DB int
	// Redis密码
	Password string

}