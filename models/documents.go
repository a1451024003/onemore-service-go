package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Documents struct {
	Id        int       `orm:"column(id);auto" description:"id"`
	Name      string    `orm:"column(name);size(200);null" description:"模块名称"`
	ParentId  int       `orm:"column(parent_id);null" description:"父级ID"`
	Type      int       `orm:"column(type);null" description:"1,产品文档2，开发文档"`
	Details   string    `orm:"column(details);null" description:"详情"`
	Edition   string    `orm:"column(edition);size(11);null" description:"版本"`
	Status    int       `orm:"column(status);null" description:"0正常，1删除"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null" description:"创建时间"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);null" description:"修改时间"`
}

func (t *Documents) TableName() string {
	return "documents"
}

func init() {
	orm.RegisterModel(new(Documents))
}

// AddDocuments insert a new Documents into database and returns
// last inserted Id on success.
func AddDocuments(m *Documents) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDocumentsById retrieves Documents by Id. Returns error if
// Id doesn't exist
func GetDocumentsById(id int) (v *Documents, err error) {
	o := orm.NewOrm()
	v = &Documents{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDocuments retrieves all Documents matches certain condition. Returns empty list if
// no records exist
func GetAllDocuments(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Documents))
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

	var l []Documents
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

// UpdateDocuments updates Documents by Id and returns error if
// the record to be updated doesn't exist
func UpdateDocumentsById(m *Documents) (err error) {
	o := orm.NewOrm()
	v := Documents{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDocuments deletes Documents by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDocuments(id int) (err error) {
	o := orm.NewOrm()
	v := Documents{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Documents{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
