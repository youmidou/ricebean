package log_cmd

/*
*
日志收集器
注意：
added by yh @ 2023-04-26
*/
type LogCmd string

const (
	//infor 日志不参与存储
	//通信日志
	Net LogCmd = "Net"
	//错误日志
	Error       LogCmd = "Error"
	CfgError    LogCmd = "CfgError"
	SkyNetAdmin LogCmd = "SkyNetAdmin" //天网管理员操作日志
	//失败
	Fail      LogCmd = "Fail"
	ErrorCfg  LogCmd = "ErrorCfg"  //配置文件错误
	ErrorItem LogCmd = "ErrorItem" //使用非法物品
	//超时
	Timeout LogCmd = "timeout"

	//-----定制日志----------------------------
	//登录日志
	Login LogCmd = "Login"
	//登录日志
	Logout LogCmd = "Logout"
	//新手
	//new_guide="new_guide"
	//
	Facebook LogCmd = "Facebook"
	//today_begin
	TodayBegin LogCmd = "today_begin"
	//每日奖励日志
	DailyBonus LogCmd = "DailyBonus"
	//天网
	SkyNet LogCmd = "sky_net"
	//天网系统 ->定点 监控服务器性能 资源使用状况 1:在线人数; 2:cpu占用百分百;3:内存占用百分百;4:数据库链接数;5:最大链接数;
	SkyNetSys LogCmd = "sky_net_sys"

	//----------------------
	Role      LogCmd = "Role"
	RoleEquip LogCmd = "RoleEquip"

	Money            LogCmd = "Money" //玩家货币使用
	MoneyGoldBindUse LogCmd = "MoneyGoldBindUse"
	Knapsack         LogCmd = "Knapsack" //背包
	//排行榜
	Ranking LogCmd = "ranking"
	//记录1:创建房间;2:结算;3:解散房间;
	Room         LogCmd = "room"
	RoomGameOver LogCmd = "RoomGameOver"
	HeroAttack   LogCmd = "HeroAttack"
	//天网监控房间信息
	SkyRoom         LogCmd = "sky_room"
	CreateMatchRoom LogCmd = "CreateMatchRoom"
	//收集火龙宝藏值
	CollectFireDragonTreasure LogCmd = "CollectFireDragonTreasure"
	//----商店----------------------
	BuyItem LogCmd = "BuyItem" //购买物品
)
