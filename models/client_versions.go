package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ClientVersions struct {
	Id           int       `orm:"column(id);auto"`
	Type         string    `orm:"column(type);size(191)" description:"版本类型"`
	Version      string    `orm:"column(version);size(191)" description:"版本名称"`
	VersionCode  string    `orm:"column(version_code);size(11)" description:"版本号"`
	InternalCode string    `orm:"column(internal_code);size(11);null"`
	Description  string    `orm:"column(description)" description:"更新说明"`
	Link         string    `orm:"column(link);size(191)" description:"链接地址"`
	IsForced     int8      `orm:"column(is_forced)" description:"是否强制更新"`
	CreatedAt    time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(timestamp);null"`
}

func (t *ClientVersions) TableName() string {
	return "client_versions"
}

func init() {
	orm.RegisterModel(new(ClientVersions))
}

// AddClientVersions insert a new ClientVersions into database and returns
// last inserted Id on success.
func AddClientVersions(m *ClientVersions) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetClientVersionsById retrieves ClientVersions by Id. Returns error if
// Id doesn't exist
func GetClientVersionsById(id int) (v *ClientVersions, err error) {
	o := orm.NewOrm()
	v = &ClientVersions{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllClientVersions retrieves all ClientVersions matches certain condition. Returns empty list if
// no records exist
func GetAllClientVersions(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ClientVersions))
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

	var l []ClientVersions
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

// UpdateClientVersions updates ClientVersions by Id and returns error if
// the record to be updated doesn't exist
func UpdateClientVersionsById(m *ClientVersions) (err error) {
	o := orm.NewOrm()
	v := ClientVersions{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteClientVersions deletes ClientVersions by Id and returns error if
// the record to be deleted doesn't exist
func DeleteClientVersions(id int) (err error) {
	o := orm.NewOrm()
	v := ClientVersions{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ClientVersions{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
