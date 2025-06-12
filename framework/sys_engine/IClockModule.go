package sys_engine

import "time"

/*
	时钟模块接口

added by yh @ 2024/05/11 11:24
注意:
*/
type IClockModule interface {
	IModule //继承模块接口
	//获取当前帧
	GetCurrentFrame() uint64
	//获取帧时间
	GetFrameTime() uint64
	GetFPS() uint64

	SetFPS(fps uint64)
	//获取当前时间
	Time() time.Time
	//日期ID
	DayID() uint32
	//本地时间
	LocalTime() time.Time
	//UTC时间
	AscTime() string
	//时间戳-下一分钟间隔
	NextMinuteInterval(sceond int32) int64
	//时间戳-下一小时间隔
	NextHourInterval(minute int32, sceond int32) int64
	//时间戳-下一天间隔
	NextDayInterval(hour int32, minute int32, sceond int32) int64
	//时间戳-下一周间隔
	NextWeekInterval(week int32, hour int32, minute int32, sceond int32) int64
	//时间戳-下一月间隔
	NextMouthInterval(day int32, hour int32, minute int32, sceond int32) int64
	//时间戳-比较时间与本地时间间隔
	LocalTimeInterval(compareTime time.Time) int64
}
