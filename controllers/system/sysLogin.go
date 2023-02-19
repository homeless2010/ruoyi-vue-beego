package system

import (
	"github.com/beego/beego/v2/core/logs"
	"ruoyi-vue-beego/controllers"
	"ruoyi-vue-beego/service"
	"strings"
)

type SysLoginController struct {
	controllers.BaseController
}

func (c *SysLoginController) Login() {
	username := strings.Trim(c.GetString("username"), " ")
	password := c.GetString("password")
	token, err := service.GetUserService().Login(username, password)
	if err != nil {
		logs.Error("login error", err)
	}
	c.SuccessWithData(token)
}
