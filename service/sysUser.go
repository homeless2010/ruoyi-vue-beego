package service

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"ruoyi-vue-beego/models/system"
)

var service *userService

type userService struct{}

func init() {
	service = &userService{}
}

func GetUserService() *userService {
	return service
}

func (userService *userService) Login(username string, password string) (token string, err error) {
	user, err := userService.LoadUserByUsername(username)
	fmt.Println(user)
	return "jwt", err
}

func (userService *userService) LoadUserByUsername(username string) (sysUser *system.SysUser, err error) {
	o := orm.NewOrm()
	user := system.SysUser{UserName: username}
	err = o.Read(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}
