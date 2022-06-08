package model

type User struct {
	Id   int
	Name string
}

type GetUserListOutput []*struct {
	Id   int
	Name string
}
