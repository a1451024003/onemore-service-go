package servers

import (
	"onemore-service-go/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hprose/hprose-golang/rpc"
)

type NoticeServer struct {
}

func (c *NoticeServer) Init() {
	server := rpc.NewHTTPService()
	server.AddAllMethods(&NoticeServer{})
	beego.Handler("/rpc/notice", server)
}

type Notice struct {
}

//发布动态
func (c *NoticeServer) AddDynamic(tetle string, name string, url string, choice_type int, notice_type string, user_id string, teacher_id int, group_type int, send_id int) error {
	_, err := models.AddDynamic(tetle, name, url, choice_type, notice_type, user_id, teacher_id, group_type, send_id)
	return err
}

//发布任务
func (c *NoticeServer) AddWork(title map[string]interface{}, name string, url_id int, work_notice map[string]interface{}, notice_type string) error {
	_, err := models.AddWork(title, name, url_id, work_notice, notice_type)
	return err
}

//编辑阅读状态
func (c *NoticeServer) SaveRead(user_id int, send_id int, group_type int) error {
	o := orm.NewOrm()
	var notice []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("nr.id").From("notice as n").LeftJoin("notice_relation as nr ").
		On("n.id = nr.notice_id").Where("nr.user_id = ? and nr.send_id = ? and n.group_type = ? ").String()
	_, err := o.Raw(sql, user_id, send_id, group_type).Values(&notice)
	for _, v := range notice {
		_, err = o.QueryTable("notice_relation").Filter("id", v["id"]).Update(orm.Params{
			"read_type": 2,
		})
	}
	return err
}

//系统通知
func (c *NoticeServer) PostSystem(value string) (err error) {
	_, err = models.AddSystem(value)
	return err
}

//邀请通知
func (c *NoticeServer) InviteSystem(value string) (err error) {
	_, err = models.InviteSystem(value)
	return err
}

//认证幼儿园通知
func (c *NoticeServer) AttestKindergarten(value string) (err error) {
	_, err = models.AttestKindergarten(value)
	return err
}

//编辑阅读状态
func (c *NoticeServer) NoticeRead(notice_id int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("notice_relation").Filter("notice_id", notice_id).Update(orm.Params{
		"read_type": 2,
	})
	return err
}

// 取消邀请 改变通知状态 type 1 教师   type 2 学生
func (c *NoticeServer) ResetStatus(id int, director string, ty int, user_id int, kinder_name string) (err error) {
	o := orm.NewOrm()
	var notice []orm.Params
	if ty == 1 {
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("nr.id").From("notice as n").LeftJoin("notice_relation as nr ").
			On("n.id = nr.notice_id").Where("nr.user_id = ? and n.types = ?").String()
		_, err = o.Raw(sql, id, 1).Values(&notice)
		if err != nil {
			return err
		}
		for _, v := range notice {
			_, err = o.QueryTable("notice").Filter("id", v["id"]).Update(orm.Params{
				"title": "邀请通知", "content": kinder_name + "取消了对您的邀请。", "notice_type": 6, "type": 3, "name": director, "invite_type": 0,
			})
			_, err = o.QueryTable("notice_relation").Filter("notice_id", v["id"]).Update(orm.Params{
				"read_type": 1,
			})
		}
	} else {
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("nr.id").From("notice as n").LeftJoin("notice_relation as nr ").
			On("n.id = nr.notice_id").Where("nr.baby_id = ? and n.types = ?").String()
		_, err = o.Raw(sql, id, 2).Values(&notice)
		if err != nil {
			return err
		}
		for _, v := range notice {
			_, err = o.QueryTable("notice").Filter("id", v["id"]).Update(orm.Params{
				"title": "邀请通知", "content": kinder_name + "取消了对您的邀请。", "notice_type": 6, "type": 3, "name": director, "invite_type": 0,
			})
			_, err = o.QueryTable("notice_relation").Filter("notice_id", v["id"]).Update(orm.Params{
				"read_type": 1,
			})
		}
	}
	return err
}

// 拒绝/接受邀请 改变通知状态 type 1 教师   type 2 学生
func (c *NoticeServer) UpdateNoticeStatus(notice_id int, ty int, director string, invite int, kinderName string) (err error) {
	o := orm.NewOrm()
	if ty == 1 {
		if invite == 1 {
			_, err = o.QueryTable("notice").Filter("id", notice_id).Update(orm.Params{
				"title": "邀请通知", "content": "您的账号已接受邀请，成为" + kinderName + "老师。", "notice_type": 6, "type": 3, "name": director, "invite_type": 0,
			})
			_, err = o.QueryTable("notice_relation").Filter("notice_id", notice_id).Update(orm.Params{
				"read_type": 1,
			})
		} else {
			_, err = o.QueryTable("notice").Filter("id", notice_id).Update(orm.Params{
				"title": "邀请通知", "content": "您拒绝了" + kinderName + "的邀请。", "notice_type": 6, "type": 3, "name": director, "invite_type": 0,
			})
			_, err = o.QueryTable("notice_relation").Filter("notice_id", notice_id).Update(orm.Params{
				"read_type": 1,
			})
		}
	} else {
		if invite == 1 {
			_, err = o.QueryTable("notice").Filter("id", notice_id).Update(orm.Params{
				"title": "邀请通知", "content": "您的账号已接受邀请，成为" + kinderName + "学生", "notice_type": 6, "type": 3, "name": director, "invite_type": 0,
			})
			_, err = o.QueryTable("notice_relation").Filter("notice_id", notice_id).Update(orm.Params{
				"read_type": 1,
			})
		} else {
			_, err = o.QueryTable("notice").Filter("id", notice_id).Update(orm.Params{
				"title": "邀请通知", "content": "您拒绝了" + kinderName + "的邀请。", "notice_type": 6, "type": 3, "name": director, "invite_type": 0,
			})
			_, err = o.QueryTable("notice_relation").Filter("notice_id", notice_id).Update(orm.Params{
				"read_type": 1,
			})
		}
	}
	return err
}
