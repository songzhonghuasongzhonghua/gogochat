package service

import (
	"gogochat/dao"
	"gogochat/models"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(user *models.UserBasic) error {
	newUserDao := dao.NewUserDao()
	return newUserDao.CreateUser(user)
}

func (us *UserService) UpdateUser(id int, user *models.UserBasic) error {
	newUserDao := dao.NewUserDao()

	return newUserDao.UpdateUser(id, user)

}

func (us *UserService) DeleteUser(id int) error {
	newUserDao := dao.NewUserDao()

	return newUserDao.DeleteUser(id)

}
