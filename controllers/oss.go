package controllers

import (
	"onemore-service-go/models"

	"github.com/astaxie/beego"
	"fmt"
)

// oss
type OssController struct {
	beego.Controller
}

type JSONStruct struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

func (c *OssController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("GetSts", c.GetSts)
}

// @Title oss直传token
// @Description oss直传token
// @Param   path     		query    string  	 false       "路径"
// @Success 0 {json} JSONStruct
// @router / [get]
func (c *OssController) Get() {
	path := c.GetString("path")
	res := models.GetPolicyToken(path)

	c.Data["json"] = JSONStruct{"success", 0, res, "获取成功"}

	c.ServeJSON()
}

// @Title 获取oss配置
// @Description 获取oss配置
// @Success 0 {string} success
// @router /oss [get]
func (c *OssController) GetOss() {
	res := make(map[string]string)
	res["Endpoint"] = beego.AppConfig.String("OSS_WEB_SERVER")
	res["Keyid"] = beego.AppConfig.String("OSS_KEY_ID")
	res["Secret"] = beego.AppConfig.String("OSS_KEY_SECRET")
	res["Ossbucket"] = beego.AppConfig.String("OSS_BUCKET")
	c.Data["json"] = JSONStruct{"success", 0, res, "获取成功"}
	c.ServeJSON()
}

// @Title 获取sts
// @Description 获取sts
// @Success 0 {json} JSONStruct
// @Failure 1005 获取失败
// @router /sts [get]
func (c *OssController) GetSts() {
	if v, err := models.GetSecurityToken(); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	} else {
		fmt.Println(err)
		c.Data["json"] = JSONStruct{"error", 1005, "", "获取失败"}
	}

	c.ServeJSON()
}
