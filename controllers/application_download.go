package controllers

import (
	"log"
	"onemore-service-go/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// 应用下载
type ApplicationDownloadController struct {
	beego.Controller
}

// URLMapping ...
func (c *ApplicationDownloadController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create 应用下载-添加
// @Param	icon				formData 	string			true		"图标"
// @Param	name				formData 	name			true		"名称"
// @Param	download_url		formData 	download_url	true		"下载地址"
// @Param	type				formData 	int				true		"类型"
// @Param	package_name		formData 	string			true		"应用包名"
// @Param	version_code		formData 	string			true		"应用版本号"
// @Param	size				formData 	string			true		"应用apk包大小（MB）"
// @Success 201 {object} models.ApplicationDownload
// @Failure 403 body is empty
// @router / [post]
func (c *ApplicationDownloadController) Post() {
	var v models.ApplicationDownload
	Icon := c.GetString("icon")
	Name := c.GetString("name")
	DownloadUrl := c.GetString("download_url")
	Type, _ := c.GetInt("type")
	PackageName := c.GetString("package_name")
	VersionCode := c.GetString("version_code")
	Size := c.GetString("size")
	v.Icon = Icon
	v.Name = Name
	v.DownloadUrl = DownloadUrl
	v.Type = Type
	v.PackageName = PackageName
	v.VersionCode = VersionCode
	v.Size = Size
	valid := validation.Validation{}
	valid.Required(v.Icon, "icon").Message("图标不能为空")
	valid.Required(v.Name, "name").Message("应用名称不能为空")
	valid.Required(v.DownloadUrl, "download_url").Message("应用下载地址不能为空")
	valid.Required(v.Type, "type").Message("类型不能为空")
	valid.Required(v.PackageName, "package_name").Message("包名称不能为空")
	valid.Required(v.VersionCode, "version_code").Message("类型不能为空")
	valid.Required(v.Size, "size").Message("类型不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		v := models.AddApplicationDownload(&v)
		if v == nil {
			c.Data["json"] = JSONStruct{"error", 1003, nil, "保存失败"}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, nil, "保存成功"}
		}
		c.ServeJSON()
	}
}

// GetOne ...
// @Title GetOne
// @Description get ApplicationDownload by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ApplicationDownload
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ApplicationDownloadController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get 应用下载-列表
// @Param	type    	query	int	 true		"类型：1儿歌，2游戏"
// @Param	page     query	int	 false		"页数"
// @Param	per_page query	int	 false		"每页显示条数"
// @Success 200 {object} models.ApplicationDownload
// @Failure 403
// @router / [get]
func (c *ApplicationDownloadController) GetAll() {
	var prepage int = 20
	var page int
	var ty int
	if v, err := c.GetInt("per_page"); err == nil {
		prepage = v
	}
	if v, err := c.GetInt("page"); err == nil {
		page = v
	}
	if v, err := c.GetInt("type"); err == nil {
		ty = v
	} else {
		valid := validation.Validation{}
		valid.Required(ty, "Type").Message("请补全参数")
		if valid.HasErrors() {
			log.Println(valid.Errors)
			c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
			c.ServeJSON()
		}
	}
	v := models.GetAllApplicationDownload(ty, page, prepage)
	if v == nil {
		c.Data["json"] = JSONStruct{"error", 1005, v, "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update 应用下载-修改
// @Param	id		    		path 		int	            true		"应用编号"
// @Param	icon				formData 	string			true		"图标"
// @Param	name				formData 	name				true		"名称"
// @Param	download_url		formData 	download_url		true		"下载地址"
// @Param	type				formData 	int				true		"类型"
// @Param	package_name		formData 	string			true		"应用包名"
// @Param	version_code		formData 	string			true		"应用版本号"
// @Param	size				formData 	string			true		"应用apk包大小（MB）"
// @Success 200 {object} models.ApplicationDownload
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ApplicationDownloadController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ApplicationDownload{Id: id}
	Icon := c.GetString("icon")
	Name := c.GetString("name")
	DownloadUrl := c.GetString("download_url")
	Type, _ := c.GetInt("type")
	PackageName := c.GetString("package_name")
	VersionCode := c.GetString("version_code")
	Size := c.GetString("size")
	v.Icon = Icon
	v.Name = Name
	v.DownloadUrl = DownloadUrl
	v.Type = Type
	v.PackageName = PackageName
	v.VersionCode = VersionCode
	v.Size = Size
	valid := validation.Validation{}
	valid.Required(v.Icon, "Icon").Message("图标不能为空")
	valid.Required(v.Name, "Name").Message("应用名称不能为空")
	valid.Required(v.DownloadUrl, "DownloadUrl").Message("应用下载地址不能为空")
	valid.Required(v.Type, "Type").Message("类型不能为空")
	valid.Required(v.PackageName, "package_name").Message("包名称不能为空")
	valid.Required(v.VersionCode, "version_code").Message("类型不能为空")
	valid.Required(v.Size, "size").Message("类型不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		v := models.UpdateApplicationDownloadById(&v)
		if v == nil {
			c.Data["json"] = JSONStruct{"error", 1003, nil, "编辑失败"}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, nil, "编辑成功"}
		}
		c.ServeJSON()
	}
}

// Delete ...
// @Title Delete
// @Description delete 应用下载-删除
// @Param	id		path 	int	true	 "应用编号"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ApplicationDownloadController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.DeleteApplicationDownload(id)
	if v == nil {
		c.Data["json"] = JSONStruct{"error", 1003, nil, "删除失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "删除成功"}
	}
	c.ServeJSON()
}
