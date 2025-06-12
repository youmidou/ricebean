package sys_mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"phoenix-tudou/framework/log4"
	"phoenix-tudou/framework/log4/log_cmd"
	"time"
	//"github.com/jinzhu/gorm"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
)

//import "github.com/go-gorm/gorm"

var db *gorm.DB
var dsn string

/*
	Init 数据库 连接初始化

added by yh @ 2023/5/9 10:37
注意:这里只负责收 和删除
*/
func Init(info *MysqlInfo) {
	//dsn := "root:root@tcp(127.0.0.1:3306)/db_lqyy?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		//"root", "root", "127.0.0.1:3306", "db_lqyy")
		info.Root, info.Password, info.Address, info.Database)
	// Gorm 配置
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error), // 设置日志级别
	}

	_db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log4.Panic(fmt.Sprintf("mysql连接失败... %s err=>%s", dsn, err.Error()))
	}

	db = _db

	//-------避免出现 Error 1040: Too many connections-----------------------------------------------
	// 设置最大连接数
	sqlDB, err := db.DB()
	if err != nil {
		log4.Panic("Failed to get DB instance")
	}

	// 设置最大打开连接数和最大空闲连接数
	sqlDB.SetMaxOpenConns(50000)
	sqlDB.SetMaxIdleConns(10000)
	// 5秒内连接没有活跃的话则自动关闭连接
	sqlDB.SetConnMaxLifetime(time.Second * 5)

	// 获取连接池状态
	stats := sqlDB.Stats()
	// 打印连接池状态信息
	log4.LogData(log_cmd.SkyNet, "mysql",
		zap.Int("OpenConnections", stats.OpenConnections),
		zap.Int("Idle", stats.Idle),
		zap.Int("InUse", stats.InUse),
		zap.Int("MaxOpenConnections", stats.MaxOpenConnections),
		zap.Int64("MaxIdleClosed", stats.MaxIdleClosed),
		zap.Int64("WaitCount", stats.WaitCount),
	)

	db.AutoMigrate(
	//&dbtable.DBUser{},
	/*
		&dbtable.DBRole{},
		&dbtable.DBRoleOther{},
		&dbtable.DBProperty{},
		&dbtable.DBKnapsack{},
		&dbtable.DBInbox{},
		&dbtable.DBShop{},
		&dbtable.DBTask{},
		&dbtable.DBActivity{},
		&dbtable.DBGuide{},

		&dbtable.DBMatch3Theme{},
		&dbtable.DBHeroCard{},

		&dbtable.DBSystemInbox{},
		&dbtable.DBSystemShop{},
		&dbtable.DBCardBattle{},
	*/
	)

	// 执行原生SQL，设置自增的起始值为100
	//db.Exec("ALTER TABLE db_users AUTO_INCREMENT = 4575")

}

func GetDBSession() *gorm.DB {
	// 开始一个 Session
	//tx := db.Begin()
	//return tx
	return db.Session(&gorm.Session{})
}

type MysqlInfo struct {
	Address string
	//root
	Root     string
	Password string
	//数据库
	Database string
}
