package v1

import "gin-plus/internal/pkg/ginp"

type UserReq struct {
	ginp.Api `path:"/user" method:"GET"`
	Id       int    `binding:"required"`
	Name     string `binding:"required"`
}

type UserRes []*struct {
	Id   int
	Name string
}
