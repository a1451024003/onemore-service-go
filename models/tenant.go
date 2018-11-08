package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Tenant struct {
	Id        int       `json:"tenant_id" orm:"column(tenant_id);auto" description:"编号"`
	Name      string    `json:"name";orm:"column(name);size(45);null" description:"名称"`
	IsFree    int       `json:"is_free";orm:"column(is_free)" description:"是否免费：0免费1付费"`
	FeeType   int       `json:"fee_type";orm:"column(fee_type)" description:"付费类型：1包年2包季度3包月"`
	FeePeriod string    `json:"fee_period";orm:"column(fee_period);size(20)" description:"租期时长"`
	StartTime time.Time `json:"start_time";orm:"column(start_time);type(datetime)" description:"开始时间"`
	EndTime   time.Time `json:"end_time";orm:"column(end_time);type(datetime)" description:"结束时间"`
	RentType  int       `json:"rent_type";orm:"column(rent_type)" description:"租借状态0正常1到期未付费2到期停"`
	CreatedAt time.Time `json:"created_at";orm:"column(created_at);type(timestamp);null"`
	UpdatedAt time.Time `json:"updated_at";orm:"column(updated_at);type(timestamp);null"`
	Status    int8      `json:"status";orm:"column(status);null" description:"状态：0:正常，1:删除"`
}

func (t *Tenant) TableName() string {
	return "tenant"
}

func init() {
	orm.RegisterModel(new(Tenant))
}

/*
租户详情
*/
func GetTenant(id int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("tenant").Where("tenant_id = ?").String()
	_, err = o.Raw(sql, id).Values(&v)
	if err == nil {
		ml := make(map[string]interface{})
		ml["data"] = v
		return ml, nil
	}
	return nil, err
}

/*
租户列表
*/
func TenantList() (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("tenant").Where("status = 0").String()
	_, err = o.Raw(sql).Values(&v)
	if err == nil {
		ml := make(map[string]interface{})
		ml["data"] = v
		return ml, nil
	}
	return nil, err
}

/*
添加租户
*/
func Post(name string, is_free int, fee_period string, fee_type int, start_time string, end_time string, rent_type int) (err error) {
	o := orm.NewOrm()
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	loc, _ := time.LoadLocation("")
	timenow := time.Now().Format("2006-01-02 15:04:05")
	createTime, _ := time.ParseInLocation(timeLayout, timenow, loc)
	sql := "insert into tenant set name = ?,is_free = ?,fee_period = ?,fee_type = ?,start_time = ?,end_time = ?,rent_type = ?,created_at = ? "
	_, err = o.Raw(sql, name, is_free, fee_period, fee_type, start_time, end_time, rent_type, createTime).Exec()
	return err
}

/*
删除租户
*/
func Delete(id int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("tenant").Filter("tenant_id", id).Update(orm.Params{
		"status": 1,
	})
	fmt.Println(err)
	return err
}

/*
编辑租户
*/
func Update(id int, name string, is_free int, fee_period string, fee_type int, start_time string, end_time string, rent_type int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("tenant").Filter("tenant_id", id).Update(orm.Params{
		"name": name, "is_free": is_free, "fee_period": fee_period, "fee_type": fee_type, "start_time": start_time, "end_time": end_time, "rent_type": rent_type,
	})
	return err
}
