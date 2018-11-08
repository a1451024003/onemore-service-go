package controllers

import (
	"onemore-service-go/models"
	"strconv"

	"github.com/astaxie/beego"
)

//租户
type TenantController struct {
	beego.Controller
}

// Tenant ...
// @Title 租户详情
// @Description 租户详情
// @Param	tenant_id		path 	int	true		"租户id"
// @Success 200 {object} models.Tenant
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TenantController) Tenant() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTenant(id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, err.Error(), "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// TenantList ...
// @Title 租户列表
// @Description 租户列表
// @Success 200 {object} models.Tenant
// @Failure 403 :id is empty
// @router / [get]
func (c *TenantController) TenantList() {
	v, err := models.TenantList()
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, err.Error(), "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// Post ...
// @Title 添加租户
// @Description 添加租户
// @Param	name		    path 	string	true		"租户名字"
// @Param	is_free		    path 	int	    true		"是否免费：0免费1付费"
// @Param	fee_type		path 	int	    true		"付费类型：1包年2包季度3包月"
// @Param	fee_period		path 	string	true		"租期时长"
// @Param	start_time		path 	string	true		"开始时间"
// @Param	end_time		path 	string	true		"结束时间"
// @Param	rent_type		path 	int 	true		"租借状态0正常1到期未付费2到期停"
// @Success 200 {object} models.Tenant
// @Failure 403 :id is empty
// @router / [post]
func (c *TenantController) Post() {
	name := c.GetString("name")
	is_free, _ := c.GetInt("is_free")
	fee_type, _ := c.GetInt("fee_type")
	fee_period := c.GetString("fee_period")
	start_time := c.GetString("start_time")
	end_time := c.GetString("end_time")
	rent_type, _ := c.GetInt("rent_type")
	err := models.Post(name, is_free, fee_period, fee_type, start_time, end_time, rent_type)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, err.Error(), "添加失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "添加成功"}
	}
	c.ServeJSON()
}

// Delete ...
// @Title 删除租户
// @Description 删除租户
// @Param	tenant_id		    path 	int	true		"租户id"
// @Success 200 {object} models.Tenant
// @Failure 403 :id is empty
// @router /:id [delete]
func (c *TenantController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	err := models.Delete(id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, err.Error(), "删除失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "删除成功"}
	}
	c.ServeJSON()
}

// Update ...
// @Title 编辑租户
// @Description 编辑租户
// @Param	tenant_id		    path 	int	true		"租户id"
// @Success 200 {object} models.Tenant
// @Failure 403 :id is empty
// @router /:id [put]
func (c *TenantController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	name := c.GetString("name")
	is_free, _ := c.GetInt("is_free")
	fee_type, _ := c.GetInt("fee_type")
	fee_period := c.GetString("fee_period")
	start_time := c.GetString("start_time")
	end_time := c.GetString("end_time")
	rent_type, _ := c.GetInt("rent_type")
	err := models.Update(id, name, is_free, fee_period, fee_type, start_time, end_time, rent_type)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, err.Error(), "编辑失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "编辑成功"}
	}
	c.ServeJSON()
}
