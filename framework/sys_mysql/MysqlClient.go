package sys_mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"phoenix-tudou/framework/log4"
)

type MysqlClient struct {
	db *gorm.DB
}

func NewMysqlClient(info *MysqlInfo) *MysqlClient {
	t := &MysqlClient{}
	t._Init(info)
	return t
}
func (t *MysqlClient) _Init(info *MysqlInfo) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		//"root", "root", "127.0.0.1:3306", "db_lqyy")
		info.Root, info.Password, info.Address, info.Database)
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log4.Panic(fmt.Sprintf("mysql连接失败... %s err=>%s", dsn, err.Error()))
		return
	}
	t.db = _db
	// Ensure the table exists
	/*
		t.db.AutoMigrate(
			&dbtable.DBUser{},
		)
	*/
	// 执行原生SQL，设置自增的起始值为100
	//t.db.Exec("ALTER TABLE db_users AUTO_INCREMENT = 4575")
}
func (t *MysqlClient) GetDBSession() *gorm.DB {
	// 开始一个 Session
	session := t.db.Session(&gorm.Session{})
	return session
	//return t.db
}

func (t *MysqlClient) AutoMigrate(dst ...interface{}) {
	t.db.AutoMigrate(dst...)
}
func (t *MysqlClient) Exec(sql string) {
	//t.db.Exec("ALTER TABLE db_users AUTO_INCREMENT = 4575")
	t.db.Exec(sql)
}
