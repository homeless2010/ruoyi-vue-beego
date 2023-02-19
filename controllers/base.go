package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"sync"
)

type BaseController struct {
	web.Controller
	Wg sync.WaitGroup
}

var L = logs.GetLogger()

func (c *BaseController) Result(code uint, msg string) {
	c.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}
func (c *BaseController) ResultWithData(code uint, msg string, data interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}
func (c *BaseController) Success() {
	c.Data["json"] = map[string]interface{}{
		"code": http.StatusOK,
		"msg":  "操作成功",
	}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}

func (c *BaseController) SuccessWithMsg(msg string) {
	c.Data["json"] = map[string]interface{}{
		"code": http.StatusOK,
		"msg":  msg,
	}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}
func (c *BaseController) SuccessWithData(data interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code": http.StatusOK,
		"data": data,
	}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}

func (c *BaseController) SuccessWithMsgData(msg string, data interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}

func (c *BaseController) Error() {
	c.Data["json"] = map[string]interface{}{
		"code": http.StatusInternalServerError,
		"msg":  "操作失败",
	}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}

func (c *BaseController) ErrorWithMsg(msg string) {
	c.Data["json"] = map[string]interface{}{
		"code": http.StatusInternalServerError,
		"msg":  msg,
	}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}
func (c *BaseController) ErrorWithMsgData(msg string, data interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code": http.StatusInternalServerError,
		"msg":  msg,
		"data": data,
	}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}
func (c *BaseController) ErrorWithCodeMsg(code uint, msg string) {
	c.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}
