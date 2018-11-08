package controllers

//codes
import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CodesController struct {
	beego.Controller
}

type JSONStruct1 struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}
type JSONStruct2 struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	Page_num int         `json:"page_num"`
}
type JSONStruct3 struct {
	Data interface{} `json:"data"`
}
type JSONStruct4 struct {
	Id bool `json:"id"`
}
type JSONStruct5 struct {
	Id int `json:"id"`
}

type Codes struct {
	Name       string `json:"name"`
	Id         int    `json:"id"`
	Version    string `json:"version"`
	Service_id int    `json:"service_id"`
	Code       int    `json:"code"`
	Content    string `json:"content"`
	Status     int    `json:"status"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Code struct {
	Id         int    `json:"id"`
	Version    string `json:"version"`
	Service_id int    `json:"service_id"`
	Code       int    `json:"code"`
	Content    string `json:"content"`
	Status     int    `json:"status"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

// URLMapping ...
func (c *CodesController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @router / [get]
func (c *CodesController) GetAll() {
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("s.name", "c.id", "c.code", "s.version", "c.service_id", "c.content", "c.status", "c.created_at", "c.updated_at").
		From("code as c").
		InnerJoin("service as s").
		On("c.service_id = s.id").String()
	o := orm.NewOrm()
	var codes []Codes
	num, err := o.Raw(sql).QueryRows(&codes)
	var ml []interface{}
	if err == nil && num > 0 {
		for _, v := range codes {
			ml = append(ml, v)
		}
		c.Data["json"] = JSONStruct1{Status: "success", Code: 0, Result: JSONStruct2{Data: ml, Total: num, Page_num: 0}, Msg: "获取列表成功！"}
	} else {
		c.Data["json"] = JSONStruct1{Status: "error", Code: 1, Result: "", Msg: "获取列表失败！"}
	}
	c.ServeJSON()
}

// @router /:id [get]
func (c *CodesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("code").Where("id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	var code Code
	err := o.Raw(sql, id).QueryRow(&code)
	if err == nil {
		c.Data["json"] = JSONStruct1{Status: "success", Code: 0, Result: code, Msg: "获取详情成功！"}
	} else {
		c.Data["json"] = JSONStruct1{Status: "error", Code: 1, Result: "", Msg: "获取详情失败！"}
	}
	c.ServeJSON()
}

// @router / [post]
func (c *CodesController) Post() {
	service_id, _ := c.GetInt("service_id")
	code, _ := c.GetInt("code")
	content := c.GetString("content")
	switch {
	case service_id == 0:
		c.Data["json"] = JSONStruct1{Status: "error", Code: 1001, Result: "", Msg: "补全参数！"}
	case code == 0:
		c.Data["json"] = JSONStruct1{Status: "error", Code: 1001, Result: "", Msg: "补全参数！"}
	case content == "":
		c.Data["json"] = JSONStruct1{Status: "error", Code: 1001, Result: "", Msg: "补全参数！"}
	default:
		o := orm.NewOrm()
		var codes []Code
		num, err := o.Raw("select * from code where code = ?", code).QueryRows(&codes)
		if err == nil && num > 0 {
			c.Data["json"] = JSONStruct1{Status: "error", Code: 501, Result: "", Msg: "添加失败！"}
		} else {
			status := 0
			created_at := time.Now()
			res, err := o.Raw("insert into code (service_id,code,content,status,created_at) values (?,?,?,?,?)", service_id, code, content, status, created_at).Exec()
			id, err := res.LastInsertId()
			if err == nil && id > 0 {
				c.Data["json"] = JSONStruct1{Status: "success", Code: 0, Result: JSONStruct4{Id: true}, Msg: "添加成功！"}
			}
		}
	}
	c.ServeJSON()
}

// @router /:id [put]
func (c *CodesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println(id)
	switch {
	case id == 0:
	default:
		code, _ := c.GetInt("code")
		content := c.GetString("content")
		datetime := time.Now()
		o := orm.NewOrm()
		res, err := o.Raw("update code set code = ?, content = ?, updated_at = ? where id = ?", code, content, datetime, id).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num > 0 {
				c.Data["json"] = JSONStruct1{Status: "success", Code: 0, Result: JSONStruct5{Id: 1}, Msg: "修改成功！"}
			} else {
				c.Data["json"] = JSONStruct1{Status: "error", Code: 1, Result: "", Msg: "修改失败！"}
			}
		} else {
			c.Data["json"] = JSONStruct1{Status: "error", Code: 1, Result: "", Msg: "修改失败！"}
		}
	}
	c.ServeJSON()
}

// @router /:id [delete]
func (c *CodesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	datetime := time.Now()
	o := orm.NewOrm()
	res, err := o.Raw("update code set status = 1, updated_at = ? where id = ?", datetime, id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			c.Data["json"] = JSONStruct1{Status: "success", Code: 0, Result: "", Msg: "删除成功！"}
		} else {
			c.Data["json"] = JSONStruct1{Status: "error", Code: 1, Result: "", Msg: "删除失败！"}
		}
	}
	c.ServeJSON()
}
