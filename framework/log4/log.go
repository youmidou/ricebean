/*
	日志记录器

added by yh @ 2023/4/28 11:41
注意:
*/
package log4

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"ricebean/framework/log4/log_cmd"
	"ricebean/framework/sys_json"
)

var log_config = []log_cmd.LogCmd{
	log_cmd.Net, log_cmd.Error, log_cmd.Fail, log_cmd.Timeout, log_cmd.CfgError,
	log_cmd.ErrorCfg, log_cmd.ErrorItem,
	log_cmd.Login, log_cmd.Logout,
	log_cmd.Facebook,
	//天网系统
	log_cmd.SkyNet, log_cmd.SkyNetSys,
	log_cmd.SkyNetAdmin,

	log_cmd.Money, log_cmd.MoneyGoldBindUse,
	log_cmd.Ranking,
	log_cmd.Room, log_cmd.RoomGameOver, log_cmd.HeroAttack,
	log_cmd.SkyRoom,
	log_cmd.Knapsack, //背包日志
	log_cmd.CollectFireDragonTreasure,
	log_cmd.BuyItem,
}
var nodeRoot = "./"
var directory = "./log"          //../root/log
var game_log_directory = "./log" //../root/game_log

// 初始化日志配置
func Init(_nodeRoot string) {
	//# 每行开头的日期格式
	nodeRoot = _nodeRoot
	//port := "1249"
	//# 根据配置初始化日志
	for _, log_name := range log_config {
		if log_name == log_cmd.Net {
			//# 原始日志
			logger := setNetLogger(fmt.Sprintf("%s/%s/%s.log", directory, nodeRoot, log_name))
			loggerDict[log_name] = logger
		} else {
			// # 格式化日志
			logger := setLogger(fmt.Sprintf("%s/%s/%s.log", directory, nodeRoot, log_name))
			loggerDict[log_name] = logger
		}
	}

}

// LogData 输出Splunk日志
func LogData(logName log_cmd.LogCmd, msg string, fields ...zap.Field) {
	logger := GetLogger(logName)
	logger.Info(msg, fields...)
}

//--------定制日志------------------------------------------------------------

func InfoUser(logName log_cmd.LogCmd, user any) {
	str, err := sys_json.MarshalToString(user)
	if err != nil {
		Error("Log4.InfoUser->json %s", err.Error())
	}
	LogData(logName, str)
}

// Debug 调试用不参与储存
func Debug(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	logger := GetLogger(log_cmd.Net)
	logger.Debug(msg)
	//logger.Debug(msg, fields...)
}

/*
Info %s 字符串 %d（整数）、%f（浮点数）和%t（布尔值）。
fmt.Sprintf("%s %d", a, b)
*/
func Info(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	logger := GetLogger(log_cmd.Net)
	logger.Info(msg)
}
func Warn(msg string, fields ...zap.Field) {
	log.Print(msg)
	logger := GetLogger(log_cmd.Net)
	logger.Warn(msg, fields...)
}
func Error(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	//log.Print(msg)
	logger := GetLogger(log_cmd.Net)
	logger.Error(msg)

	logger = GetLogger(log_cmd.Error)
	logger.Error(msg)
}
func Exception(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	//log.Print(msg)
	logger := GetLogger(log_cmd.Net)
	logger.Error(msg)

	logger = GetLogger(log_cmd.Error)
	logger.Error(msg)
}

/*
DPanic 是一个“双重危险”的级别，它表示程序遇到了一个错误，但程序仍然可以继续执行。
在这种情况下，DPanic会记录错误并打印堆栈跟踪，但不会导致程序崩溃。
added by yh @ 2023/4/28 15:55
*/
func DPanic(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	logger := GetLogger(log_cmd.Net)
	logger.DPanic(msg)

	logger = GetLogger(log_cmd.Error)
	logger.DPanic(msg)
}

/*
Panic 是一个更高级别的错误级别，它表示程序遇到了一个无法恢复的错误，需要立即停止执行。
在这种情况下，Panic会记录错误并打印堆栈跟踪，然后导致程序崩溃
*/
func Panic(msg string, fields ...zap.Field) {
	logger := GetLogger(log_cmd.Net)
	logger.Panic(msg, fields...)
}

func Fatal(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	logger := GetLogger(log_cmd.Net)
	logger.Fatal(msg)
	logger = GetLogger(log_cmd.Error)
	logger.Fatal(msg)
}
