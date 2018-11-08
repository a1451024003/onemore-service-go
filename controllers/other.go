package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"io"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

//other
type OtherController struct {
	beego.Controller
}

// @Title 密码加密
// @Description 密码加密
// @Param   pwd     query    string  	 true       "密码"
// @Success 0 {string} success
// @router /encryptPwd [get]
func (c *OtherController) GetPwd() {
	pwd := c.GetString("pwd")
	valid := validation.Validation{}
	valid.Required(pwd, "pwd").Message("密码不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		h := md5.New()
		io.WriteString(h, pwd)
		io.WriteString(h, "dameifenglin")
		md5String := h.Sum(nil)
		ml := make(map[string]interface{})
		NewPwd := base64.StdEncoding.EncodeToString(md5String)
		ml["encryptPwd"] = NewPwd
		c.Data["json"] = JSONStruct{"success", 0, ml, "获取成功"}
		c.ServeJSON()
	}
}
