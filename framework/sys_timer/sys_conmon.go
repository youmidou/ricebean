package sys_timer

import "time"

// 将时间戳转换time
func TimestampToTime(timestamp int64) time.Time {
	// 将 Unix 时间戳转换为 time.Time 对象
	timeFromTimestamp := time.Unix(timestamp, 0)
	return timeFromTimestamp
}

// 持续时间是否大于间隔
func IsDurationGreaterThanInterval(startTime time.Time, interval time.Duration) bool {
	currentTime := time.Now()
	elapsedTime := currentTime.Sub(startTime)
	elapsedTime.Nanoseconds()
	return elapsedTime > interval
}

// 计算两个时间点之间的间隔
func CheckIsTimeout(startTime time.Time, interval time.Duration) bool {
	currentTime := time.Now()
	elapsedTime := currentTime.Sub(startTime)
	return elapsedTime > interval
}

// 计算两个时间点之间的间隔
func CheckIsTimeout1(startTime time.Time, endTime time.Time) bool {
	//currentTime = time.Now()
	elapsedTime := endTime.Sub(startTime)
	return elapsedTime > 0
}

// 计算两个时间点之间的间隔
func CheckIsTimeout2(startTime time.Time, endTime time.Time, interval time.Duration) bool {
	//currentTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	return elapsedTime > interval
}

// 判断 之前时间 加一天 是不是下一天
func IsNextDay(timestamp time.Time) bool {
	// 将时间戳转换为时间
	//t := time.Unix(timestamp, 0)

	// 获取当前时间
	now := time.Now()

	// 获取明天的时间
	nextDay := now.Add(24 * time.Hour)

	// 检查是否是下一天
	return timestamp.After(now) && timestamp.Before(nextDay)
}

// 检查是否是下一天凌晨00:00 零点
func IsNextDay2(timestamp time.Time) bool {
	return true
}

// 间隔天数 计算开始时间 与结束时间相差几天
func DaysBetween(start time.Time, end time.Time) int {
	// 将时间调整到午夜
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())

	duration := end.Sub(start)
	return int(duration.Hours() / 24)
}
func DaysBetween2(start time.Time, end time.Time) int {
	duration := end.Sub(start)
	return int(duration.Hours() / 24)
}

/*
// 延迟下一帧执行 fun-回调函数

	func NextFrameRun(fun func()) *TimerTask {
		t := NewTimerTask(10*time.Millisecond, fun)
		//t.AddFunc(fun)
		t.Start()
		return t
	}

	func DelayTimeRunLoop(time time.Duration, fun func()) *TimerTask {
		if time == 0 {
			fun()
			return nil
		}
		t := NewTimerTask(time, fun)
		//t.AddFunc(fun)
		t.Start()
		return t
	}

func DelayTimeAlarmClock(time time.Duration, loop int, fun func(), endBack func(int)) {

}

// 延迟时间执行 time-延迟时间单位 秒 ,fun-回调函数

	func DelayTimeRun(time time.Duration, fun func()) *TimerTask {
		if time == 0 {
			fun()
			return nil
		}
		t := NewTimerTask(time, fun)
		//t.AddFunc(fun)
		t.Start()
		return t
	}
*/
/*
// 取消延迟时间执行
func CancelDelayRun(t *TimerTask) {
	if t != nil {
		t.Stop()
	}
	t = nil
}
*/
