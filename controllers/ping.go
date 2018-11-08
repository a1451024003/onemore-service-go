package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

// ping
type PingController struct {
	beego.Controller
}

// URLMapping ...
func (c *PingController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @Title 测试服务是否正常
// @Description 测试服务是否正常
// @Success 0 {json} JSONStruct
// @router / [get]
func (c *PingController) Get() {
	var balance [2]string

	balance[0] = beego.AppConfig.String("appname")
	balance[1] = time.Now().Format("2006-01-02 15:04:05")

	c.Data["json"] = JSONStruct{"success", 0, balance, "获取成功"}
	c.ServeJSON()
}
