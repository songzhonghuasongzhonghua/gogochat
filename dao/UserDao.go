package dao

import (
	"gogochat/models"
	"gogochat/tool"
)

type UserDao struct {
	*tool.Orm
}

func NewUserDao() *UserDao {
	return &UserDao{tool.DBEngine}
}

func (ud *UserDao) CreateUser(user *models.UserBasic) error {

	tx := ud.Create(user)
	return tx.Error
}

func (ud *UserDao) UpdateUser(id int, user *models.UserBasic) error {

	tx := ud.Where("id = ?", id).Updates(user)
	return tx.Error
}

func (ud *UserDao) DeleteUser(id int) error {

	tx := ud.Where("id = ?", id).Delete(&models.UserBasic{})
	return tx.Error
}

func (ud *UserDao) FindUserByName(name string) *models.UserBasic {
	user := models.UserBasic{}
	ud.Where("name = ?", name).Find(&user)
	return &user
}

func (ud *UserDao) FindUserByPhone(phone string) *models.UserBasic {
	user := models.UserBasic{}
	ud.Where("phone = ?", phone).Find(&user)
	return &user
}

func (ud *UserDao) FindUserByEmail(email string) *models.UserBasic {
	user := models.UserBasic{}
	ud.Where("email = ?", email).Find(&user)
	return &user
}
