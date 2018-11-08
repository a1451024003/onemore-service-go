package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"onemore-service-go/utils"
)

type Weather struct {
	Id int `json:"id"`
	Province string `json:"province"`
	CityName string `json:"city_name"`
	Date string `json:"date"`
	DayWeather string `json:"day_weather"`
	NightWeather string `json:"night_weather"`
	DayTemp string `json:"day_temp"`
	NightTemp string `json:"night_temp"`
	DayWind string `json:"day_wind"`
	DayWindScale string `json:"day_wind_scale"`
	NightWind string `json:"night_wind"`
	NightWindScale string `json:"night_wind_scale"`
	AirQuality string `json:"air_quality"`
	Time string `json:"time"`
	Weather string `json:"weather"`
	Temp string `json:"temp"`
	Humidity string `json:"humidity"`
	Wind string `json:"wind"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now"`
}

type City struct {
	Id   		int				`json:"id" orm:"column(id);auto"`
	Name 		string 			`json:"name" orm:"column(name)"`
	ParentId 	int 			`json:"parent_id" orm:"column(parent_id)"`
	CreatedAt   time.Time		`json:"created_at" orm:"column(created_at);auto_now_add"`
	UpdatedAt	time.Time		`json:"updated_at" orm:"column(updated_at);auto_now"`
}

var TimeInterval = []int{2, 5, 8, 11, 14, 17, 20, 23}

func (w *Weather) TableName() string {
	return "weather"
}

func (c *City) TableName() string {
	return "city"
}

func init() {
	orm.RegisterModel(new(Weather), new(City))
}


func (w *Weather) Save(ws []Weather, o orm.Ormer) error {
	_, err := o.InsertMulti(len(ws), ws)

	return err
}

func (w *Weather) Clear(o orm.Ormer) error {
	 _, err := o.Raw("DELETE FROM `weather` WHERE `id` > 0").Exec()

	 return err
}

func (w *Weather) Get() error {
	o := orm.NewOrm()
	r := utils.Redix.Get()
	defer r.Close()

	today := time.Now().Format("2006-01-02")
	rdToday := time.Now().Format("20060102")
	nowHour := time.Now().Hour()
	index := nowHour / 3
	nowTime := fmt.Sprintf("%d:00:00", TimeInterval[index])
	rdNowTime := fmt.Sprintf("%d", TimeInterval[index])
	rdKey := fmt.Sprintf("weather:%s:%s:%s:%s", w.Province, w.CityName, rdToday, rdNowTime)

	if res, _ := redis.String(r.Do("get", rdKey)); res == "" {
		if err := o.QueryTable(w).Filter("province", w.Province).Filter("city_name", w.CityName).Filter("date", today).Filter("time", nowTime).One(w); err == nil {
			v, _ := json.Marshal(w)
			r.Do("setex", rdKey, 10800, string(v))
			return err
		} else {
			return err
		}
	} else {
		json.Unmarshal([]byte(res), &w)

		return nil
	}
}

func (c *City) Get() ([]City, int64, error) {
	o := orm.NewOrm()

	var cs []City
	num, err := o.QueryTable(c).Filter("parent_id", c.ParentId).All(&cs)

	return cs, num, err
}
