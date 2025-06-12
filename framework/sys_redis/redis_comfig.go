package sys_redis

/*
redis 数据库结构设计
注意：为了解决单个db压力大 我们将更具数据类型分库存储 0-15 db 中
2,将动态生成配置 便于多个redis集群部署
added by yh @ 2023-04-26
*/
func GetRedisInfoList() map[string]*RedisInfo {
	//Redis map[string] *RedisInfo

	list := make(map[string]*RedisInfo)
	list[SystemData] = &RedisInfo{"127.0.0.1:6379", 13, ""}
	list[User] = &RedisInfo{"127.0.0.1:6379", 14, ""}
	list[User] = &RedisInfo{"127.0.0.1:6379", 12, ""}
	list[Pub] = &RedisInfo{"127.0.0.1:6379", 13, ""}

	list[Predict] = &RedisInfo{"127.0.0.1:6379", 15, ""}

	list[Property] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[DailyBonus] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[UserMission] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[UserProduct] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[ProductCoupon] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[StorePackageCoupon] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[MoneyBankCoupon] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[UserInbox] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[UserGuide] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[SevenDaysPrize] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[Activity] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[MoneyBank] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[Sapphire] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[ActivityCenter] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[Recoup] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[MarketingData] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	list[RateUs] = &RedisInfo{"127.0.0.1:6379", 15, ""}
	/*
		list[Quest] =&RedisInfo{"127.0.0.1:6379",4,""}
		list[Ranking] =&RedisInfo{"127.0.0.1:6379",4,""}
		list[CoinRush] =&RedisInfo{"127.0.0.1:6379",4,""}
		list[NormalQuest] =&RedisInfo{"127.0.0.1:6379",4,""}

		list[LoginBonus] =&RedisInfo{"127.0.0.1:6379",7,""}
		list[Freebie] =&RedisInfo{"127.0.0.1:6379",7,""}
	*/

	return list
}
