package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gogochat/dao"
	"gogochat/models"
	"gogochat/service"
	"gogochat/tool"
	"strconv"
)

type UserController struct {
}

func (uc *UserController) Router(engine *gin.Engine) {
	engine.GET("/user/create", uc.createUser)
	engine.GET("/user/update", uc.update)
	engine.DELETE("/user/delete", uc.deleteUser)
	engine.GET("/user/login", uc.login)
}

func (uc *UserController) createUser(context *gin.Context) {
	name := context.Query("name")
	password := context.Query("password")
	rePassword := context.Query("re_password")
	email := context.Query("email")
	phone := context.Query("phone")

	if password != rePassword {
		tool.Failed(context, "两次密码必须相同")
		return
	}

	//校验
	userDao := dao.NewUserDao()
	userValidate := userDao.FindUserByPhone(phone)
	if userValidate.ID != 0 {
		tool.Failed(context, "该电话已存在")
		return
	}

	userValidate = userDao.FindUserByEmail(email)
	if userValidate.ID != 0 {
		tool.Failed(context, "该邮箱已存在")
		return
	}
	userValidate = userDao.FindUserByName(name)
	if userValidate.ID != 0 {
		tool.Failed(context, "该名字已存在")
		return
	}

	hashPwd := tool.HashPasswordSha256(password)

	user := models.UserBasic{Password: hashPwd, Email: email, Phone: phone, Name: name}
	userService := service.NewUserService()
	err := userService.CreateUser(&user)
	if err != nil {
		tool.Failed(context, err)
		return
	}

	tool.Success(context, "创建成功")

}

func (uc *UserController) update(context *gin.Context) {
	id := context.Query("id")
	password := context.Query("password")
	email := context.Query("email")
	phone := context.Query("phone")
	idConv, err := strconv.Atoi(string(id))
	if err != nil {
		tool.Failed(context, "id解析失败")
		return
	}
	hashPwd := tool.HashPasswordSha256(password)
	user := models.UserBasic{Password: hashPwd, Email: email, Phone: phone}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		tool.Failed(context, "邮箱电话格式错误")
		return
	}

	userService := service.NewUserService()
	err = userService.UpdateUser(idConv, &user)
	if err != nil {
		tool.Failed(context, err)
		return
	}

	tool.Success(context, "更新成功")
}

func (uc *UserController) deleteUser(context *gin.Context) {
	id := context.Query("id")
	idConv, err := strconv.Atoi(string(id))
	if err != nil {
		tool.Failed(context, "id解析失败")
		return
	}

	userService := service.NewUserService()
	err = userService.DeleteUser(idConv)
	if err != nil {
		tool.Failed(context, err)
		return
	}

	tool.Success(context, "删除成功")
}

func (uc *UserController) login(context *gin.Context) {
	name := context.Query("name")
	pwd := context.Query("password")

	userDao := dao.NewUserDao()
	user := userDao.FindUserByName(name)
	if user.ID == 0 {
		tool.Failed(context, "用户不存在")
		return
	}

	isValid := tool.ValidHashPassword(pwd, user.Password)
	if !isValid {
		tool.Failed(context, "登录密码错误")
		return
	}
	tool.Success(context, "登录成功")
}
