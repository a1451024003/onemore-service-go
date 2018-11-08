package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type MaterialImg struct {
	Id         int       `orm:"column(id);auto"`
	MaterialId int       `orm:"column(material_id);null"`
	Url        string    `orm:"column(url);size(200);null" description:"路径"`
	Name       string    `orm:"column(name);size(50);null" description:"图片名称"`
	Status     int       `orm:"column(status);null" description:"-1,审核未通过，0待审核，1，审核成功"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);null"`
}

func (t *MaterialImg) TableName() string {
	return "material_img"
}

func init() {
	orm.RegisterModel(new(MaterialImg))
}

// AddMaterialImg insert a new MaterialImg into database and returns
// last inserted Id on success.
func AddMaterialImg(m *MaterialImg) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMaterialImgById retrieves MaterialImg by Id. Returns error if
// Id doesn't exist
func GetMaterialImgById(id int) (v *MaterialImg, err error) {
	o := orm.NewOrm()
	v = &MaterialImg{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMaterialImg retrieves all MaterialImg matches certain condition. Returns empty list if
// no records exist
func GetAllMaterialImg(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MaterialImg))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MaterialImg
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateMaterialImg updates MaterialImg by Id and returns error if
// the record to be updated doesn't exist
func UpdateMaterialImgById(m *MaterialImg) (err error) {
	o := orm.NewOrm()
	v := MaterialImg{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMaterialImg deletes MaterialImg by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMaterialImg(id int) (err error) {
	o := orm.NewOrm()
	v := MaterialImg{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MaterialImg{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
