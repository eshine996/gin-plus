package ginp

import (
	"gorm.io/gorm"
)

var Config *inConfig
var Mysql *gorm.DB

func init() {
	if Config == nil {
		Config = NewConfig()
	}
	if Mysql == nil {
		Mysql = GetMysqlDatabase()
	}
}
