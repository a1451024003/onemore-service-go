package models

import (
	"errors"
	"fmt"

	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TenantKindergarten struct {
	Id             int `orm:"column(id);auto" description:"租户幼儿园  编号"`
	TenantId       int `orm:"column(tenant_id);null" description:"租户ID"`
	KindergartenId int `orm:"column(kindergarten_id);null" description:"幼儿园ID"`
}

func (t *TenantKindergarten) TableName() string {
	return "tenant_kindergarten"
}

func init() {
	orm.RegisterModel(new(TenantKindergarten))
}

// 添加租户幼儿园
// last inserted Id on success.
func (w *TenantKindergarten) AddTenantKindergarten() error {
	o := orm.NewOrm()
	if _, err := o.Insert(w); err != nil {
		return err
	}

	return nil
}

// GetTenantKindergartenById retrieves TenantKindergarten by Id. Returns error if
// Id doesn't exist
func GetTenantKindergartenById(id int) (v *TenantKindergarten, err error) {
	o := orm.NewOrm()
	v = &TenantKindergarten{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTenantKindergarten retrieves all TenantKindergarten matches certain condition. Returns empty list if
// no records exist
func GetAllTenantKindergarten(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TenantKindergarten))
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

	var l []TenantKindergarten
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

// UpdateTenantKindergarten updates TenantKindergarten by Id and returns error if
// the record to be updated doesn't exist
func UpdateTenantKindergartenById(m *TenantKindergarten) (err error) {
	o := orm.NewOrm()
	v := TenantKindergarten{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTenantKindergarten deletes TenantKindergarten by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTenantKindergarten(id int) (err error) {
	o := orm.NewOrm()
	v := TenantKindergarten{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TenantKindergarten{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
