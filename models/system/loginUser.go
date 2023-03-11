package system

import "time"

type LoginUser struct {
	UserId        int64     `json:"userId"`
	DeptId        *int64    `json:"deptId"`
	UserName      string    `json:"userName"`
	LoginTime     time.Time `json:"loginTime"`
	ExpireTime    time.Time `json:"expireTime"`
	Ipaddr        string    `json:"ipaddr"`
	LoginLocation string    `json:"loginLocation"`
	Browser       string    `json:"browser"`
	Permissions   []string  `json:"permissions"`
	SysUser       SysUser
}
