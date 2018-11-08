package models

import (
	"encoding/json"
	"math"
	"strings"

	"github.com/astaxie/beego/orm"
)

type HotChcp struct {
	Id                 int    `json:"id"`
	Name               string `json:"name" orm:"size(128)"`
	IosIdentifier      string `json:"ios_identifier" orm:"size(128)"`
	AndroidIdentifier  string `json:"android_identifier" orm:"size(128)"`
	Update             string `json:"update" orm:"size(128)"`
	Manifest           string `json:"manifest" orm:"size(128)"`
	ContentUrl         string `json:"content_url" orm:"size(128)"`
	Release            string `json:"release" orm:"size(128)"`
	MinNativeInterface int    `json:"min_native_interface"`
}

type Page struct {
	PageNum int          `json:"page_num"`
	PerPage int          `json:"per_page"`
	Total   int64        `json:"total"`
	Data    []orm.Params `json:"data"`
}

func (t *HotChcp) TableName() string {
	return "hot_chcp"
}

func init() {
	orm.RegisterModel(new(HotChcp))
}

//新增
func (w *HotChcp) AddHot() error {
	o := orm.NewOrm()
	if _, err := o.Insert(w); err != nil {
		return err
	}

	return nil
}

// GetAllHot_chap retrieves all Hot_chap matches certain condition. Returns empty list if
// no records exist
func (f *HotChcp) GetAll(prepage, page int) (Page, error) {
	var con []interface{}
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("count(*)").From(f.TableName()).String()

	println(sql)

	var total int64
	err := o.Raw(sql, con).QueryRow(&total)

	if err == nil {
		var smslog []orm.Params
		limit := 20
		if page <= 0 {
			page = 1
		}
		start := (page - 1) * limit

		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("*").From(f.TableName()).OrderBy("id").Desc().Limit(limit).Offset(start).String()

		if _, err := o.Raw(sql, con).Values(&smslog); err == nil {
			pageNum := int(math.Ceil(float64(total) / float64(limit)))

			return Page{pageNum, limit, total, smslog}, nil
		}
	}

	return Page{}, nil
}

//删除记录
func (f *HotChcp) Delete() (int64, error) {
	o := orm.NewOrm()
	if err := o.Read(f); err == nil {
		return o.Delete(f)
	} else {
		return 0, err
	}
}

//获取最新
func (f *HotChcp) GetNew() (ml map[string]interface{}, err error) {
	var chcp HotChcp
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	sql := qb.Select("*").From(f.TableName()).OrderBy("id").Desc().Limit(1).Offset(0).String()

	if err := o.Raw(sql).QueryRow(&chcp); err == nil {
		chcpjson, _ := json.Marshal(chcp)
		json.Unmarshal(chcpjson, &ml)
		var manifest []interface{}
		jsonstr := strings.Replace(ml["manifest"].(string), "\\", "", 0)
		json.Unmarshal([]byte(jsonstr), &manifest)
		ml["manifest"] = manifest
		delete(ml, "manifest")
		return ml, err
	}
	return ml, err
}

//编辑评分占比
func (a HotChcp) Save() error {
	tmp := a
	o := orm.NewOrm()
	if err := o.Read(&tmp, "id"); err == nil {
		a.Id = tmp.Id
		if _, err := o.Update(&a); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

//获取详情
func (f *HotChcp) GetOne(id int) (chcp HotChcp, err error) {
	o := orm.NewOrm()
	var hotChcp HotChcp
	err = o.QueryTable("hot_chcp").Filter("id", id).One(&hotChcp)

	return hotChcp, err
}

//获取最新版本信息
func (f *HotChcp) GetNewInfo() (ml map[string]interface{}, err error) {
	var chcp HotChcp
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	sql := qb.Select("*").From(f.TableName()).OrderBy("id").Desc().Limit(1).Offset(0).String()

	if err := o.Raw(sql).QueryRow(&chcp); err == nil {
		chcpjson, _ := json.Marshal(chcp)
		json.Unmarshal(chcpjson, &ml)
		var manifest []interface{}
		jsonstr := strings.Replace(ml["manifest"].(string), "\\", "", 0)
		json.Unmarshal([]byte(jsonstr), &manifest)
		ml["manifest"] = manifest
		delete(ml, "id")
		delete(ml, "android_identifier")
		delete(ml, "content_url")
		delete(ml, "ios_identifier")
		delete(ml, "min_native_interface")
		delete(ml, "name")
		delete(ml, "release")
		delete(ml, "update")
		return ml, err
	}
	return ml, err
}
