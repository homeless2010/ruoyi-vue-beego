package system

import (
	"github.com/beego/beego/v2/client/orm"
	"ruoyi-vue-beego/models"
)

type SysUser struct {
	UserId      int64  `json:"userId,string" orm:"column(user_id);pk"`
	DeptId      *int64 `json:"deptId,string" orm:"column(dept_id)"`
	UserName    string `json:"userName" orm:"column(user_name)""`
	NickName    string `json:"nickName" orm:"column(nick_name)"`
	UserType    string `json:"userType" orm:"column(user_type)"`
	Email       string `json:"email" orm:"column(email)""`
	Avatar      string `json:"avatar" orm:"column(avatar)""`
	Phonenumber string `json:"phonenumber" orm:"column(phonenumber)"`
	Sex         string `json:"sex" orm:"column(sex)"`
	Password    string `json:"password" orm:"column(password)"`
	Status      string `json:"status" orm:"column(status)"`
	Remark      string `json:"remake" orm:"column(remark)"`
	models.BaseEntity
}

func init() {
	orm.RegisterModel(new(SysUser))
}
