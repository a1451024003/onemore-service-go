package models

import (
	"fmt"
	"strconv"
	"onemore-service-go/utils"
	"github.com/astaxie/beego/orm"
)

type Conf struct {
	Id      int    `json:"id" orm:"column(id);auto" description:"编号"`
	Type    int    `json:"type" orm:"column(type);" description:"环境 1.local 2.dev 3.test 4.prod"`
	Service int    `json:"service" orm:"column(service);" description:"数据库类型"`
	Config  string `json:"config";orm:"column(config);size(45);null" description:"配置"`
}

func (t *Conf) TableName() string {
	return "system"
}

func init() {
	orm.RegisterModel(new(Conf))
}

/*
配置列表
*/
func Configs(ty int, service int) (ml interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	types := strconv.Itoa(ty)
	services := strconv.Itoa(service)
	key := "" + types + "_" + services
	r := utils.Redix.Get()
	defer r.Close()
	is_key_exit, err := r.Do("EXISTS", key)
	if is_key_exit != 0 {
		conf, err := r.Do("GET", key)
		return conf, err
	} else {
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("*").From("config").Where("type = ? and service = ? ").String()
		_, err = o.Raw(sql, ty, service).Values(&v)
		return v, err
	}
}

/*
添加配置
*/
func SystemPost(ty int, service int, config string) (err error) {
	o := orm.NewOrm()
	types := strconv.Itoa(ty)
	services := strconv.Itoa(service)
	key := "" + types + "_" + services
	r := utils.Redix.Get()
	defer r.Close()
	sql := "insert into config set type = ?,config = ?,service = ? "
	_, err = o.Raw(sql, ty, config, service).Exec()
	if err == nil {
		_, err = r.Do("SET", key, config)
		fmt.Println(err, 11)
		defer r.Close()
	}
	return err
}
