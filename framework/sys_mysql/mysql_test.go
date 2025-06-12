package sys_mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  uint8
}

/*
*
db := // your gorm database instance
tx := db.Begin()

// perform database operations
// ...

// rollback transaction if there is an error

	if err != nil {
	    tx.Rollback()
	}

// commit transaction if everything is successful
tx.Commit()
*/
func mainxx() {
	dsn := "user:password@tcp(127.0.0.1:3306)/database_name"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// 开始一个 Session
	session := db.Session(&gorm.Session{})

	// 添加数据
	user1 := &User{Name: "John", Age: 25}
	user2 := &User{Name: "Mary", Age: 30}
	if err := session.Create(user1).Error; err != nil {
		panic(err)
	}
	if err := session.Create(user2).Error; err != nil {
		panic(err)
	}

	// 删除数据
	if err := session.Where("name = ?", "John").Delete(&User{}).Error; err != nil {
		panic(err)
	}

	// 查询数据
	var users []User
	if err := session.Where("age > ?", 25).Find(&users).Error; err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}
