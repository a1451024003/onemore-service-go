package controllers

import (
	"fmt"
	"onemore-service-go/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 热更
type HotChcpController struct {
	beego.Controller
}

// URLMapping ...
func (c *HotChcpController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description  新增热更信息
// @Success 0 {object} models.HotChcp
// @Failure 1003 添加失败
// @router / [post]
func (c *HotChcpController) Post() {
	name := c.GetString("name")
	ios_identifier := c.GetString("ios_identifier")
	android_identifier := c.GetString("android_identifier")
	update := c.GetString("update")
	manifest := c.GetString("manifest")
	content_url := c.GetString("content_url")
	release := c.GetString("release")

	u := models.HotChcp{
		Name:              name,
		IosIdentifier:     ios_identifier,
		AndroidIdentifier: android_identifier,
		Update:            update,
		Manifest:          manifest,
		ContentUrl:        content_url,
		Release:           release,
	}

	if err := u.AddHot(); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, "", "添加成功"}
	} else {
		fmt.Println(err)
		c.Data["json"] = JSONStruct{"error", 1003, "", "添加失败"}
	}

	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description 获取全部信息
// @Success 0 {object} models.HotChcp
// @Failure 1001 获取失败
// @router / [get]
func (c *HotChcpController) GetAll() {
	var f *models.HotChcp
	var prepage int = 20
	var page int = 1

	if v, err := c.GetInt("per_page"); err == nil {
		prepage = v
	}
	if v, err := c.GetInt("page"); err == nil {
		page = v
	}

	if l, err := f.GetAll(prepage, page); err == nil {
		mystruct := JSONStruct{"success", 0, l, "获取成功"}
		c.Data["json"] = mystruct
	} else {
		mystruct := JSONStruct{"error", 1001, l, "获取失败"}
		c.Data["json"] = mystruct
	}

	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description 编辑热更信息
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.HotChcp	true		"body for Hot_chap content"
// @Success 0 {object} models.HotChcp
// @Failure 1003 编辑失败
// @router /:id [put]
func (c *HotChcpController) Put() {
	name := c.GetString("name")
	ios_identifier := c.GetString("ios_identifier")
	android_identifier := c.GetString("android_identifier")
	update := c.GetString("update")
	manifest := c.GetString("manifest")
	content_url := c.GetString("content_url")
	release := c.GetString("release")

	//验证参数是否为空
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	w := models.HotChcp{Id: id,
		Name:              name,
		IosIdentifier:     ios_identifier,
		AndroidIdentifier: android_identifier,
		Update:            update,
		Manifest:          manifest,
		ContentUrl:        content_url,
		Release:           release,
	}
	if err := w.Save(); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, "", "编辑成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1003, "", "编辑失败"}
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description 删除热更信息
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 0 {string} delete success!
// @Failure 1002 热更信息删除失败
// @Failure 1003 热更信息不存在
// @router /:id [delete]
func (c *HotChcpController) Delete() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	f := &models.HotChcp{Id: id}
	if _, err := f.Delete(); err == orm.ErrNoRows {
		c.Data["json"] = JSONStruct{"error", 1003, err, "热更信息不存在"}
	} else if err == nil {
		c.Data["json"] = JSONStruct{"success", 0, err, "热更信息删除成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1002, err, "热更信息删除失败"}
	}

	c.ServeJSON()
}

// GetNew ...
// @Title GetNew
// @Description 获取最新
// @Success 0 {object} models.HotChcp
// @Failure 1001 获取最新热更信息失败
// @router /getNew [get]
func (c *HotChcpController) GetNew() {
	var f *models.HotChcp
	if l, err := f.GetNew(); err == nil {

		mystruct := JSONStruct{"success", 0, l, "获取成功"}
		c.Data["json"] = mystruct
	} else {
		mystruct := JSONStruct{"error", 1001, l, "获取失败"}
		c.Data["json"] = mystruct
	}

	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description 获取详情
// @Success 0 {object} models.HotChcp
// @Failure 1001 获取详情失败
// @router /:id [get]
func (c *HotChcpController) GetOne() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	f := &models.HotChcp{Id: id}
	if l, err := f.GetOne(id); err == nil {
		mystruct := JSONStruct{"success", 0, l, "获取成功"}

		c.Data["json"] = mystruct
	} else {
		mystruct := JSONStruct{"error", 1001, "", "获取失败"}

		c.Data["json"] = mystruct
	}

	c.ServeJSON()
}

// GetNew ...
// @Title GetNew
// @Description 获取最新
// @Success 0 {object} models.HotChcp
// @Failure 1001 获取最新热更信息失败
// @router /GetNewInfo [get]
func (c *HotChcpController) GetNewInfo() {
	var f *models.HotChcp
	if l, err := f.GetNewInfo(); err == nil {
		mystruct := JSONStruct{"success", 0, l["manifest"], "获取成功"}
		c.Data["json"] = mystruct
	} else {
		mystruct := JSONStruct{"error", 1001, l, "获取失败"}
		c.Data["json"] = mystruct
	}

	c.ServeJSON()
}
