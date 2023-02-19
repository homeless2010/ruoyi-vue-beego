package routers

import (
	"github.com/beego/beego/v2/server/web"
	"ruoyi-vue-beego/controllers/system"
)

func init() {
	prefix := ""
	system := web.NewNamespace(prefix+"/v1",
		web.NSRouter("/login", &system.SysLoginController{}, "post:Login"),
		//web.NSRouter("/logout", &system.SysLoginController{}, "delete:Logout"),
	)
	web.AddNamespace(system)
}
