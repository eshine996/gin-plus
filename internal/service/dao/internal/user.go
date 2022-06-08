package internal

import (
	"context"
	"gin-plus/internal/model/entity"
	"gin-plus/internal/pkg/ginp"
	"gorm.io/gorm"
)

var User = doUser{}

type doUser struct {
	entity.User
}

func (d *doUser) TableName() string {
	return "test"
}

func (d *doUser) Mysql() *gorm.DB {
	return ginp.Mysql
}

func (d *doUser) Ctx(ctx context.Context) *gorm.DB {
	return d.Mysql().Model(d).WithContext(ctx)
}
