package service

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"ruoyi-vue-beego/models/system"
)

var userServiceImpl *userService

type userService struct{}

func init() {
	userServiceImpl = &userService{}
}

func GetUserService() *userService {
	return userServiceImpl
}

func (userService *userService) Login(username string, password string) (token string, err error) {
	user, err := userService.LoadUserByUsername(username)
	fmt.Println(user)
	t, err := GetTokenService().createToken(user)
	return t, err
}

func (userService *userService) LoadUserByUsername(username string) (sysUser system.LoginUser, err error) {
	o := orm.NewOrm()
	user := system.SysUser{UserName: username}
	err = o.Read(user)
	if err != nil {
		return system.LoginUser{}, err
	}
	loginUser := createLoginUser(user)
	return loginUser, err
}

func createLoginUser(user system.SysUser) system.LoginUser {
	//TODO permissions
	loginUser := system.LoginUser{
		UserId:  user.UserId,
		DeptId:  user.DeptId,
		SysUser: user}
	return loginUser
}
