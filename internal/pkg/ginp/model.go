package ginp

import (
	"context"
	"gorm.io/gorm"
)

type BaseModel struct{}

func (b *BaseModel) Mysql() *gorm.DB {
	return Mysql
}

func (b *BaseModel) Ctx(ctx context.Context) *gorm.DB {
	return b.Mysql().Model(b).WithContext(ctx)
}
