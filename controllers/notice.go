package controllers

import (
	"encoding/json"
	"onemore-service-go/models"
	"strconv"

	"github.com/hprose/hprose-golang/rpc"

	"github.com/astaxie/beego"
	"onemore-service-go/utils"
)

var User *UserService

// NoticeController operations for Notice
type NoticeController struct {
	beego.Controller
}

type UserService struct {
	GetOneByUserId func(userId int) (interface{}, error)
	TeacherNotice  func(class_type int, kindergarten_id int) (interface{}, error)
	StudentNotice  func(class_type int, kindergarten_id int) (interface{}, error)
	Pub            func(int, int, string) error
	GetBabyInfo    func(id int) (map[string]interface{}, error)
	GetUsers       func(types int) ([]map[string]interface{}, error)
}

// URLMapping ...
func (c *NoticeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Delete", c.Delete)
}

// Post 发布通知
// @Title Post 发布通知
// @Description 发布通知
// @Param	body		body 	models.Notice	true		"body for Notice content"
// @Success 0 {int} models.Notice
// @Failure 403 body is empty
// @router / [post]
func (c *NoticeController) Post() {
	var v models.Request
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddNotice(&v); err == nil {
			c.Data["json"] = models.JSONStruct{"success", 0, nil, "保存成功！"}
		} else {
			c.Data["json"] = models.JSONStruct{"error", 1003, nil, err.Error()}
		}
	} else {

		c.Data["json"] = models.JSONStruct{"error", 1001, nil, "请输入正确格式的json格式数据"}
	}
	c.ServeJSON()
}

// GetOne 通知详情
// @Title 通知详情
// @Description 通知详情
// @Param	id		path 	string	true		"通知ID"
// @Success 0 {object} models.Notice
// @Failure 403 :id is empty
// @router /:id [get]
func (c *NoticeController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetNoticeById(id)
	user_id, _ := c.GetInt("user_id")
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1005, nil, "获取失败！"}
	} else {
		if v == nil {
			c.Data["json"] = models.JSONStruct{"success", 1002, nil, "没有相关数据"}
		} else {
			//更改阅读状态
			models.UpdateNoticeById(id, user_id)
			c.Data["json"] = models.JSONStruct{"success", 0, v, "获取成功！"}
		}
	}
	c.ServeJSON()
}

// GetAll 通知列表
// @Title Get All 通知列表
// @Description 通知列表
// @Success 0 {object} models.Notice
// @Failure 403
// @router / [get]
func (c *NoticeController) GetAll() {
	ty, _ := c.GetInt("type")
	search := c.GetString("search")
	user_id, _ := c.GetInt("user_id")
	notice_type, _ := c.GetInt("notice_type")
	l, err := models.GetAllNotice(user_id, ty, search, notice_type)
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1005, nil, "获取失败！"}
	} else {
		if l == nil {
			c.Data["json"] = models.JSONStruct{"success", 1002, nil, "没有相关数据"}
		} else {
			c.Data["json"] = models.JSONStruct{"success", 0, l, "获取成功！"}
		}
	}
	c.ServeJSON()
}

// 系统通知列表
// @Title 系统通知列表
// @Description 系统通知列表
// @Success 0 {object} models.Notice
// @Failure 403
// @router /notice [get]
func (c *NoticeController) GetNotice() {
	page, _ := c.GetInt("page")
	prepage, _ := c.GetInt("per_page", 20)
	ty, _ := c.GetInt("type")
	search := c.GetString("search")
	user_id, _ := c.GetInt("user_id")
	notice_type, _ := c.GetInt("notice_type")
	l, err := models.GetNotice(user_id, ty, search, notice_type, page, prepage)
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1005, nil, "没有相关数据"}
	} else {
		c.Data["json"] = models.JSONStruct{"success", 0, l, "获取成功！"}
	}
	c.ServeJSON()
}

// Delete 删除通知
// @Title 删除通知
// @Description 删除通知
// @Param	id		path 	string	true		"通知ID"
// @Success 0 {string} delete success!
// @Failure 403 id is empty
// @router /del_notice [post]
func (c *NoticeController) Delete() {
	id, _ := c.GetInt("notice_id")
	user_id, _ := c.GetInt("user_id")
	err := models.DeleteNotice(id, user_id)
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1008, nil, "删除失败"}
	} else {
		c.Data["json"] = models.JSONStruct{"success", 0, nil, "删除成功！"}
	}
	c.ServeJSON()
}

// PostSystem 发布系统通知
// @Title PostSystem 发布系统通知
// @Description 发布系统通知
// @Param	title		formData 	string	true		"标题"
// @Param	content		formData 	string	true		"内容"
// @Param	notice_type	formData 	string	true		"通知分类"
// @Param	type		formData 	string	true		"选择分类"
// @Success 0 {int} models.Notice
// @Failure 403 body is empty
func (c *NoticeController) PostSystem() {
	title := c.GetString("title")
	content := c.GetString("content")
	notice_type := c.GetString("notice_type")
	choice_type, _ := c.GetInt("type")
	data := make(map[string]interface{})
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	user_id, _ := User.GetUsers(choice_type)
	data["user_id"] = user_id
	data["title"] = title
	data["content"] = content
	data["notice_type"] = notice_type
	data["choice_type"] = choice_type
	result, _ := json.Marshal(data)
	r := utils.Redix.Get()
	defer r.Close()
	_, err := r.Do("LPUSH", "onemore:postSystem", result)
	if err == nil {
		c.Data["json"] = models.JSONStruct{"success", 0, nil, "保存成功！"}
	} else {
		c.Data["json"] = models.JSONStruct{"error", 1003, nil, err.Error()}
	}
	c.ServeJSON()
}

// DeleteOms oms删除通知
// @Title oms删除通知
// @Description oms删除通知
// @Param	id		path 	string	true		"通知ID"
// @Success 0 {string} delete success!
// @Failure 403 id is empty
// @router /notice_oms/:id [delete]
func (c *NoticeController) DeleteOms() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteOms(id)
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1004, nil, "删除失败！"}
	} else {
		c.Data["json"] = models.JSONStruct{"success", 0, nil, "删除成功！"}
	}
	c.ServeJSON()
}
