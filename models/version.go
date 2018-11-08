package models

import (
	"time"

	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"onemore-service-go/utils"
)

type Version struct {
	Id          int       `json:"id"`
	HtmlVersion string    `json:"html_version"`
	HtmlUrl     string    `json:"html_url"`
	Type        int8      `json:"type"`
	AppType     int8      `json:"app_type"`
	CreatedAt   time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt   time.Time `json:"updated_at" orm:"auto_now"`
}

func (v *Version) TableName() string {
	return "version"
}

func init() {
	orm.RegisterModel(new(Version))
}

func (v *Version) GetNew() (map[string]string, error) {
	res := map[string]string{"html_version": "", "html_url": "", "version": "", "store_url": ""}

	o := orm.NewOrm()
	r := utils.Redix.Get()
	defer r.Close()
	replyH, _ := redis.Values(r.Do("hgetall", fmt.Sprintf("version:%d:1", v.AppType)))
	if len(replyH) > 0 {
		if string(replyH[0].([]byte)) == "url" {
			res["html_version"] = string(replyH[3].([]byte))
			res["html_url"] = string(replyH[1].([]byte))
		} else {
			res["html_version"] = string(replyH[1].([]byte))
			res["html_url"] = string(replyH[3].([]byte))
		}
	} else {
		var h5 Version
		if err := o.QueryTable(v).Filter("type", 1).Filter("app_type", v.AppType).OrderBy("-id", "-html_version").One(&h5); err == nil {
			res["html_version"] = h5.HtmlVersion
			res["html_url"] = h5.HtmlUrl

			r.Do(
				"hmset",
				fmt.Sprintf("version:%d:1", v.AppType),
				"version",
				h5.HtmlVersion,
				"url",
				h5.HtmlUrl,
			)
			} else if err != orm.ErrNoRows {
				return res, err
			}
		}
	replyO, _ := redis.Values(r.Do("hgetall", fmt.Sprintf("version:%d:%d", v.AppType, v.Type)))
	if len(replyO) > 0 {
		if string(replyO[0].([]byte)) == "url" {
			res["version"] = string(replyO[3].([]byte))
			res["store_url"] = string(replyO[1].([]byte))
		} else {
			res["version"] = string(replyO[1].([]byte))
			res["store_url"] = string(replyO[3].([]byte))
		}
	} else {
		var other Version
		if err := o.QueryTable(v).Filter("type", v.Type).Filter("app_type", v.AppType).OrderBy("-id", "-html_version").One(&other); err == nil {
			res["version"] = other.HtmlVersion
			res["store_url"] = other.HtmlUrl
			r.Do(
				"hmset",
				fmt.Sprintf("version:%d:%d", v.AppType, v.Type),
				"version",
				other.HtmlVersion,
				"url",
				other.HtmlUrl,
			)
		} else if err != orm.ErrNoRows {
			return res, err
		}
	}

	return res, nil
}

func (v *Version) Save() error {
	r := utils.Redix.Get()
	defer r.Close()
	o := orm.NewOrm()
	o.Begin()

	if _, err := orm.NewOrm().Insert(v); err != nil {
		o.Rollback()

		return err
	}

	if _, err := r.Do(
		"hmset",
		fmt.Sprintf("version:%d:%d", v.AppType, v.Type),
		"version",
		v.HtmlVersion,
		"url",
		v.HtmlUrl,
	); err != nil {
		o.Rollback()

		return err
	}
	o.Commit()

	return nil
}
