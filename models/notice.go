package models

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/hprose/hprose-golang/rpc"
)

type Request struct {
	Id             int    `json:"id";orm:"column(id);auto"`
	Title          string `json:"title";orm:"column(title);size(100)"`
	Content        string `json:"content";orm:"column(content)"`
	ClassGroup     string `json:"class_group";orm:"column(class_group)"`
	Teacher        string `json:"teacher";orm:"column(teacher)"`
	ParentGroup    string `json:"parent_group";orm:"column(parent_group)"`
	UserId         string `json:"user_id";orm:"column(user_id)"`
	NoticeType     string `json:"notice_type";orm:"column(notice_type)"`
	Name           string `json:"name";orm:"column(name);size(30)"`
	KindergartenId string `json:"kindergarten_id";orm:"column(kindergarten_id);size(30)"`
	Url            string `json:"url";orm:"column(url);size(100)"`
	MeetTime       string `json:"meet_time";orm:"column(meet_time);type(datetime);auto_now_add"`
	MeetAddress    string `json:"meet_address";orm:"column(meet_address);size(30)"`
	CreatedAt      string `json:"created_at";orm:"column(created_at);size(30)"`
}

type Notice struct {
	Id             int    `json:"id";orm:"column(id);auto"`
	Title          string `json:"title";orm:"column(title);size(100)"`
	Content        string `json:"content";orm:"column(content)"`
	NoticeType     int    `json:"notice_type";orm:"column(notice_type)"`
	Name           string `json:"name";orm:"column(name);size(30)"`
	MeetTime       string `json:"meet_time";orm:"column(meet_time);type(datetime);auto_now_add"`
	MeetAddress    string `json:"meet_address";orm:"column(meet_address);size(30)"`
	Url            string `json:"url";orm:"column(url);size(100)"`
	UrlId          int    `json:"url_id"`
	Type           int    `json:"type";orm:"column(type);"`
	kindergertenId int    `json:"kindergerten_id";orm:"column(kindergerten_id);"`
	InviteType     int    `json:"invite_type";orm:"column(invite_type);"`
	Types          int    `json:"types";orm:"column(types);"`
	GroupType      int    `json:"group_type";orm:"column(group_type);"`
	CreatedAt      string `json:"created_at";orm:"column(created_at);type(datetime);auto_now_add"`
}
type JSONStruct struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}
type NoticeRelation struct {
	Id       int `json:"auto";orm:"column(id);auto"`
	NoticeId int `json:"notice_id";orm:"column(notice_id)"`
	UserId   int `json:"user_id";orm:"column(user_id)"`
	SendId   int `json:"send_id";orm:"column(send_id)"`
	ReadType int `json:"read_type";orm:"column(read_type)"`
}

var User *UserService
var Mqtt *MqttService
var Kg *KgService

type UserService struct {
	GetOneByUserId func(userId int) (interface{}, error)
	GetBabyInfo    func(id int) (map[string]interface{}, error)
	GetUsers       func(types int) ([]map[string]interface{}, error)
}

type MqttService struct {
	Pub func(int, int, string) error
}

type KgService struct {
	UserPost      func(user_id int) (string, error)
	TeacherNotice func(class_type int, kindergarten_id int) (interface{}, error)
	StudentNotice func(class_type int, kindergarten_id int) (interface{}, error)
}

func (t *Notice) TableName() string {
	return "notice"
}

func init() {
	orm.RegisterModel(new(Notice))
	orm.RegisterModel(new(NoticeRelation))
}

// AddNotice 发布通知
func AddNotice(m *Request) (id int64, err error) {
	o := orm.NewOrm()
	var v Notice
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	user_id := m.UserId
	i, _ := strconv.Atoi(user_id)
	user, err := User.GetOneByUserId(i)
	if err != nil {
		return 0, err
	}
	users := user.(map[string]interface{})
	v.Title = m.Title
	v.Content = m.Content
	t, err := strconv.Atoi(m.NoticeType)
	v.NoticeType = t
	user_name := users["name"]
	v.Name = user_name.(string)
	v.MeetAddress = m.MeetAddress
	v.MeetTime = m.MeetTime
	v.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	o.Begin()
	//插入通知表
	id, err = o.Insert(&v)

	if err == nil && id != 0 {
		//获取教师组的user_id
		class_group := m.ClassGroup
		class_client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_KG_SERVER"))
		class_client.UseService(&Kg)
		class_str := ""
		kindergarten_id := m.KindergartenId
		k_id, _ := strconv.Atoi(kindergarten_id)
		if len(class_group) > 0 {
			class_group_num := strings.Split(class_group, ",")
			for _, c := range class_group_num {
				cg, _ := strconv.Atoi(c)
				class, _ := Kg.TeacherNotice(cg, k_id)
				if class != nil {
					class_ids := class.(map[interface{}]interface{})
					data := class_ids["data"].([]interface{})
					for _, v := range data {
						user_id := v.(map[interface{}]interface{})
						ids := user_id["user_id"]
						class_str += ids.(string) + ","
					}
				}
			}
		}

		//获取家长组的user_id
		parent_group := m.ParentGroup
		parent_str := ""
		strc := make(map[int]int)
		if len(parent_group) > 0 {
			parent_group_num := strings.Split(parent_group, ",")
			for _, g := range parent_group_num {
				gg, _ := strconv.Atoi(g)
				parent, _ := Kg.StudentNotice(gg, k_id)
				if parent != nil {
					parent_ids := parent.(map[interface{}]interface{})
					data := parent_ids["data"].([]interface{})
					for _, gv := range data {
						parent_id := gv.(map[interface{}]interface{})
						gids := parent_id["baby_id"]
						uid := gids.(string)
						uuid, _ := strconv.Atoi(uid)
						home, _ := User.GetBabyInfo(uuid)
						if home != nil {
							strc[int(home["creator"].(int))] = 0
						}
					}
				}
			}
		}
		if user_id != "" {
			tea, _ := strconv.Atoi(user_id)
			strc[tea] = 0
		}
		for k, _ := range strc {
			home_str := strconv.Itoa(k)
			parent_str += home_str + ","
		}
		ids := strings.TrimLeft(m.Teacher+class_str+parent_str, ",")
		lens := len(ids)
		s := string([]rune(ids)[:(lens - 1)])
		user_ids := strings.Split(s, ",")
		for _, v := range user_ids {
			sql := "insert into notice_relation set user_id = ?,notice_id = ?,read_type = ?"
			_, err = o.Raw(sql, v, id, 1).Exec()
		}
		if err != nil {
			o.Rollback()
			return 0, err
		} else {
			var notice []orm.Params
			qb, _ := orm.NewQueryBuilder("mysql")
			sql := qb.Select("n.*", "nr.*").From("notice as n").LeftJoin("notice_relation as nr ").
				On("n.id = nr.notice_id").Where("n.id = ?").String()
			_, err = o.Raw(sql, id).Values(&notice)
			var mString string
			client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_MQTT_SERVER"))
			client.UseService(&Mqtt)
			for _, v := range notice {
				jsonData, _ := json.Marshal(v)
				mString = string(jsonData) //text
				for _, v := range user_ids {
					cg, _ := strconv.Atoi(v)
					err := Mqtt.Pub(3, cg, mString) // 1所有平台用户 2所有幼儿园用户 3用户
					// 设置文件名
					filename := time.Now().Format("20060102")
					logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"storage/logs/%s"}`, filename))
					logs.Info(err)
				}
				break
			}
			o.Commit()
			return 1, nil
		}
	} else {
		o.Rollback()
		return 0, err
	}
	return 0, nil
}

// GetNoticeById 通知详情
func GetNoticeById(id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	s := strconv.Itoa(id)
	qb, _ := orm.NewQueryBuilder("mysql")
	where := " n.id = " + s
	sql := qb.Select("*").From("notice as n").Where(where).OrderBy("n.id").Desc().String()
	if num, err := o.Raw(sql).Values(&maps); err == nil && num > 0 {
		for _, v := range maps {
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}

// GetAllNotice 通知列表
func GetAllNotice(user_id int, ty int, search string, notice_type int) (ml interface{}, err error) {
	o := orm.NewOrm()
	var notice []orm.Params
	var con []interface{}
	where := "1 "
	if ty != 0 {
		where += "AND n.type = ? "
		con = append(con, ty)
	}
	if notice_type != 0 {
		where += "AND n.notice_type = ? "
		con = append(con, notice_type)
	}
	if search != "" {
		where += "AND n.title like ? "
		con = append(con, "%"+search+"%")
	}
	if user_id > 0 {
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("n.*", "nr.*").From("notice as n").LeftJoin("notice_relation as nr ").
			On("n.id = nr.notice_id").Where("nr.user_id = ?").OrderBy("n.created_at").Desc().String()
		_, err = o.Raw(sql, user_id).Values(&notice)
	} else {
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("n.*").From("notice as n").Where(where).And("notice_type = ?").OrderBy("n.created_at").Desc().String()
		_, err = o.Raw(sql, con, 6).Values(&notice)
	}
	return notice, err
}

//系统通知列表
func GetNotice(user_id int, ty int, search string, notice_type int, page int, prepage int) (ml interface{}, err error) {
	o := orm.NewOrm()
	var notice []orm.Params
	var con []interface{}
	where := "1 "
	if ty != 0 {
		where += "AND n.type = ? "
		con = append(con, ty)
	}
	if notice_type != 0 {
		where += "AND n.notice_type = ? "
		con = append(con, notice_type)
	}
	if search != "" {
		where += "AND n.title like ? "
		con = append(con, "%"+search+"%")
	}
	// 构建查询对象
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("count(*)").From("notice as n").Where(where).And("notice_type = ?").OrderBy("n.created_at").Desc().String()
	var total int64
	err = o.Raw(sql, con, 6).QueryRow(&total)
	if err == nil {
		//根据nums总数，和prepage每页数量 生成分页总数
		totalpages := int(math.Ceil(float64(total) / float64(prepage))) //page总数
		if page > totalpages {
			page = totalpages
		}
		if page <= 0 {
			page = 1
		}
		limit := (page - 1) * prepage
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("n.*").From("notice as n").Where(where).And("notice_type = ?").OrderBy("n.created_at").Desc().Limit(prepage).Offset(limit).String()
		num, err := o.Raw(sql, con, 6).Values(&notice)
		if err == nil && num > 0 {
			paginatorMap := make(map[string]interface{})
			paginatorMap["total"] = total         //总条数
			paginatorMap["data"] = notice         //分页数据
			paginatorMap["page_num"] = totalpages //总页数
			return paginatorMap, nil
		}
	}
	return nil, err
}

// GetOneNoticeById 新插入的通知列表
func GetOneNoticeById(id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	s := strconv.Itoa(id)
	qb, _ := orm.NewQueryBuilder("mysql")
	where := " n.id = " + s
	sql := qb.Select("n.id as notice_id,n.title,n.notice_type,nr.read_type,n.created_at").From("notice as n").LeftJoin("notice_relation as nr").On("n.id = nr.notice_id").Where(where).GroupBy("n.id").String()
	if num, err := o.Raw(sql).Values(&maps); err == nil && num > 0 {
		for _, v := range maps {
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}

// DeleteNotice 删除通知
func DeleteNotice(id int, user_id int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("notice_relation").Filter("user_id", user_id).Filter("notice_id", id).Delete()
	return err

}

// UpdateNotice 更改阅读状态
func UpdateNoticeById(id int, user_id int) (err error) {
	o := orm.NewOrm()
	_, errs := o.QueryTable("notice_relation").Filter("notice_id", id).Filter("user_id", user_id).Update(orm.Params{
		"read_type": 2,
	})
	return errs
}

//发布动态
func AddDynamic(tetle string, name string, url string, choice_type int, notice_type string, user_id string, teacher_id int, group_type int, send_id int) (id int64, err error) {
	o := orm.NewOrm()
	var v Notice
	v.Title = tetle
	v.Name = name
	v.GroupType = group_type
	v.Url = url
	t, err := strconv.Atoi(notice_type)
	v.NoticeType = t
	v.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	o.Begin()
	//插入通知表
	id, err = o.Insert(&v)
	var relation NoticeRelation
	var Relation []NoticeRelation
	var ids string
	if err == nil && id != 0 {
		if choice_type == 1 {
			//调取rpc获取用户ID
			client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_KG_SERVER"))
			client.UseService(&Kg)
			UserId, _ := strconv.Atoi(user_id)
			baby, _ := Kg.UserPost(UserId)
			ids = baby
		} else {
			ids = user_id
		}
		user_ids := strings.Split(strings.TrimRight(ids, "  "), ",")
		for _, v := range user_ids {
			i, _ := strconv.Atoi(v)
			relation.UserId = i
			relation.SendId = send_id
			relation.NoticeId = int(id)
			relation.ReadType = 1
			Relation = append(Relation, relation)
		}
		if _, err := o.InsertMulti(1, Relation); err != nil {
			o.Rollback()
			return 0, err
		} else {
			o.Commit()
			var notice []orm.Params
			qb, _ := orm.NewQueryBuilder("mysql")
			sql := qb.Select("n.*", "nr.*").From("notice as n").LeftJoin("notice_relation as nr ").
				On("n.id = nr.notice_id").Where("n.id = ?").String()
			_, err = o.Raw(sql, id).Values(&notice)
			var mString string
			client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_MQTT_SERVER"))
			client.UseService(&Mqtt)
			for _, v := range notice {
				jsonData, _ := json.Marshal(v)
				mString = string(jsonData) //text
				for _, v := range user_ids {
					cg, _ := strconv.Atoi(v)
					err = Mqtt.Pub(3, cg, mString) // 1所有平台用户 2所有幼儿园用户 3用户
					fmt.Println(err)
				}
			}
			return 1, nil
		}
	} else {
		o.Rollback()
		return 0, err
	}
	return 0, nil
}

//发布系统通知
func AddSystem(value string) (id int64, err error) {
	o := orm.NewOrm()
	var data map[string]interface{}
	json.Unmarshal([]byte(value), &data)
	var userids []map[string]string
	user_ids, _ := json.Marshal(data["user_id"])
	json.Unmarshal(user_ids, &userids)
	fmt.Println(userids)
	choice_type := int(data["choice_type"].(float64))
	sql := "insert into notice set title = ?,content = ?,notice_type = ? ,type = ?, name = ? , created_at = ?"
	res, err := o.Raw(sql, data["title"], data["content"], data["notice_type"], choice_type, "system", time.Now().Format("2006-01-02 15:04:05")).Exec()
	id, _ = res.LastInsertId()
	if err != nil {
		o.Rollback()
		return 0, nil
	}
	if err == nil && id != 0 {
		client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_MQTT_SERVER"))
		client.UseService(&Mqtt)
		if choice_type == 1 || choice_type == 2 || choice_type == 3 {
			for _, v := range userids {
				sql := "insert into notice_relation set user_id = ?,notice_id = ?,read_type = ?"
				_, err = o.Raw(sql, v["user_id"], id, 1).Exec()
				fmt.Println(err)
			}
			if err != nil {
				o.Rollback()
				return 0, nil
			} else {
				o.Commit()
				var notice []orm.Params
				qb, _ := orm.NewQueryBuilder("mysql")
				sql := qb.Select("n.*", "nr.*").From("notice as n").LeftJoin("notice_relation as nr ").
					On("n.id = nr.notice_id").Where("n.id = ?").String()
				_, err = o.Raw(sql, id).Values(&notice)
				var mString string
				for _, v := range notice {
					jsonData, _ := json.Marshal(v)
					mString = string(jsonData)
				}
				if choice_type == 1 || choice_type == 2 {
					Mqtt.Pub(choice_type, 0, mString)
				} else {
					for _, v := range notice {
						jsonData, _ := json.Marshal(v)
						mString = string(jsonData)
						for _, v := range userids {
							cg, _ := strconv.Atoi(v["user_id"])
							Mqtt.Pub(choice_type, cg, mString)
						}
						break
					}
				}
				return 1, nil
			}
		}
	} else {
		o.Rollback()
		return 0, err
	}
	return 0, nil
}

//认证幼儿园通知
func AttestKindergarten(value string) (id int64, err error) {
	o := orm.NewOrm()
	var data map[string]interface{}
	json.Unmarshal([]byte(value), &data)
	choice_type := int(data["choice_type"].(float64))
	sql := "insert into notice set title = ?,content = ?,notice_type = ? ,type = ?, name = ? , created_at = ?"
	res, err := o.Raw(sql, data["title"], data["content"], data["notice_type"], choice_type, "system", time.Now().Format("2006-01-02 15:04:05")).Exec()
	id, _ = res.LastInsertId()
	if err != nil {
		o.Rollback()
		return 0, nil
	}
	if err == nil && id != 0 {
		client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_MQTT_SERVER"))
		client.UseService(&Mqtt)
		if choice_type == 1 || choice_type == 2 || choice_type == 3 {
			sql := "insert into notice_relation set user_id = ?,notice_id = ?,read_type = ?"
			_, err = o.Raw(sql, data["user_id"], id, 1).Exec()
			if err != nil {
				o.Rollback()
				return 0, nil
			} else {
				o.Commit()
				var notice []orm.Params
				qb, _ := orm.NewQueryBuilder("mysql")
				sql := qb.Select("n.*", "nr.*").From("notice as n").LeftJoin("notice_relation as nr ").
					On("n.id = nr.notice_id").Where("n.id = ?").String()
				_, err = o.Raw(sql, id).Values(&notice)
				var mString string
				for _, v := range notice {
					jsonData, _ := json.Marshal(v)
					mString = string(jsonData)
				}
				if choice_type == 1 || choice_type == 2 {
					Mqtt.Pub(choice_type, 0, mString)
				} else {
					for _, v := range notice {
						jsonData, _ := json.Marshal(v)
						mString = string(jsonData)
						cg, _ := strconv.Atoi(data["user_id"].(string))
						Mqtt.Pub(choice_type, cg, mString)
					}
				}
				return 1, nil
			}
		}
	} else {
		o.Rollback()
		return 0, err
	}
	return 0, nil
}

//邀请通知
func InviteSystem(value string) (id int64, err error) {
	o := orm.NewOrm()
	var data map[string]interface{}
	json.Unmarshal([]byte(value), &data)
	choice_type := int(data["choice_type"].(float64))
	sql := "insert into notice set title = ?,content = ?,notice_type = ? ,type = ?, name = ? , created_at = ?,kindergerten_id = ?,invite_type = ?,types = ?"
	res, err := o.Raw(sql, data["title"], data["content"], data["notice_type"], choice_type, "system", time.Now().Format("2006-01-02 15:04:05"), data["kindergarten_id"], 1, data["type"]).Exec()
	id, _ = res.LastInsertId()
	if err != nil {
		o.Rollback()
		return 0, nil
	}
	if err == nil && id != 0 {
		client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_MQTT_SERVER"))
		client.UseService(&Mqtt)
		sql := "insert into notice_relation set user_id = ?,notice_id = ?,read_type = ?,baby_id = ?"
		_, err = o.Raw(sql, data["user_id"], id, 1, data["baby_id"]).Exec()
		if err != nil {
			o.Rollback()
			return 0, nil
		} else {
			o.Commit()
			var notice []orm.Params
			qb, _ := orm.NewQueryBuilder("mysql")
			sql := qb.Select("n.*", "nr.*").From("notice as n").LeftJoin("notice_relation as nr ").
				On("n.id = nr.notice_id").Where("n.id = ?").String()
			_, err = o.Raw(sql, id).Values(&notice)
			var mString string
			for _, v := range notice {
				v["invite_type"] = 1
				v["kindergarten_id"] = data["kindergarten_id"]
				jsonData, _ := json.Marshal(v)
				mString = string(jsonData)
				cg := int(data["user_id"].(float64))
				err := Mqtt.Pub(choice_type, cg, mString)
				fmt.Println(err)
			}
			return 1, nil
		}
	} else {
		o.Rollback()
		return 0, err
	}
	return 0, nil
}

//发布工作通知
func AddWork(title map[string]interface{}, name string, url_id int, work_notice map[string]interface{}, notice_type string) (id int64, err error) {
	o := orm.NewOrm()
	var v Notice
	//出入接受人
	v.Name = name
	t, err := strconv.Atoi(notice_type)
	v.NoticeType = t
	v.UrlId = url_id
	v.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	v.Title = title["cc"].(string)
	o.Begin()
	//插入通知表
	cid, err := o.Insert(&v)
	v.Title = title["operator"].(string)
	v.Id = 0
	oid, err := o.Insert(&v)
	var relation NoticeRelation
	var Relation []NoticeRelation
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_MQTT_SERVER"))
	client.UseService(&Mqtt)
	if err == nil && cid != 0 {
		cids := work_notice["cc"]
		cuser_id := strings.Replace(strings.Trim(fmt.Sprint(cids), "[]"), " ", ",", -1)
		cuser_ids := strings.Split(cuser_id, ",")
		for _, v := range cuser_ids {
			i, _ := strconv.Atoi(v)
			relation.UserId = i
			relation.NoticeId = int(cid)
			relation.ReadType = 1
			Relation = append(Relation, relation)
		}
		if ids, err := o.InsertMulti(1, Relation); err != nil {
			o.Rollback()
			return 0, err
		} else {
			if oid != 0 {
				var ouser_idd string
				oids := work_notice["operator"]
				ouser_id := strings.Replace(strings.Trim(fmt.Sprint(oids), "[]"), " ", ",", -1)
				ouser_idd += ouser_id + ","
				ouser_ids := strings.Split(ouser_idd, ",")
				for _, v := range ouser_ids {
					if v != "" {
						sql := "insert into notice_relation set user_id = ?,notice_id = ?,read_type = ?"
						_, err = o.Raw(sql, v, oid, 1).Exec()
					}
				}
			}
			o.Commit()
			var o_notice []orm.Params
			var notice []orm.Params
			qb, _ := orm.NewQueryBuilder("mysql")
			sql := qb.Select("n.*", "nr.*").From("notice as n").LeftJoin("notice_relation as nr ").
				On("n.id = nr.notice_id").Where("n.id = ?").String()
			_, err = o.Raw(sql, cid).Values(&notice)

			qb, _ = orm.NewQueryBuilder("mysql")
			sql = qb.Select("n.*", "nr.*").From("notice as n").LeftJoin("notice_relation as nr ").
				On("n.id = nr.notice_id").Where("n.id = ?").String()
			_, err = o.Raw(sql, oid).Values(&o_notice)

			var mString string
			for _, v := range notice {
				jsonData, _ := json.Marshal(v)
				mString = string(jsonData)
			}
			for _, v := range notice {
				i, _ := strconv.Atoi(v["user_id"].(string))
				err = Mqtt.Pub(3, i, mString)
				fmt.Println(err)
			}
			for _, v := range o_notice {
				i, _ := strconv.Atoi(v["user_id"].(string))
				err = Mqtt.Pub(3, i, mString)
				fmt.Println(err)
			}
			return ids, nil
		}
	} else {
		o.Rollback()
		return 0, err
	}
	return 0, nil
}

// DeleteOms oms删除通知
func DeleteOms(id int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("notice").Filter("id", id).Delete()
	if err == nil {
		_, err := o.QueryTable("notice_relation").Filter("notice_id", id).Delete()
		if err != nil {
			return err
		}
	}
	return nil
}
