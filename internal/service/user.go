package service

import (
	"context"
	"gin-plus/internal/model"
	"gin-plus/internal/model/entity"
	"gin-plus/internal/service/dao"
)

var User = sUser{}

type sUser struct {
}

func (s *sUser) GerList(ctx context.Context) (out *model.GetUserListOutput, err error) {
	var users []entity.User
	if err = dao.User.Ctx(ctx).Find(&users).Error; err != nil {
		return
	}

	arr := make(model.GetUserListOutput, len(users))
	for i, u := range users {
		arr[i] = &struct {
			Id   int
			Name string
		}{Id: u.Id, Name: u.Name}
	}

	out = &arr
	return
}
