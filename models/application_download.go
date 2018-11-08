package models

import (
	"math"
	"time"

	"github.com/astaxie/beego/orm"
)

type ApplicationDownload struct {
	Id          int    `json:"id"; orm:"column(id);auto" description:"自增ID"`
	Icon        string `json:"icon"; orm:"column(icon);size(200);null" description:"应用图标"`
	Name        string `json:"name"; orm:"column(name);size(45);null" description:"应用名称"`
	DownloadUrl string `json:"download_url"; orm:"column(download_url);size(200);null" description:"原唱音频"`
	Type        int    `json:"type"; orm:"column(type);null" description:"类型：1儿歌，2游戏"`
	PackageName string `json:"package_name"; orm:"column(package_name);null" description:"应用包名"`
	VersionCode string `json:"version_code"; orm:"column(version_code);null" description:"应用版本号"`
	Size        string `json:"size"; orm:"column(size);null" description:"应用apk包大小（MB）"`
	CreatedAt   string `json:"created_at"; orm:"auto_now_add" description:"创建时间"`
	UpdatedAt   string `json:"updated_at"; orm:"auto_now" description:"修改时间"`
}

func init() {
	orm.RegisterModel(new(ApplicationDownload))
}

// 应用下载-添加
// last inserted Id on success.
func AddApplicationDownload(m *ApplicationDownload) map[string]interface{} {
	o := orm.NewOrm()
	m.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	id, err := o.Insert(m)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = id //返回数据
		return paginatorMap
	}
	return nil
}

// GetApplicationDownloadById retrieves ApplicationDownload by Id. Returns error if
// Id doesn't exist
func GetApplicationDownloadById(id int) (v *ApplicationDownload, err error) {
	o := orm.NewOrm()
	v = &ApplicationDownload{Id: id}
	if err = o.QueryTable(new(ApplicationDownload)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 应用下载-列表
// no records exist
func GetAllApplicationDownload(ty int, page, prepage int) map[string]interface{} {
	var v []ApplicationDownload
	o := orm.NewOrm()
	nums, err := o.QueryTable("application_download").Filter("Type", ty).OrderBy("id").All(&v)
	if err == nil && nums > 0 {
		//根据nums总数，和prepage每页数量 生成分页总数
		totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
		if page > totalpages {
			page = totalpages
		}
		if page <= 0 {
			page = 1
		}
		limit := (page - 1) * prepage
		num, err := o.QueryTable("application_download").Filter("Type", ty).Limit(prepage, limit).All(&v)
		if err == nil && num > 0 {
			paginatorMap := make(map[string]interface{})
			paginatorMap["total"] = nums          //总条数
			paginatorMap["data"] = v              //分页数据
			paginatorMap["page_num"] = totalpages //总页数
			return paginatorMap
		}
	}
	return nil
}

// 应用下载-修改
// the record to be updated doesn't exist
func UpdateApplicationDownloadById(m *ApplicationDownload) map[string]interface{} {
	o := orm.NewOrm()
	m.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	v := ApplicationDownload{Id: m.Id}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			paginatorMap := make(map[string]interface{})
			paginatorMap["data"] = num //返回数据
			return paginatorMap
		}
	}
	return nil
}

// 应用下载-删除
// the record to be deleted doesn't exist
func DeleteApplicationDownload(id int) map[string]interface{} {
	o := orm.NewOrm()
	v := ApplicationDownload{Id: id}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ApplicationDownload{Id: id}); err == nil {
			paginatorMap := make(map[string]interface{})
			paginatorMap["data"] = num //返回数据
			return paginatorMap
		}
	}
	return nil
}
