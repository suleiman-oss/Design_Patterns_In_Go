package service

type UserService struct {
	Db Database
}

func (u *UserService) Add() {
	q := u.Db.Query()
	q.ExecuteQuery()
}
