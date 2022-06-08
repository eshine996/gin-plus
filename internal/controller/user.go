package controller

import (
	"context"
	v1 "gin-plus/api/v1"
	"gin-plus/internal/model"
	"gin-plus/internal/service"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {
	var out *model.GetUserListOutput
	if out, err = service.User.GerList(ctx); err != nil {
		return
	}

	arr := make(v1.UserRes, len(*out))
	for i, o := range *out {
		arr[i] = o
	}

	res = &arr
	return
}
