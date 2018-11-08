package controllers

import (
	"onemore-service-go/models"

	"github.com/astaxie/beego"
)

//配置中心
type SystemController struct {
	beego.Controller
}

// Setting ...
// @Title 读取数据库配置
// @Description 读取数据库配置
// @Param	type		    path 	int	true		"环境 1.local 2.dev 3.test 4.prod"
// @Param	service		path 	int	true		"数据库服务"
// @Success 200 {object} models.Config
// @Failure 403 :id is empty
// @router / [get]
func (c *SystemController) Settings() {
	ty, _ := c.GetInt("type")
	service, _ := c.GetInt("service")
	v, err := models.Configs(ty, service)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, err.Error(), "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// Post ...
// @Title 数据库配置
// @Description 数据库配置
// @Param	config		path 	int	true		"配置"
// @Success 200 {object} models.Config
// @Failure 403 :id is empty
// @router / [post]
func (c *SystemController) Post() {
	ty, _ := c.GetInt("type")
	service, _ := c.GetInt("service")
	config := c.GetString("config")
	err := models.SystemPost(ty, service, config)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, err.Error(), "配置失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "配置成功"}
	}
	c.ServeJSON()
}
