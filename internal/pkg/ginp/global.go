package ginp

import (
	"gorm.io/gorm"
)

var Config *iConfig
var Mysql *gorm.DB

//进行初始化配置文件及数据库orm对象
func init() {
	if Config == nil {
		Config = NewConfig()
		if err := Config.ReadData(); err != nil {
			panic(err)
		}
	}
	if Mysql == nil {
		Mysql = GetMysqlDatabase()
	}
}
