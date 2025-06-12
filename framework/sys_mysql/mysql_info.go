/*
*
added by yh @ 2023/4/27 18:16
注意:
*/
package sys_mysql

/*
19:51>>>在Golang中使用GORM连接MySQL的步骤如下：

1. 安装MySQL驱动

在终端中执行以下命令安装MySQL驱动：

go get -u github.com/go-sql-driver/mysql

2. 导入GORM和MySQL驱动

在项目中导入GORM和MySQL驱动：

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

3. 建立数据库连接

使用Open函数建立数据库连接：

dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

其中，dsn参数包含数据库连接信息，包括用户名、密码、主机地址、端口、数据库名、字符集等。

4. 定义数据模型

定义数据模型并使用AutoMigrate函数自动创建表：

type User struct {
    gorm.Model
    Name string
    Age  int
}

db.AutoMigrate(&User{})

5. 数据库操作

使用db变量进行数据库操作，例如插入数据：

user := User{Name: "test", Age: 18}
db.Create(&user)

查询数据：

var users []User
db.Find(&users)

更新数据：

db.Model(&user).Update("Age", 20)

删除数据：

db.Delete(&user)

使用GORM连接MySQL可以方便地进行数据库操作，同时也支持事务、链式操作等高级特性。
*/
