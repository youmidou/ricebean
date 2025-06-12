/*
added by yh @ 2023/4/28 11:57
注意:防止丢失

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
*/
package log4

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"ricebean/framework/log4/log_cmd"
	"ricebean/framework/sys_getenv"

	"time"
)

var loggerDict = make(map[log_cmd.LogCmd]*zap.Logger)

func GetLogger(filename log_cmd.LogCmd) *zap.Logger {
	if filename == "" {
		panic("GetLogger name is nil")
		return nil
	}
	var rv *zap.Logger
	rv = loggerDict[filename]
	if rv == nil {
		//rv = setLogger(filename)
		//loggerDict[filename] = rv
		panic("GetLogger no init filename =" + filename)
		return nil
	}
	return rv
}

// net日志 原始日志
func setNetLogger(filename string) *zap.Logger {

	if !sys_getenv.IsLinux() {

		loggerD, _ := zap.NewDevelopment()
		options := zap.AddStacktrace(zapcore.InfoLevel)
		loggerD.WithOptions(options)
		return loggerD
	}

	// 创建一个FileWriter
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		//Filename:   "./log/net.log", // 日志文件路径
		Filename:   filename, // 日志文件路径
		MaxSize:    30,       // 每个日志文件最大大小（MB）
		MaxBackups: 31,       // 最多保留的日志文件数量
		MaxAge:     31,       // 日志文件最大保存时间（天）
	})
	// 配置logger
	encoderConfig := zap.NewProductionEncoderConfig()
	/*
		EpochTimeEncoder
		EpochMillisTimeEncoder
		EpochNanosTimeEncoder
		ISO8601TimeEncoder
		RFC3339TimeEncoder
		RFC3339NanoTimeEncoder
	*/
	encoderConfig.EncodeTime = RFC3339TimeEncoder //zapcore.ISO8601TimeEncoder // 时间格式
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		fileWriter,
		zap.InfoLevel, // 日志级别
	)
	//options := zap.AddCaller()
	//options := zap.Development()
	//options := zap.WithCaller(true)
	//options := zap.AddStacktrace(zapcore.ErrorLevel)
	logger := zap.New(core)
	return logger
}

// 格式化日志
func setLogger(filename string) *zap.Logger {
	// 创建一个FileWriter
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		//Filename:   "./log/net.log", // 日志文件路径
		Filename:   filename, // 日志文件路径
		MaxSize:    30,       // 每个日志文件最大大小（MB）
		MaxBackups: 31,       // 最多保留的日志文件数量
		MaxAge:     31,       // 日志文件最大保存时间（天）
	})

	// 配置logger
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = RFC3339TimeEncoder //zapcore.ISO8601TimeEncoder // 时间格式
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		fileWriter,
		zap.InfoLevel, // 日志级别
	)
	options := zap.Development()
	logger := zap.New(core, options)
	return logger
}
func RFC3339TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, time.DateTime, enc)
}
func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}
func NewProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "t",
		LevelKey:       "l",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "m",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
