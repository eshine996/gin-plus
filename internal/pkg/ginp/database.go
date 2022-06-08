package ginp

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMysqlDatabase() *gorm.DB {
	var host, port, user, password, database string
	if host = Config.GetString("mysql.host"); host == "" {
		panic("config.mysql.host cannot be empty")
	}
	if port = Config.GetString("mysql.port"); port == "" {
		panic("config.mysql.port cannot be empty")
	}
	if user = Config.GetString("mysql.user"); user == "" {
		panic("config.mysql.user cannot be empty")
	}
	if password = Config.GetString("mysql.password"); password == "" {
		panic("config.mysql.password cannot be empty")
	}

	if database = Config.GetString("mysql.database"); password == "" {
		panic("config.mysql.password cannot be empty")
	}

	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)

	mysqlConfig := mysql.Config{
		DSN:                      mysqlDSN, // DSN data source name
		DefaultStringSize:        256,      // string 类型字段的默认长度
		DisableDatetimePrecision: true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:   true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true, // 根据当前 MySQL 版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic(err.Error())

	}

	return db
}
